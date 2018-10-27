package main

import (
	"fmt"
	"time"
)

func main()  {
	for i := 0; i < 100 ; i++ {
		go func(a int) {
			fmt.Printf("call: %d \n", i)
		}(i)
	}
	time.Sleep(1 * time.Millisecond)
}
