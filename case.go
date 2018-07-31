package main

import "fmt"

func main() {

	var m, sum int
	m = 100

	for i := 1; i < m; i++ {
		sum += i
		fmt.Print(sum)
		fmt.Printf("\n")
	}

	arr := [5]int{1,2,3,4,5}
	for k,v := range arr {
		fmt.Printf("%d", k)
		fmt.Printf("%d\n", v)
	}

	ma :=  make(map[string]string)
	ma["1234"] = "1234"
	ma["23"] = "345"

	for k, v := range ma {
		fmt.Printf("%s", k)
		fmt.Printf("%s\n", v)
	}


	st := "string"
	for _,i := range st {
		fmt.Printf("%c", i)
	}


}
