package cmd

import (
	"fmt"

	"github.com/daniilperestoronin/go-chain/core"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := core.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
