package main

import (
	"fmt"
	"regexp"
)

const text = "my email is liangchangyoujackson@gmail.com aa@qq.com"

func main()  {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	for _, row := range match {
		fmt.Println(row)
	}
}

