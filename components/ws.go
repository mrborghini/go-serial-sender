package components

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type Subscription struct {
	Subscribe int    `json:"subscribe"`
	Metadata  string `json:"metadata"`
}

type WS struct {
	conn *websocket.Conn
}

func NewWebsocket() *WS {
	return &WS{conn: nil}
}

func (w *WS) Connect(channel string) {
	const wsUrl = "ws://192.168.178.24:8080/realtime"
	conn, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	w.conn = conn
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	jsonSubscribe, _ := json.Marshal(Subscription{Subscribe: 4, Metadata: channel})
	w.send(string(jsonSubscribe))

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		messageString := string(message)
		fmt.Println(messageString)
	}
}

func (w *WS) send(data string) {
	if w.conn == nil {
		fmt.Println("No conn yet...")
		return
	}

	w.conn.WriteMessage(websocket.TextMessage, []byte(data))
	fmt.Println("Sent: " + data)
}

func (w *WS) Transmit(channel string, data string) {
	jsonSubscribe, _ := json.Marshal(Subscription{Subscribe: 5, Metadata: fmt.Sprintf("%s:%s", channel, data)})
	w.send(string(jsonSubscribe))
}
