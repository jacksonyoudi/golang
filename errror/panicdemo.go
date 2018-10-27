package main

import (
	"fmt"
	"errors"
)

func recoverDemo() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("err:", err)
		} else {
			panic(r)
		}

	}()


	panic(errors.New("123"))
}


func main()  {
	recoverDemo()
}
