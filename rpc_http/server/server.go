package server

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"log"

	"../contract"
)

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

func StartServer(port int) {
	handler := &HelloWorldHandler{}
	rpc.Register(handler)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer l.Close()

	log.Printf("Server starting on port %v\n", port)

	http.Serve(l, nil)
}
