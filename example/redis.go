package main


import (
	"github.com/astaxie/goredis"
)

func main() {

	var client goredis.Client
	client.Addr = "127.0.0.1:6379"

	//client.Set("youdi", []byte("laji"))


	//v, _ := client.Get("youdi")
	//fmt.Println(string(v))

	f := make(map[string]interface{})
	f["youdi"] = 123
	f["hello"] = "yyy"

	err := client.Hmset("test_mset", f)
	if err != nil {
		panic(err)
	}

	//m := map[string]interface{} {
	//	"123":123,
	//	"hello": "hello"
	//}


}
