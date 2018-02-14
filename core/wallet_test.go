package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"bytes"
)

func TestWallet(t *testing.T) {
	private, public := newKeyPair()
	wallet := &Wallet{private, public}
	twallet := NewWallet(private, public)
	assert.Equal(
		t,
		wallet,
		twallet,
		"Level 1 hash 1 is correct",
	)
}

func TestHashPubKey(t *testing.T) {
	_, pubKey := newKeyPair()

	ripemd160Hash := ripemd160.New()
	sha256PubKey := sha256.Sum256(pubKey)
	_, err := ripemd160Hash.Write(sha256PubKey[:])
	if err != nil {
		t.Error("ripemd160Hash Error", err)
	}
	expectedRes := ripemd160Hash.Sum(nil)
	testRes := HashPubKey(pubKey)
	if !bytes.Equal(expectedRes, testRes) {
		t.Error("Expected", expectedRes, " got ", testRes)
	}
}