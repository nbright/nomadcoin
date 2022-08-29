package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/nbright/nomadcoin/db"
	"github.com/nbright/nomadcoin/utils"
)

const difficulty int = 2

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"preHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprint(b))))
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(data string, preHash string, height int) *Block {
	block := &Block{
		Data:       data,
		Hash:       "",
		PrevHash:   preHash,
		Height:     height,
		Difficulty: difficulty,
		Nonce:      0,
	}
	block.mine()
	block.persist()
	return block
}
