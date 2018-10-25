package main
import "C"

func main()  {
	println("hello Cgo!")
	C.puts(C.CString("Hello, World\n"))
}
