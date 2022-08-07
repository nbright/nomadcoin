package main

import (
	"github/nbright/nomadcoin/explorer"
	"github/nbright/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
