package client

import (
	"fmt"
	"log"
	"net/rpc"

	"../contract"
)

var port = 8080

func CreateClient() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		// log.Fatal("dialing:" + err)
	}
	return client
}

func PerformRequest(client *rpc.Client, name string) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: name}
	var reply contract.HelloWorldResponse

	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error: ", err)
	}

	return reply
}
