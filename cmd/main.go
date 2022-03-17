package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	addr             = ":8080"
	chatServerBanner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |
Websocket Chat Room Demo ：Server listening on:%s
`
)

func main() {
	fmt.Printf(chatServerBanner, addr)
	//need some init function
	//register
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		//TODO: create HTTP template
		fmt.Println(res, "Hello")
	})
	log.Fatalln(http.ListenAndServe(addr, nil))
}
