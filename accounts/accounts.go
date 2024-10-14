package accounts

import "github.com/ethereum/go-ethereum/crypto"

type Account struct {
}

func (a *Account) Generate() map[string]string {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	account := make(map[string]string)
	account["privateKey"] = key.D.String()
	account["address"] = crypto.PubkeyToAddress(key.PublicKey).Hex()

	return account
}
