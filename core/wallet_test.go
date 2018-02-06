package core

import (
	"testing"

  	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T){
  private, public := newKeyPair()
  wallet := Wallet{private, public}
  twallet := NewWallet()
  assert.Equal(
    t,
    wallet,
    twallet,
    "Level 1 hash 1 is correct",
  )
}
