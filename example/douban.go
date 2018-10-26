package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
	url := "https://movie.douban.com/subject/26752088/?from=showing"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		panic("error")
	}

	cont, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(cont))
}
