package main

import (
	"fmt"
)

/*
缓存 channel

c := make(chan int, n)


无缓存
c := make(chan, int) => c := make(chan, int, 0)

 */

 var ch chan int

func testChan(ch chan int)  {
	ch <- 1
	ch <- 1
	fmt.Println("test")
}

func main()  {
	ch := make(chan int, 1)
	go testChan(ch)
	fmt.Printf("running end")
	<- ch
	<- ch
	<- ch

}

