package main

import (
	"strconv"
	"time"
	"fmt"
	"os"
)

func sample(ch chan string)  {
	for i := 0; i< 5; i++ {
		ch <- "i'm sample1 num:" + strconv.Itoa(i)
	}

	time.Sleep(4 * time.Second)

}


func sample1(ch chan int)  {
	for i := 0; i< 19; i++ {
		ch <- i
	}
	time.Sleep(3 * time.Second)

}


func main()  {
	chan1 := make(chan string, 3)
	chan2 := make(chan int, 5)
	for i := 0; i< 10 ; i++ {
		go sample(chan1)
		go sample1(chan2)
	}
	for  {
		select {
		case str, ok := <-chan1:
			if !ok {
				fmt.Println("chan1,failure")
			}
			fmt.Println(str)
		case one, err := <-chan2:
			if !err {
				fmt.Println("chan2,failure")
			}
			fmt.Println(one)
		case <- time.After(2 * time.Second): // 一直读取不到数据，timeout 就会有数据，就会执行
			fmt.Println("come to timeout")
			os.Exit(1)
		}
	}

}
