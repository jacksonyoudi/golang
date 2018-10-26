package main

import (
	"net"
	"bufio"
	"os"
	"strings"
	"fmt"
)

func checkError(err error)  {
	if err != nil {
		panic(err)

	}
}

func messageSent(conn net.Conn)  {
	var input string
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _, _  := reader.ReadLine()
		input = string(data)

		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			break
			//os.Exit(0)
		}

		_, err := conn.Write(data)
		if err != nil {
			conn.Close()
			fmt.Println("client connenct failure! " + err.Error())
			break

		}

	}
}



func main()  {
	//conn, err := net.Dial("tcp", "127.0.0.1:8090")
	//checkError(err)
	//
	//defer conn.Close()
	//conn.Write([]byte("hello, world"))
	//fmt.Println("sent ok")

	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	checkError(err)
	defer conn.Close()
	go messageSent(conn)

	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)
		checkError(err)
		fmt.Printf("recevier message: %s", string(buf))

	}

}

