package main

import "fmt"

type person struct {
	name string
	age int
}


type stu struct {
	person //匿名字段
		   // 字段提升
	no int
}

func main()  {
	jane := person{"youdi", 12}
	fmt.Printf("%v", jane)

	ma := person{age: 13, name: "hello"}
	fmt.Printf("%v", ma)


	// 匿名字段

	student  := stu{person{"12", 12}, 123}
	fmt.Printf("%v", student.person.name)
}
