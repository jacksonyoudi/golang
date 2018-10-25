package main

import (
	"net/rpc"
	"net"
	"log"
)

type HelloService struct {
	
}

func (hs *HelloService) Hello(request string, reply *string)  error{
	*reply = "hello:" + request
	return nil

}

func main() {
	//注册一个服务
	rpc.RegisterName("HelloService", new(HelloService))

	// 监听在1234端口
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	// 服务启动
	go rpc.ServeConn(conn)
}