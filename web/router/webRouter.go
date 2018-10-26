/*package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
)

func main() {
	r := httprouter.New()
	r.PUT()

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, c interface{}) {
		log.Printf("Recovering from panic, Reason: %#v", c.(error))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(c.(error).Error()))
	}



}
*/

//package main
//
//import (
//	"fmt"
//	"github.com/julienschmidt/httprouter"
//	"net/http"
//	"log"
//)
//
//func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Fprint(w, "Welcome!\n")
//}
//
//func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
//}
//
//func main() {
//	router := httprouter.New()
//	router.GET("/", Index)
//	router.GET("/hello/:name", Hello)
//
//	log.Fatal(http.ListenAndServe(":8080", router))
//}

package router

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("hell" + ps.ByName("hello")))
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("hell0" + ps.ByName("name")))
}

func main() {
	// router就是一个handler函数，http库上扩展的
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	http.ListenAndServe(":8080", router)
}
