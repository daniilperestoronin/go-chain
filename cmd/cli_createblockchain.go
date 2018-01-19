package main

import (
	"fmt"
	"log"

	"github.com/daniilperestoronin/go-chain/core"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !core.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.CreateBlockchain(address, nodeID)
	defer bc.CloseConnection()

	UTXOSet := core.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
