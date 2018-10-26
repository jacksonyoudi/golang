package main

import (
	"time"
	"fmt"
)

/*
select
linux 很早就引入的函数,用来实现非阻塞的一种方式
基于事件机制 epoll

go语言直接在语言级别上支持select关键字,用于处理异步io问题





 */
//
// select {
// case <-chan1:
// 	//如果chan1成功读取数据，则进行case处理语句
// 	case chan2 <-: //如果成功向chan2写入数据,则进行case处理语句
// 	default: //如果上面都没成功, 则进入default处理流程

//}

//func main()  {
//	ch := make(chan  int)
//
//	go func(ch chan int) {
//		ch <- 1
//	}(ch)
//	time.Sleep( 1 * time.Second)
//	select {
//	case <-ch:
//		fmt.Printf("come to read")
//	default:
//		fmt.Println("default")
//	}
//}


/*
超时控制



 */

//func main()  {
//	ch := make(chan int)
//	timeout := make(chan bool, 1)
//
//	go func() {
//		time.Sleep(time.Second)
//		timeout <- true
//	}()
//
//	select {
//		case <-ch: //一直等待读取数据
//			fmt.Printf("come to read chan!")
//		case <- timeout: // 一直读取不到数据，timeout 就会有数据，就会执行
//			fmt.Println("come to timeout")
//	}
//	fmt.Println("end of code!")
//}



func main()  {
	ch := make(chan int)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()

	select {
	case <-ch: //一直等待读取数据
		fmt.Printf("come to read chan!")
	case <- time.After(time.Second): // 一直读取不到数据，timeout 就会有数据，就会执行
		fmt.Println("come to timeout")
	}
	fmt.Println("end of code!")
}

//
//for {
//	select {
//
//}
//}