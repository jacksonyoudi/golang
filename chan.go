package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
ch <-
阻塞


 */


func read(ch chan int)  {
	value := <- ch //没写之前，会阻塞
	fmt.Println("value:" + strconv.Itoa(value))
}

func write(ch chan int)  {
	ch <- 10
}


func main()  {
	ch := make(chan int)

	go read(ch)
	go write(ch)
	fmt.Println("test")
	time.Sleep( 1 * time.Second)
	
	
	
}
