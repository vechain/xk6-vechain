package transactions

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vechain/thor/thor"
	"github.com/vechain/thor/tx"
	"math/big"
)

type Transactions struct {
	BaseURL string
}

// Init initializes the Transactions module.
func (t *Transactions) Init(url string) {
	t.BaseURL = url
}

// SignTransaction signs a raw transaction with a private key and returns the signed transaction.
func (t *Transactions) SignTransaction(rawTx string, privateKey string) string {
	return "TODO"
}

func (t *Transactions) BuildTransaction() string {
	to := thor.MustParseAddress("0x45A1b2d88d8C59f98Fa1aB81060Dee191ce8F326")
	clause := tx.NewClause(&to).WithValue(big.NewInt(1))

	newTx := new(tx.Builder).Clause(clause).Build()

	// rawTxWriter is a io.Writer
	rawTxWriter := new(bytes.Buffer)
	err := newTx.EncodeRLP(rawTxWriter)

	if err != nil {
		panic(err)
	}

	return common.Bytes2Hex(rawTxWriter.Bytes())
}

//
//func decode(raw string) (*tx.Transaction, error) {
//	data, err := hexutil.Decode(raw)
//	if err != nil {
//		return nil, err
//	}
//	var tx *tx.Transaction
//	if err := rlp.DecodeBytes(data, &tx); err != nil {
//		return nil, err
//	}
//	return tx, nil
//}
