package cmd

import (
	"fmt"
	"log"

	"github.com/daniilperestoronin/go-chain/core"
	"github.com/daniilperestoronin/go-chain/crypto"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !core.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{bc}
	defer bc.CloseConnection()

	balance := 0
	pubKeyHash, _ := crypto.BitcoinBase58.Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
