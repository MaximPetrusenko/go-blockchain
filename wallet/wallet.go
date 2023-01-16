package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"io"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey *ecdsa.PublicKey
}

func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey((elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey
	return w
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PubliceKey() *ecdsa.PubliceKey {
	return w.publicKey
}

func (w *Wallet) PubliceKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}