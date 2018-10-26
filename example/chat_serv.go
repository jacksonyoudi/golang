//package main
//
//import (
//	"net"
//	"fmt"
//)
//
//func CheckError(err error)  {
//	if err != nil {
//		panic(err)
//
//	}
//}
//
//func ProcessInfo(conn net.Conn)  {
//	buf := make([]byte, 1024)
//	defer conn.Close()
//	for {
//		Numbuf ,err := conn.Read(buf)
//		if err != nil {
//			break
//		}
//		if Numbuf != 0 {
//			fmt.Printf("has recived: %s", string(buf))
//		}
//
//	}
//}
//
//func main()  {
//	listen_soct, err := net.Listen("tcp", "127.0.0.1:8090")
//	CheckError(err)
//	defer listen_soct.Close()
//
//	for {
//		conn, err := listen_soct.Accept()
//		CheckError(err)
//		go ProcessInfo(conn)
//	}
//}
//
