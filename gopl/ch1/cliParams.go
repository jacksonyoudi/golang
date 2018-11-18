package main

import (
	"fmt"
	"os"
	"strings"
)

/*
var Args []string 切片
 */

func main() {
	//var s, sep string
	/*	for i := 1; i< len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	*/
	/*
		for _, par := range os.Args[1:] {
			s += sep + par
			sep = " "
		}

	*/

	// func Join(a []string, sep string) string {}
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])



}
