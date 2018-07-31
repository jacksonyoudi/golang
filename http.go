package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

/*
http的包
net/http


路由

 */

func main()  {
	//http.HandleFunc("/hello", func(w http.ResponseWriter, request *http.Request) {
	//	w.Write([]byte("hello, world"))
	//
	//})
	//http.ListenAndServe("127.0.0.1:8090", nil)
	url := "http://www.baidu.com"

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic("error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}