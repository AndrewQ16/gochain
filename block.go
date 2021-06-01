package main

import (
	"time"
)

/*
	For this implementation, each block will contain a full list
	of the transactions.

	BTC for example only will contain the "root"
	of a merkel tree as a hash and if the user wants, they can download
	the full block
*/
type Block struct {
	prevHash     []byte
	transactions []Transaction
	nonce        int
	timestamp    time.Time // will probably use the "wall clock", this is mainly used for telling/displaying time
}

/*
	If I feel like expanding the project to get as close to BTC
	then I can consider this definition of a block.
*/
type FullBlock struct {
	prevHash  []byte
	nonce     int
	transactions []byte
	timestamp time.Time
}

/*
	Represent a node in the merkel tree.
	Not going to do this in this implementation of a blockchain
*/
type Node struct {
	hash  []byte
	left  *Node
	right *Node
}

/*
Recall:

Sign(Message, sk) = Signature

Verify(Message, Signature, pk) = T/F
*/

/*
	How to store transactions in blocks:

*/

func NewBlock(tx *Transaction) *Block {
	
	return nil
}

func (b *Block) addTransaction(tx* Transaction) error{
	return nil
}

func (b *Block) mine(c *Chain) error {
	return nil
}
