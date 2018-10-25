package main

import (
	"net/rpc"
	"net"
	"log"
	"fmt"
)

// 定义服务名称常量
const HelloServiceName = "path/to/pkg.HelloService"

// 服务接口
type HelloServiceInterface = interface {
	// 不同方法规范
	Hello(request string, reply *string) error
}

// 注册服务的函数
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type Service struct {
	HelloServiceInterface
}

func (s *Service) Hello(request string, reply *string) error {
	fmt.Println("hello:" + request)
	return nil
}

func main()  {

	s := Service{}
	RegisterHelloService(s)

	// 初始化sokcet
	listen ,err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}




}



/* func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
*/