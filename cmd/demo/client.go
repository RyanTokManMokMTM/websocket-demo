package main

import (
	"context"
	"log"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://127.0.0.1:2021/ws", nil)
	if err != nil {
		panic(err)
	}

	defer c.Close(websocket.StatusInternalError, "internal error")

	//send json
	err = wsjson.Write(ctx, c, "Hello WS server")
	if err != nil {
		panic(err)
	}

	var v interface{}
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		panic(err)
	}
	log.Printf("Received from server %v", v)

	c.Close(websocket.StatusNormalClosure, "")

}
