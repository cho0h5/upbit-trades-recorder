package main

import (
	"log"

	"github.com/gorilla/websocket"

	"go.mongodb.org/mongo-driver/bson"
)

const URL string = "wss://api.upbit.com/websocket/v1"
const REQUEST string = `[{"ticket":"test"},{"type":"trade","codes":["KRW-XRP"]},{"format":"SIMPLE"}]`

func main() {
	// DB initialize
	db := ConnectDB()
	defer db.Disconnect()

	// connect websocket
	c, _, err := websocket.DefaultDialer.Dial(URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// request message type
	err = c.WriteMessage(websocket.TextMessage, []byte(REQUEST))

	for {
		//read message
		_, message, err := c.ReadMessage()

		// print message
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(string(message))

		// insert to DB
		var doc interface{}
		err = bson.UnmarshalExtJSON(message, false, &doc)
		if err != nil {
			log.Fatal(err)
		}
		db.InsertOne(doc)
	}
}
