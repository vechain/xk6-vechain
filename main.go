package xk6_vechain

import (
	"github.com/darrenvechain/xk6-vechain/account"
	"github.com/darrenvechain/xk6-vechain/transactions"
	"go.k6.io/k6/js/modules"
)

type RootModule struct {
	Account      account.Acccount
	Transactions transactions.Transactions
}

func init() {
	modules.Register("k6/x/vechain", new(RootModule))
	modules.Register("k6/x/vechain/accounts", new(account.Acccount))
	modules.Register("k6/x/vechain/transactions", new(transactions.Transactions))
}
