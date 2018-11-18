package main

import (
	"golang/spider/frontend/view"
	"net/http"
)

func main()  {
	http.Handle("/search", view.CreateearchResultHandler("index.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}

