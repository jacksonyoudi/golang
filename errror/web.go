package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"log"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func errWarapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			log.Fatal(err)
			switch  {
			case os.IsNotExist(err):
				http.Error(w, "File not exit", http.StatusNotFound)
			default:
				http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			}
		}
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	w.Write(all)
	return nil
}

func main() {
	http.HandleFunc("/", errWarapper(listHandler))
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}
