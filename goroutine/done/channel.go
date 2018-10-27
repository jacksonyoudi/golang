package main

import (
	"fmt"
	"sync"
)


func worker(id int, c chan int, done chan bool)  {

	for n := range c {
		fmt.Printf("worker %d received %d\n", id, <-n)
	}
	done <- true
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


	var wg sync.WaitGroup
	wg.Add(20)


	wg.Done()

}


/*
单向 chan <-
 */