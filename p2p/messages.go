package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/nbright/nomadcoin/blockchain"
	"github.com/nbright/nomadcoin/utils"
)

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

// kind와 payload로 JSON으로 인코딩
func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m)
}

// 최신블록을 찾아 메시지로 만들어서 inbox로 보낸다.
func sendNewestBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.BlockChain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		var payload blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		fmt.Println(payload)
	case MessageAllBlocksRequest:

	case MessageAllBlocksResponse:

	}
	fmt.Printf("Peer: %s, Sent a message with of: %d", p.key, m.Kind)
}
