package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)


func determineEncoding(r *bufio.Reader) encoding.Encoding {
	i, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(i, "")
	return e
}


var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error)  {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	fmt.Println(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)

	utf8Resp := transform.NewReader(bodyReader, e.NewDecoder())

	return  ioutil.ReadAll(utf8Resp)
}
