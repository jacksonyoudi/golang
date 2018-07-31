package main

import (
	"encoding/json"
	"fmt"
)

/*
json ending/json
ffjson


func Marshal()


md5
 */

 type student struct {
 	Name string `json:"student_name"` //重命名
 	Age int
 }


func main()  {
	x := [5]int{1,2,3,4,5}
	s, err := json.Marshal(x)
	if err != nil {
		panic("err")
	}
	fmt.Println(string(s))

	m := make(map[int]string)
	m[1] = "string"
	m[2] = "you"
	m[3] = "ng"
	m[4] = "ing"

	n, ok := json.Marshal(m)
	if ok != nil {
		panic("err")
	}
	fmt.Println(string(n))


	stu := student{"youdi", 12}

	j, _ := json.Marshal(stu)
	fmt.Println(string(j))


	//解码
	var s3 interface{}
	json.Unmarshal([]byte(j), &s3)
	fmt.Printf("%v", s3)

}


