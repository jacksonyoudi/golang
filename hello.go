package main

import "fmt"

func main() {
	//var x int
	//x = 1
	//fmt.Print(x)

	y, z := "youdi", "name"
	fmt.Print(y, z)

	//var x complex64 = 6 + 2i
	//fmt.Print(x * x)

	var x [10]int
	x[0] = 10
	x[9] = 20

	fmt.Printf("%k", x)


	v := [10]int{1,2,3,4,5}
	fmt.Printf("%v", v)

	m := len(v)
	fmt.Printf("%d", m)

	s := make([]int, 20, 30)
	fmt.Print(len(s))
	fmt.Print(cap(s))

	s = append(s, 1,2,3,4)
	fmt.Print(cap(s))



}
