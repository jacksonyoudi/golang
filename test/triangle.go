package test

import "fmt"

func Triangle(a, b, c int) bool {
	switch {
	case a + b <= c:
		return false
	case a + c <= b:
		return false
	case b + c <= a:
		return false
	default:
		return true
	}

}




func main()  {
	var a, b , c = 3, 4, 5
	fmt.Println(Triangle(a, b, c))
}
