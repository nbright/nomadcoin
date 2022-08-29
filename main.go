package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	//go explorer.Start(3000)
	//rest.Start(4000)

	//fmt.Println(os.Args[2:])

	//db.DB()

	// blockchain.BlockChain().AddBlock("First")
	// blockchain.BlockChain().AddBlock("Second")
	// blockchain.BlockChain().AddBlock("Third")

	//defer db.Close()
	//cli.Start()
	difficulty := 6
	target := strings.Repeat("0", difficulty)
	nonce := 1
	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nonce))))
		fmt.Printf("Hash:%s\nTarget:%s\nNonce:%d\nn", hash, target, nonce)
		if strings.HasPrefix(hash, target) {
			return
		} else {
			nonce++
		}
	}
}
