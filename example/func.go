package main

import (
	"fmt"
)

func swap(a, b int) (int, int) {
	return b, a
}

func add(a *int) (*int) {
	*a = *a + 1
	return a
}

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	a, b := 1, 2
	a, b = swap(a, b)
	fmt.Printf("a:%d,b:%d", a, b)

	add(&a)
	fmt.Printf("a:%d,b:%d\n", a, b)

	m := []int{1, 2, 3, 4, 5}
	sum := sum(m)
	fmt.Printf("%d\n", sum)

	f := func(x, y int) int {
		return x + y
	}
	fmt.Printf("%d\n", f(1, 2))

	defer func() {
		fmt.Printf("start")
	}()

	fc := func() {
		fmt.Printf("start")
	}
	defer fc()

	//panic("123")

	for i := 1; i < 10; i++ {
		defer fmt.Printf("%d\n", i)
	}

	fmt.Printf("close")

	type sumFun func(a, b int) int

	var sum_f sumFun
	sum_f = func(a, b int) int {
		return a + b
	}
	s := sum_f(1, 5)
	fmt.Printf("%d", s)

}
