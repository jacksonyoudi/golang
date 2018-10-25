package main

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)





func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		// NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
