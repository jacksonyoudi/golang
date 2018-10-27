package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()
	timeTick := time.Tick(time.Millisecond * 100)

	for {
		select {
		case n := <-c1:
			fmt.Println("c1", n)
		case n := <-c2:
			fmt.Println("c2", n)
			//default:
			//	fmt.Println("no")
		case <-timeTick:
			fmt.Println("crontab")
		case <-time.After(1 * time.Second):
			return
		}
	}
}
