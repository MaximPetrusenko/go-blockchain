package main

import (
	"fmt"
	"log"

	"../goblockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PubliceKeyStr())
	fmt.Println(w.BlockchainAddresss())
}
