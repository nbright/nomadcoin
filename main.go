package main

import (
	"fmt"
	"time"
)

func countToTen(c chan int) {
	for i := range [10]int{} {
		time.Sleep(10 * time.Second)
		c <- i

	}
}

func main() {
	// //go explorer.Start(3000)
	// //rest.Start(4000)
	// //fmt.Println(os.Args[2:])
	// //db.DB()
	// // blockchain.BlockChain().AddBlock("First")
	// // blockchain.BlockChain().AddBlock("Second")
	// // blockchain.BlockChain().AddBlock("Third")
	// defer db.Close()
	// cli.Start()
	// //wallet.Wallet()
	c := make(chan int)
	go countToTen(c)
	fmt.Println("blocking")
	a := <-c //blocking 기다리고 있음.
	fmt.Println("unblocked")
	fmt.Println(a)

}
