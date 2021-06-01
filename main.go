package main

import (
	"crypto/sha256"
	"fmt"
	"unsafe"
)

func main() {

	// var initialTransactions []Transaction
	// Create some wallets and transactions, the first transaction is the genesis transaction
	// i.e. create some money out of thin air first

	// blockchain := NewChain(initialTransactions)
	genesisWallet := NewWallet("genesis")
	pops := NewWallet("pops")
	

	// empty random address that essentially
	// generates 10 free coins to the user
	sk, address, err:= genesisWallet.newKeyForTransaction()
	if err != nil {
		panic(err)
	}

	tx1 := GenesisTransaction(address, 10)
	tx1.sign(sk)

}
