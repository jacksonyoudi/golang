package main

import (
	"fmt"
)

//协程和进程相比，协程优势在于轻量级，可以轻松创建
//上百万个而不会导致系统资源衰竭
// go协程的实现
/*
go 函数名


channel的写和读

ch <- c
ch -> c


 */

 //func testRoutine() {
 //	fmt.Printf("this is one routine!!!")
 //}

func Add(x, y int, quit chan int)  {
	z := x + y
	fmt.Println(z)
	quit <- z
}






func main()  {
	//go testRoutine()
	//ch := make(chan int)
	//

	chs := make([]chan int, 10)
	for i := 0;i< 10; i++ {
		chs[i] = make(chan int)
		go Add(i, i, chs[i])
	}

	//time.Sleep( 1 * time.Second)

	for _,v := range chs {
		<- v //阻塞着
	}
}

