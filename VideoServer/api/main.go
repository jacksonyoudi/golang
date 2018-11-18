package main

import (
	"golang/VideoServer/api/conf"
	"net/http"
)


func main() {
	r := conf.RegisterHandlers()
	http.ListenAndServe(":8080", r)
}
