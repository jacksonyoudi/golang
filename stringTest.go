package main

import "fmt"

func main()  {

	var a string
	a = "hello"
	for _, c := range a {
		fmt.Println(string(c))
	}


}
