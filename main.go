package main

import (
	"log"

	"github.com/gorilla/websocket"
)

const URL string = "wss://api.upbit.com/websocket/v1"
const REQUEST string = `[{"ticket":"test"},{"type":"trade","codes":["KRW-XRP"]},{"format":"SIMPLE"}]`

func main() {
	c, _, err := websocket.DefaultDialer.Dial(URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = c.WriteMessage(websocket.TextMessage, []byte(REQUEST))

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(message))
	}
}
