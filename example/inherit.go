package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) NameAndAge() (string, int) {
	return p.name, p.age
}

type Student struct {
	Person
	speciaility string
}

func (s Student) GetSpeciaility() string {
	return s.speciaility
}









func main() {
	stu := new(Student)
	stu.name = "123"
	stu.age = 24
	stu.speciaility = "4555"

	age, name := stu.NameAndAge()
	fmt.Printf("%d\n%s", age, name)
}



/*
在go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口。
接口赋值

将对象实例赋值给接口
将一个接口赋值给另一个接口

非侵入式接口



*/


