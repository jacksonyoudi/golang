package main

import (
	"net/rpc"
	"log"
)

type HelloServiceClient struct {
	*rpc.Client // 类型名
}


// 实例化
var _ HelloServiceInterface = (*HelloServiceClient)(nil)



// 拨号,返回一个rpc客户端
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}


func main() {
	// rpc client
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	// 客户端调用函数
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
}