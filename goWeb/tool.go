package main

import "fmt"

/*
build 编译
intall  也是编译，与build最大的区别是编译后会将输出的文件打包成库放在pkg下面
get 获取第三方包
fmt
test  运行当前包下tests  go test go test -v

main.go
main_test.go

*/

func PrintAdd(a, b int) int {
	return a + b
}

func main()  {
	fmt.Println(PrintAdd(1,2))
}