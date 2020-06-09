package main

import (
	"fmt"

	"./client"
	"./server"
)

func main() {
	go server.StartServer(8080)
	c := client.CreateClient()
	defer c.Close()
	reply := client.PerformRequest(c, "RPC")
	fmt.Println(reply.Message)
}
