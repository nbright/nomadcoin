package blockchain

import (
	"fmt"
	"sync"

	"github.com/nbright/nomadcoin/db"
	"github.com/nbright/nomadcoin/utils"
)

// func (b *BlockChain) listBlocks() {
// 	for _, block := range b.blocks {
// 		fmt.Printf("Data: %s\n", block.data)
// 		fmt.Printf("Hash: %s\n", block.hash)
// 		fmt.Printf("PreHash: %s\n", block.prevHash)
// 	}
// }

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

type blockChain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockChain
var once sync.Once

func (b *blockChain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockChain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockChain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}
func (b *blockChain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)

		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockChain) recalculateDifficulty() int {
	allBlock := b.Blocks()
	newestBlock := allBlock[0]
	lastRecalculatedBlock := allBlock[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval
	if actualTime < (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockChain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		// recalculate the difficulty
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func (b *blockChain) txOuts() []*TxOut {
	var txOuts []*TxOut
	blocks := b.Blocks()
	for _, block := range blocks {
		for _, tx := range block.Transactions {
			txOuts = append(txOuts, tx.TxOuts...)
		}
	}

	return txOuts
}

func BlockChain() *blockChain {
	if b == nil {
		once.Do(func() {
			b = &blockChain{}
			// search for "checkpoint" on the db
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock()
			} else {
				fmt.Printf("Restoring...")
				// restore b from bytes
				b.restore(checkpoint)
			}

		})
	}
	fmt.Printf("3NewestHash: %s\nHeight:%d", b.NewestHash, b.Height)
	return b
}
