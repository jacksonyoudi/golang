package main

import "fmt"

//func compare(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}

type inter struct {
	value int
}

func (a inter) compare(b inter) inter {
	if a.value > b.value {
		return a
	} else {
		return b
	}
}


/*
	接收者可以是任何类型
 */

type Person struct {
	name string
	age  int
}

func (p Person) Name() string {
	return p.name
}

func (p Person) SetName(name string) Person {
	p.name = name
	return p
}

func main() {

	a := inter{1}
	b := inter{2}

	fmt.Printf("%v", a.compare(b))

	p1 := Person{"youdi", 23}

	p2 := &Person{name:"hello",age:12}

	fmt.Printf("%v\n", p1)
	fmt.Printf("%v\n", p2)

}
