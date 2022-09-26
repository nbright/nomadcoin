package main

import (
	"github/nbright/nomadcoin/cli"
	"github/nbright/nomadcoin/db"
)

func main() {
	//go explorer.Start(3000)
	//rest.Start(4000)

	//fmt.Println(os.Args[2:])

	//db.DB()

	// blockchain.BlockChain().AddBlock("First")
	// blockchain.BlockChain().AddBlock("Second")
	// blockchain.BlockChain().AddBlock("Third")

	defer db.Close()
	cli.Start()

	//wallet.Wallet()
}
