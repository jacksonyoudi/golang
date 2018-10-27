package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	var a [10]int
	for i := 0; i < 100 ; i++ {
		go func(i int) {
			a[i]++
			runtime.Gosched() // 主动让出参数

		}(i)
	}
	time.Sleep(1 * time.Millisecond)
	fmt.Println(a)
}
