package main

import (
	"fmt"
)

//func worker(id int) chan int {
//	c := make(chan int)
//	go func() {
//		for {
//			fmt.Printf("worker %d received %c\n", id, <-c)
//		}
//	}()
//	return c
//}

func worker(id int, c chan int)  {




	for {
		fmt.Printf("worker %d received %d\n", id, <-c)
	}
}



//func chanDemo() {
//	var channels [10]chan int
//
//	for i := 0; i < 10; i++ {
//		channels[i] = worker(i)
//	}
//
//	for i := 0; i < 10; i++ {
//		channels[i] <- 'a' + i
//	}
//
//	for i := 0; i < 10; i++ {
//		channels[i] <- 'A' + i
//	}
//
//}

func channelClose()  {
	c := make(chan int, 3)
	c <- 'A'
	c <- 'B'
	c <- 'C'
	c <- 'D'
	close(c)
}


func bufferChannel()  {
	c := make(chan int, 3)
	c <- 'A'
	c <- 'B'
	c <- 'C'
	c <- 'D'

}



func main() {
	//chanDemo()
	bufferChannel()
}


/*
单向 chan <-
 */