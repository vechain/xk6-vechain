package toolchain

import (
	"context"
	_ "embed"
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/darrenvechain/thorgo"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/darrenvechain/xk6-vechain/random"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func NewTransaction(thor *thorgo.Thor, managers []*txmanager.PKManager, address common.Address) (string, error) {
	manager := random.Element(managers)

	contract, err := NewToolchainTransactor(address, thor.Client, manager)
	if err != nil {
		return "", err
	}

	clauseAmount := 40
	clauses := make([]*tx.Clause, clauseAmount)
	for i := 0; i < clauseAmount; i++ {
		a := random.Uint8()
		b := [32]byte(random.Bytes(32))
		c := [32]byte(random.Bytes(32))
		clause, err := contract.SetBytes32AsClause(a, b, c)
		if err != nil {
			panic(err)
		}
		clauses[i] = clause
	}

	best, err := thor.Blocks.Best()
	if err != nil {
		return "", err
	}
	suggestion, err := thor.Client.FeesPriority()
	if err != nil {
		return "", err
	}

	// TODO: Something better here??
	transaction, err := thor.Transactor(clauses).Build(manager.Address(), &transactions.Options{
		MaxFeePerGas:         best.BaseFee.ToInt(),
		MaxPriorityFeePerGas: suggestion.MaxPriorityFeePerGas.ToInt(),
	})
	if err != nil {
		return "", err
	}

	signature, err := manager.SignTransaction(transaction)
	if err != nil {
		return "", err
	}
	transaction = transaction.WithSignature(signature)
	encoded, err := transaction.MarshalBinary()
	if err != nil {
		return "", err
	}

	return hexutil.Encode(encoded), nil
}

func Deploy(thor *thorgo.Thor, managers []*txmanager.PKManager, amount int) ([]*ToolchainTransactor, error) {
	contracts := make([]*ToolchainTransactor, 0, amount)

	var (
		mu sync.Mutex // mutex to protect concurrent writes
		wg sync.WaitGroup
	)

	for i := range amount {
		manager := managers[i%len(managers)]
		wg.Add(1)
		go func(m *txmanager.PKManager) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			txID, contract, err := DeployToolchain(ctx, thor.Client, manager, &transactions.Options{})
			if err != nil {
				slog.Error("failed to deploy toolchain contract", "error", err, "txID", txID)
				return
			}

			mu.Lock()
			contracts = append(contracts, contract)
			mu.Unlock()
		}(manager)
	}

	wg.Wait()

	if len(contracts) != amount {
		slog.Error("failed to deploy all contracts")
		return nil, errors.New("failed to deploy all contracts")
	}

	return contracts, nil
}
