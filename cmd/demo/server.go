package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func main() {
	//create a simple http server

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "HTTP Demo")
	})

	http.HandleFunc("/ws", func(res http.ResponseWriter, req *http.Request) {
		//handshaking and upgrade to websocket
		//Header info:
		/*
			handshaking:
			Client -> syn -> server
			Server -> syn ,ack -> client
			Client -> ack

			Connection : Upgrade  -> http upgrade
			Upgrade : websocket -> upgrade to websocket
			Sec-Websocket-Key:a random string that identify http and websocket
				using SHA-1 to encode + specific string ....
				*Server will use this information and become the value of Sec-Websocket-Accept(response)
			Sec-Websocket-Version : 13 (RFC6455 requires version 13)

			Response code : 101 instead http status code 200
		*/
		conn, err := websocket.Accept(res, req, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "server internal error")

		//setting request context with timeout
		//return a new context with timeout and cancel function
		ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
		defer cancel()

		var v interface{}
		//reading ws json message form, copy conn message to v
		err = wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Received client:%v", v)

		err = wsjson.Write(ctx, conn, "Hello WS Client")
		if err != nil {
			log.Println(err)
			return
		}

		conn.Close(websocket.StatusNormalClosure, "")
	})

	log.Fatalln(http.ListenAndServe(":2021", nil))
}
