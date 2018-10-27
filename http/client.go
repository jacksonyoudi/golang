package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main()  {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)

	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect:", req)
			return nil
	},
	}
	
	
	resp, er := client.Do(request)
	if er != nil {
		panic(er)
	}
	defer resp.Body.Close()

	httputil.DumpResponse(resp, true)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s", s)
}