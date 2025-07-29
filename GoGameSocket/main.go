package main

import (
	"fmt"
	"gamedev/backend/server"
	"net/http"
)

func main() {
	go server.QueueMessageTask()
	go server.WaitingRoomTask()
	http.HandleFunc("/ws", server.Handler)
	http.ListenAndServe(":3002", nil)
	fmt.Println("it's working.")
}
