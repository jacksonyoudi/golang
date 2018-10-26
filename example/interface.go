package main

import (
	"fmt"
)

type Animal interface {
	Fly()
	Run()
}



//只要类中实现了所有的接口方法，就可以赋值给接口， 可以把多的赋值给少的


//类型查询





//任何对象都满足一个空接口

type Animal2 interface {
	Fly()
}


type Dog struct {

}

func (d Dog) Fly()  {
	fmt.Printf("Fly.......")
}

func (d Dog) Run() {
	fmt.Printf("Run........")
}

func main()  {
	var animal Animal
	var animal2 Animal2

	animal2 = animal
	animal2.Fly()

	//接收任何类型
	var v1 interface{} = 1
	var v2 interface{} = "string"


	fmt.Printf("%v", v1)
	fmt.Printf("%v", v2)



	if v, ok := v1.(int); ok {
		fmt.Println(v, ok)
	} else {

	}

	switch v1.(type) {
	case int:
	case float64:
	case string:

	}

}