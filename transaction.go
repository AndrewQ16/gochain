package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"math/big"
)

/*
	Transactions contains the 32 byte array to represent
	the address of each wallet and the amount being transferred.

	Byte arrays are used so that transactions can be used as map keys
	with the public key that is used to verify it
*/
type Transaction struct {
	Amount         float64  `json:"amount"`         // should money be represented as this type?
	PublicKeyPayer ecdsa.PublicKey `json:"PublicKeyPayer"` // wallet public key for the PublicKeyPayer
	AddressPayee   []byte `json:"AddressPayee"`   // wallet address for the AddressPayee
	sig1, sig2     *big.Int
}

// Turn the transaction into a json string
// Currently just being used for debugging
func (t *Transaction) toJSON() ([]byte, error) {
	// Might be possible to remove the error return since
	// we know the format of what is passed always
	value, err := json.Marshal(&t)
	if err != nil {
		return nil, err
	}
	return value, nil
}


// Sign the transaction to make it official
func(t *Transaction) sign(sk *ecdsa.PrivateKey) error{

	hasher := sha256.New()
	hash, err := t.toJSON()
	if err != nil {
		return err
	}
	hasher.Write(hash)

	t.sig1, t.sig2, err = ecdsa.Sign(rand.Reader, sk, hasher.Sum(nil))

	return err
}

// Returns a new Transaction.
func NewTransaction(address []byte, pk ecdsa.PublicKey, amount float64) *Transaction {
	tx := new(Transaction)
	tx.AddressPayee = address
	tx.PublicKeyPayer = pk
	tx.Amount = amount
	return tx
}

func GenesisTransaction(address []byte, amount float64) *Transaction {
	tx := new(Transaction)
	var nilPk ecdsa.PublicKey
	tx.PublicKeyPayer = nilPk
	tx.Amount = amount
	tx.AddressPayee = address
	return tx
}

/*
Recall:

Sign(Message, sk) = Signature

Verify(Message, Signature, pk) = T/F
*/
