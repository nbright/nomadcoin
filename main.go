package main

import (
	"fmt"
	"time"
)

func countToTen(c chan<- int) { // 보내기(채널에 쓰기)전용 채널 send only channel
	for i := range [5]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d\n", i)
		c <- i
	}
	close(c)
}

func receive(c <-chan int) { // 받기(채널에서 읽기)전용 채널 receive only channel
	for {
		a, ok := <-c //blocking 기다리고 있음. block operation
		if !ok {
			fmt.Printf("we are done")
			break
		}
		fmt.Printf("received %d\n", a)
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
	receive(c)

}
