package p2p

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nbright/nomadcoin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)

	utils.HandleErr(err)

}

func AddPeer(address, port string) {

}

/** 아주 중요,서버에서  메시지 읽어서, 보내기
var conns []*websocket.Conn
func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	conns.append(conns, conn)
	utils.HandleErr(err)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {

			break
		}
		for _, aConn := range conns {
			if aConn != conn {
				aConn.WriteMessage(websocket.TextMessage, p)
			}
		}

	}

}
*/
