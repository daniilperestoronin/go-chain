package cmd

import (
	"fmt"
	"log"

	"github.com/daniilperestoronin/go-chain/core"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !core.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !core.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := core.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{bc}
	defer bc.CloseConnection()

	wallets, err := core.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := core.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := core.NewCoinbaseTX(from, "")
		txs := []*core.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], tx)
	}

	fmt.Println("Success!")
}
