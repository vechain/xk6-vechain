package xk6_vechain

import (
	"github.com/darrenvechain/thor-go-sdk/client"
	"strconv"
	"sync"
	"time"

	"github.com/darrenvechain/thor-go-sdk/crypto/hdwallet"
	"github.com/darrenvechain/thor-go-sdk/thorgo"
	"github.com/darrenvechain/thor-go-sdk/txmanager"
	"github.com/darrenvechain/xk6-vechain/toolchain"
	"github.com/ethereum/go-ethereum/common"
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
	managers map[int]*txmanager.PKManager
}

func (c *Client) DeployToolchain(amount int) ([]string, error) {
	managers := c.managerList()
	contracts, err := toolchain.Deploy(c.thor, managers, amount)
	if err != nil {
		return nil, err
	}
	addresses := make([]string, 0)
	for _, contract := range contracts {
		addresses = append(addresses, contract.Address.String())
	}
	return addresses, nil
}

func (c *Client) NewToolchainTransaction(address string) (string, error) {
	managers := c.managerList()
	addr := common.HexToAddress(address)
	return toolchain.NewTransaction(c.thor, managers, addr)
}

// WaitForTx
func (c *Client) WaitForTx(txID string) (*client.TransactionReceipt, error) {
	return c.thor.Transaction(common.HexToHash(txID)).Wait()
}

func (c *Client) managerList() []*txmanager.PKManager {
	managers := make([]*txmanager.PKManager, 0)
	for _, i := range c.managers {
		managers = append(managers, i)
	}
	return managers
}

var blocks sync.Map

func (c *Client) pollForBlocks() {
	prev, err := c.thor.Blocks.Best()
	if err != nil {
		return
	}

	for range time.Tick(500 * time.Millisecond) {
		block, err := c.thor.Blocks.Best()
		if err != nil {
			continue
		}

		if block.Number > prev.Number {
			blockTimestampDiff := time.Unix(int64(block.Timestamp), 0).Sub(time.Unix(int64(prev.Timestamp), 0))
			tps := float64(len(block.Transactions)) / float64(blockTimestampDiff.Seconds())

			prev = block

			rootTS := metrics.NewRegistry().RootTagSet()
			if c.vu != nil && c.vu.State() != nil && rootTS != nil {
				if _, loaded := blocks.LoadOrStore(c.opts.URL+strconv.FormatUint(block.Number, 10), true); loaded {
					// We already have a block number for this client, so we can skip this
					continue
				}

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
							Time:  time.Now(),
						},
						{
							TimeSeries: metrics.TimeSeries{
								Metric: c.metrics.GasUsed,
								Tags: rootTS.WithTagsFromMap(map[string]string{
									"block": strconv.Itoa(int(block.Number)),
								}),
							},
							Value: float64(block.GasUsed),
							Time:  time.Now(),
						},
						{
							TimeSeries: metrics.TimeSeries{
								Metric: c.metrics.TPS,
								Tags:   rootTS,
							},
							Value: tps,
							Time:  time.Now(),
						},
						{
							TimeSeries: metrics.TimeSeries{
								Metric: c.metrics.BlockTime,
								Tags: rootTS.WithTagsFromMap(map[string]string{
									"block_timestamp_diff": blockTimestampDiff.String(),
								}),
							},
							Value: float64(blockTimestampDiff.Milliseconds()),
							Time:  time.Now(),
						},
					},
				})
			}
		}
	}
}
