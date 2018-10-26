package main

import "fmt"

func main()  {
	var stu map[int]float32
	stu = make(map[int]float32)
	stu[12] = 123.4
	fmt.Print(stu)

	//panic: assignment to entry in nil map
}
