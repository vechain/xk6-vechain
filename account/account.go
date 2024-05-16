package account

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

// Acccount is a k6 external module that provides access to common address operations.
type Acccount struct{}

func (a *Acccount) NewPrivateKey() string {
	key := mustCreateKey()
	return hex.EncodeToString(key.D.Bytes())
}
func (a *Acccount) NewAddress() string {
	key := mustCreateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)
	return addr.String()
}

func mustCreateKey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	return key
}
