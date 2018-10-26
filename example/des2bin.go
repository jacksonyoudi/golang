package main

import (
	"fmt"
)

func mod(n int) (s, m int) {
	s = n / 2
	m = n % 2
	return s, m
}

func dec2bin(des int) []int {
	var m []int
	for {
		x, y := mod(des)
		m = append(m, y)
		if x == 0 {
			break
		} else {
			des = x
		}
	}

	return m


}

func main()  {

	a := dec2bin(4)
	fmt.Println(a)
}