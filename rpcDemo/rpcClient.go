package main

import (
	"net/rpc"
	"log"
	"fmt"
)

func main() {
	// rpc拨号
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	// rpc调用 HelloService.Hello函数
	// 第一个参数是用点号链接的RPC服务名字和方法名字
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
