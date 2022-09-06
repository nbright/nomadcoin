package blockchain

import (
	"time"

	"github.com/nbright/nomadcoin/utils"
)

const (
	minerReward int = 50
)

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	Txins     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	Owner  string
	Amount int
}
type TxOut struct {
	Owner  string
	Amount int
}

func makeConinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		Txins:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx

}
