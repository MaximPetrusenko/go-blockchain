package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{} //emtpy transaction pool for bc
	return b
}

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b

}

func (b *Block) Print() {
	fmt.Printf("timestamp	%d\n", b.timestamp)
	fmt.Printf("nonce	%d\n", b.nonce)
	fmt.Printf("previous_hash	%x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"` //From capital letter to make a it a public Property
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

type Transaction struct {
	senderBlockchainAdress     string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("  sender_blockchain_address  %s\n", t.senderBlockchainAdress)
	fmt.Printf("  recipient_blockchain_address  %s\n", t.recipientBlockchainAddress)
	fmt.Printf("  value  %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_address"`
		Recipient string  `json:"recipient_address`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAdress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()
}
