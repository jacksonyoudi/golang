package main

import (
	"os"
	"github.com/anaskhan96/soup"
	"strings"
	"fmt"
)

type Book struct {
}

func GetBookUrls(url string) ([]string, string) {
	var urls []string
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("ul", "class", "hanghang-list").FindAll("a")
	for _, link := range links {
		href := link.Attrs()["href"]
		if ! strings.HasPrefix(href, "/index.php/bookInfo/") {
			continue
		}
		url := "http://www.ireadweek.com" + href
		urls = append(urls, url)
	}
	return urls, ""

}

func GetBookDetail(url string) {
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	titleNode := doc.Find("div", "class", "hanghang-za-title")
	title := titleNode.Text()
	fmt.Println("title:" + title)

	imgNode := doc.Find("div", "class", "hanghang-shu-content-img").Find("img")
	img := imgNode.Attrs()["src"]
	fmt.Println("img:" + img)


	downloadNode := doc.Find("a", "class", "downloads")
	download := downloadNode.Attrs()["href"]
	fmt.Println("download:" + download)



	writerNode := doc.Find("div", "class", "hanghang-za-content")
	writer := writerNode.Text()
	fmt.Println("writer:" + writer)


	links := doc.Find("div", "class", "hanghang-shu-content-font").FindAll("p")

	for _, link := range links {
		fmt.Println(link.Text())
	}

	lists := doc.FindAll("div", "class", "hanghang-za-content")[1].FindAll("p")

	for _, list := range lists {
		fmt.Println(list.Text())
	}


}

func main() {

	start_url := `http://www.ireadweek.com/index.php/index/2.html`

	urls, _ := GetBookUrls(start_url)
	for _, v := range urls {
		fmt.Println(v)
		GetBookUrls(v)
	}

}
