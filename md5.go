package main

import (
	"fmt"
	"crypto/md5"
)

func main()  {
	md := md5.New()
	md.Write([]byte("youdi"))
	r := md.Sum([]byte(""))
	fmt.Printf("%x\n", r)
}
