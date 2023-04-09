package main

import (
	"fmt"
	"log"

	"../go-blockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PubliceKeyStr())
	fmt.Println(w.BlockchainAddresss())

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddresss(), "B", 1.0)
	fmt.Printf("signature %s\n", t.GenerateSignature())

}
