package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// A representation of the wallet.
type Wallet struct {
	addressAndAmount map[ecdsa.PublicKey]*miniWallet //key is the address used for that Tx
	keys             []*ecdsa.PrivateKey
	amount           float64
	name             string
}

// Returns a new wallet
func NewWallet(name string) *Wallet {
	wallet := new(Wallet)
	wallet.name = name

	// hasher := sha256.New()
	// hasher.Write([]byte(name))
	// wallet.address = hasher.Sum(nil)

	wallet.keys = make([]*ecdsa.PrivateKey, 100)
	wallet.addressAndAmount = make(map[ecdsa.PublicKey]*miniWallet)

	// generate a 100 keys for the wallet to immediately use
	for i := 0; i < 100; i++ {
		pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			panic(err)
		}
		wallet.keys[i] = pk
	}

	return wallet
}

// returns a key pair that will be used for a transaction. It also
// saves a note of this for your wallet to reference
func (w *Wallet) newKeyForTransaction() (*ecdsa.PrivateKey, []byte, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	w.keys = append(w.keys, key)

	hasher := sha256.New()
	_, err = hasher.Write([]byte(fmt.Sprintf("%v", key.PublicKey)))

	if err != nil {
		return nil, nil, err
	}

	address := hasher.Sum(nil)
	w.addressAndAmount[key.PublicKey] = new(miniWallet)
	w.addressAndAmount[key.PublicKey].address = address
	return key, address, nil
}

type miniWallet struct {
	address []byte
	amount  float64
}
