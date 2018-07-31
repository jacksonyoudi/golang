package main

import (
	"net"
	"fmt"
	"strings"
)


var onlineConn map[string]net.Conn
var messageQueue = make(chan string, 1000)
var quitChan = make(chan bool)

func CheckError(err error)  {
	if err != nil {
		panic(err)

	}
}

func ProcessInfo(conn net.Conn)  {
	buf := make([]byte, 1024)
	defer conn.Close()
	for {
		Numbuf ,err := conn.Read(buf)
		if err != nil {
			break
		}
		if Numbuf != 0 {
			message := string(buf[0:Numbuf])
			messageQueue <- message
		}

	}
}

func doProcessMessage(message string)  {
	conts := strings.Split(message, "#")
	if len(conts) > 1 {
		addr := conts[0]
		senMes := conts[1]
		if  conn, ok := onlineConn[addr]; ok {
			conn.Write([]byte(senMes))

		}
	}
}

func messageConsumer()  {
	for {
		select {
			case mess := <- messageQueue:
				//对消息解析
				doProcessMessage(mess)
			case <- quitChan:
				break
		}
	}
}

func main()  {
	listen_soct, err := net.Listen("tcp", "127.0.0.1:8090")
	CheckError(err)
	fmt.Printf("listen......")

	go messageConsumer()

	defer listen_soct.Close()

	for {
		conn, err := listen_soct.Accept()
		CheckError(err)
		//保存到映射表中
		addr :=  fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConn[addr] = conn
		go ProcessInfo(conn)
	}
}

