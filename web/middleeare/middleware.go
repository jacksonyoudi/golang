package middleeare

import (
	"net/http"
	"time"
	"fmt"
	"log"
)

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

// 装饰器
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		fmt.Println(timeElapsed)
	})
}

/*func timeCount(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		next.ServeHTTP(wr, r)

		timeLong := time.Since(timeStart)
		fmt.Println(timeLong)
	})
}*/



func main() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("err")
	}
}


/*

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

任何方法实现了 ServeHTTP，即是一个合法的 http.Handler，读到这里你可能会有一些混乱，我们先来梳理一下 http 库的 Handler，HandlerFunc 和 ServeHTTP 的关系：

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
    f(w, r)
}


 */

