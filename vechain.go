package xk6_vechain

import (
	"context"
	"errors"
	"log/slog"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/darrenvechain/thorgo"
	"github.com/darrenvechain/thorgo/builtins"
	"github.com/darrenvechain/thorgo/crypto/hdwallet"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vechain/xk6-vechain/toolchain"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

type Client struct {
	wallet   *hdwallet.Wallet
	thor     *thorgo.Thor
	chainTag byte
	vu       modules.VU
	metrics  vechainMetrics
	opts     *options
	accounts int
	managers []*txmanager.PKManager
}

func (c *Client) Accounts() []string {
	addresses := make([]string, 0)
	for _, i := range c.managers {
		addresses = append(addresses, i.Address().String())
	}
	return addresses
}

func (c *Client) DeployToolchain(amount int) ([]string, error) {
	contracts, err := toolchain.Deploy(c.thor, c.managers, amount)
	if err != nil {
		return nil, err
	}
	addresses := make([]string, 0)
	for _, contract := range contracts {
		addresses = append(addresses, contract.Address().String())
	}
	return addresses, nil
}

func (c *Client) NewToolchainTransaction(address string) (string, error) {
	addr := common.HexToAddress(address)
	return toolchain.NewTransaction(c.thor, c.managers, addr)
}

// Fund sends VET and VTHO to the accounts after the index, funded by the accounts before the index.
// The amount is the amount of VET & VTHO to send, represented as hex.
// Example: thor solo only funds the first 10 accounts [0-9], so specify 10 as the start index.
func (c *Client) Fund(start int, amount string) error {
	if start > len(c.managers) {
		return errors.New("start index is greater than the number of accounts")
	}

	// funder index -> clauses to send
	clauses := make(map[int][]*tx.Clause)
	vtho, err := builtins.NewVTHO(c.thor.Client())
	if err != nil {
		return err
	}

	for i := start; i < len(c.managers); i++ {
		fundee := c.managers[i].Address()
		funderIndex := i % start

		value := new(big.Int)
		value.SetString(amount, 16)

		vetClause := tx.NewClause(&fundee).WithValue(value)
		vthoClause, err := vtho.TransferAsClause(fundee, value)
		if err != nil {
			return err
		}

		funderClauses := clauses[funderIndex]
		if funderClauses == nil {
			funderClauses = make([]*tx.Clause, 0)
		}

		clauses[funderIndex] = append(funderClauses, vetClause, vthoClause)
	}

	var (
		wg        sync.WaitGroup
		clauseErr error
	)

	for i, clauses := range clauses {
		wg.Add(1)
		manager := c.managers[i]
		go func(i *txmanager.PKManager, clauses []*tx.Clause) {
			defer wg.Done()
			for i := 0; i < len(clauses); i += 100 {
				end := i + 100
				if end > len(clauses) {
					end = len(clauses)
				}

				tx, err := manager.SendClauses(clauses[i:end], nil)
				if err != nil {
					clauseErr = err
					return
				}
				ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
				_, err = tx.Wait(ctx)
				cancel()
				if err != nil {
					clauseErr = err
					return
				}
			}
		}(manager, clauses)
	}

	wg.Wait()

	if clauseErr != nil {
		return clauseErr
	}

	return nil
}

var blocks sync.Map

func (c *Client) pollForBlocks() {
	prev, err := c.thor.Blocks().Best()
	if err != nil {
		return
	}

	for range time.Tick(500 * time.Millisecond) {
		block, err := c.thor.Blocks().Best()
		if err != nil {
			continue
		}

		if block.Number > prev.Number {
			blockTimestampDiff := time.Unix(block.Timestamp, 0).Sub(time.Unix(prev.Timestamp, 0))
			tps := float64(len(block.Transactions)) / blockTimestampDiff.Seconds()

			prev = block

			rootTS := metrics.NewRegistry().RootTagSet()
			if c.vu != nil && c.vu.State() != nil && rootTS != nil {
				if _, loaded := blocks.LoadOrStore(c.opts.URL+strconv.FormatInt(block.Number, 10), true); loaded {
					// We already have a block number for this client, so we can skip this
					continue
				}

				baseFee, _ := block.BaseFee.ToInt().Float64()
				baseFeePercent := baseFee * 100 / 10_000_000_000_000

				slog.Info("base fee", "val", baseFeePercent, "block", block.Number)

				blockTime := time.Unix(block.Timestamp, 0)

				metrics.PushIfNotDone(c.vu.Context(), c.vu.State().Samples, metrics.ConnectedSamples{
					Samples: []metrics.Sample{
						{
							TimeSeries: metrics.TimeSeries{
								Metric: c.metrics.Block,
								Tags: rootTS.WithTagsFromMap(map[string]string{
									"transactions": strconv.Itoa(len(block.Transactions)),
									"gas_used":     strconv.Itoa(int(block.GasUsed)),
									"gas_limit":    strconv.Itoa(int(block.GasLimit)),
								}),
							},
							Value: float64(block.Number),
							Time:  blockTime,
						},
						{
							TimeSeries: metrics.TimeSeries{
								Metric: c.metrics.GasUsed,
								Tags:   rootTS,
							},
							Value: float64(block.GasUsed),
							Time:  blockTime,
						},
						{
							TimeSeries: metrics.TimeSeries{
								Metric: c.metrics.TPS,
								Tags:   rootTS,
							},
							Value: tps,
							Time:  blockTime,
						},
						{
							TimeSeries: metrics.TimeSeries{
								Metric: c.metrics.BlockTime,
								Tags: rootTS.WithTagsFromMap(map[string]string{
									"block_timestamp_diff": blockTimestampDiff.String(),
								}),
							},
							Value: float64(blockTimestampDiff.Milliseconds()),
							Time:  blockTime,
						},
						{
							TimeSeries: metrics.TimeSeries{
								Metric: c.metrics.BaseFee,
								Tags:   rootTS,
							},
							Value: baseFeePercent,
							Time:  blockTime,
							Metadata: map[string]string{
								"block": strconv.Itoa(int(block.Number)),
							},
						},
					},
				})
			}
		}
	}
}
