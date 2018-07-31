package main

import (
	"time"
	"fmt"
)

/*

有缓存的chan,当chan满的时候, <- 就会阻塞
当chan为空的时候,  -> 就会阻塞

 */



func one(message chan string) {
	message <- "hello goroutine! 1" //阻塞
	message <- "hello goroutine! 2" //阻塞
	message <- "hello goroutine! 3" //阻塞
	message <- "hello goroutine! 4" //阻塞
}

func two(message chan string)  {
	time.Sleep(2 * time.Second)
	str := <- message
	str = str + "I'm goroutine!"
	message <- str
	close(message)
}


func main()  {
	message := make(chan string, 3)
	go one(message)

	go two(message)
	time.Sleep(3 * time.Second)
	for v := range message {
		fmt.Println(v)
	}


	fmt.Println("end!")
}
