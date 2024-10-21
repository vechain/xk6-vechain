package toolchain

import (
	_ "embed"
	"errors"
	"log/slog"
	"strings"
	"sync"

	"github.com/darrenvechain/thor-go-sdk/crypto/transaction"
	"github.com/darrenvechain/thor-go-sdk/thorgo"
	"github.com/darrenvechain/thor-go-sdk/thorgo/accounts"
	"github.com/darrenvechain/thor-go-sdk/txmanager"
	"github.com/darrenvechain/xk6-vechain/random"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

//go:embed Toolchain.abi
var ABI string

//go:embed Toolchain.bin
var Bytecode string

var (
	toolchainABI, abiErr = abi.JSON(strings.NewReader(ABI))
)

func NewTransaction(thor *thorgo.Thor, managers []*txmanager.PKManager, address common.Address) (string, error) {
	manager := random.Element(managers)

	if abiErr != nil {
		return "", abiErr
	}
	contract := thor.Account(address).Contract(&toolchainABI)

	clauseAmount := 40
	clauses := make([]*transaction.Clause, clauseAmount)
	for i := 0; i < clauseAmount; i++ {
		a := random.Uint8()
		b := [32]byte(random.Bytes(32))
		c := [32]byte(random.Bytes(32))
		clause, err := contract.AsClause("setBytes32", a, b, c)
		if err != nil {
			panic(err)
		}
		clauses[i] = clause
	}

	tx, err := thor.Transactor(clauses, manager.Address()).Build()
	if err != nil {
		return "", err
	}

	signature, err := manager.SignTransaction(tx)
	if err != nil {
		return "", err
	}
	tx = tx.WithSignature(signature)

	encoded, err := tx.Encoded()
	if err != nil {
		return "", err
	}

	return encoded, nil
}

func Deploy(thor *thorgo.Thor, managers []*txmanager.PKManager, amount int) ([]*accounts.Contract, error) {
	contracts := make([]*accounts.Contract, 0, amount)
	if abiErr != nil {
		return nil, abiErr
	}
	deployer := thor.Deployer(common.Hex2Bytes(Bytecode), &toolchainABI)

	var (
		mu sync.Mutex // mutex to protect concurrent writes
		wg sync.WaitGroup
	)

	for i := range amount {
		manager := managers[i%len(managers)]
		wg.Add(1)
		go func(m *txmanager.PKManager) {
			defer wg.Done()

			contract, txID, err := deployer.Deploy(manager)
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
