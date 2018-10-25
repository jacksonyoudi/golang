package main

import (
	"path"
	"fmt"
)

func main()  {
	pathStr := "/youdi/name/a.txt"
	a, b := path.Split(pathStr)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("liangyoudi")
}
