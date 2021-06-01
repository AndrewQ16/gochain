package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

var hasher = sha256.New()

type Chain struct {
	blocks []*Block
}

// The Fireship version takes a block and then mines.
// What if instead we confirm the block here and mine in the block function
func (c *Chain) addBlock(block *Block) bool {
	fmt.Printf("%+v\n", block)
	return false
}

// could expand on this method to take a list of transactions as well so the block can contain more than just the
// single transaction for the creation of the genesis block
func (c *Chain) genesis(transactions []Transaction) {

	// verify the math of the transactions, where the first
	// transaction is the "printing" of a certain amount
	// of the coin to an owner

	// {payee: amount}

	// m := make(map[ecdsa.PublicKey]float64)

	for i := 1; i < len(transactions); i++ {

	}

	// turn the  block into a JSON string and hash it
	block := new(Block)
	hasher.Write([]byte(block.prevHash))
	block.timestamp = time.Now()

	// using the binary representation of time and add to the
	// running hash
	time, err := block.timestamp.MarshalBinary()
	if err != nil {
		panic(err)
	}
	hasher.Write(time)

	// hash := make([]byte, hex.EncodedLen(len(hasher.Sum(nil))))
	// str := hex.EncodeToString(hasher.Sum(nil))
	// fmt.Println("hex: ", str)

	// create a single transaction to add to the block
	// create the genesis transaction
	transaction := new(Transaction)
	transaction.Amount = 50.0
	// transaction.payee = publicKey

	block.transactions = append(block.transactions, *transaction)

	c.addBlock(block)
}

func (c *Chain) getBlocks() []*Block {
	return c.blocks
}

// NewChain creates a new chain by first validating the list of transactions passed,
// put them in a new block and mine it. (Then save it to disk and broadcast, not sure
// how to do either yet)
func NewChain(transactions []Transaction) *Chain {
	chain := new(Chain)

	chain.genesis(transactions)

	return chain
}

// Will take a transaction in and confirm that the amount as input
// exists in the wallet of the payer by looking at past transactions.
// In various bitcoin applications, they will have a system for 
// indexing unspent coins. Here I'll traverse the blockchain to find the
// matching transaction
func (b *Chain) verifyInput(t *Transaction) (bool, float64, error){
	// Turn the public key into a hash that represents the address
	// that received funds.
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%v", t.PublicKeyPayer)))
	payerAddress := hasher.Sum(nil)
	_ = payerAddress
	// Go through the blocks and find the matching hash

	return true, 0, nil
}