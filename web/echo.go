package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

// handler 参数 wr, Response
//ResponseWriter, Request
func echo(wr http.ResponseWriter, r *http.Request)  {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte("echo error"))
		return
	}

	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}
}

/*

func handleKafka(app *ApplicationContext, w http.ResponseWriter, r *http.Request) (int, string) {
	pathParts := strings.Split(r.URL.Path[1:], "/")
	if _, ok := app.Config.Kafka[pathParts[2]]; !ok {
		return makeErrorResponse(http.StatusNotFound, "cluster not found", w, r)
	}
	if pathParts[2] == "" {
		// Allow a trailing / on requests
		return handleClusterList(app, w, r)
	}
	if (len(pathParts) == 3) || (pathParts[3] == "") {
		return handleClusterDetail(app, w, r, pathParts[2])
	}

	switch pathParts[3] {
	case "consumer":
		switch {
		case r.Method == "DELETE":
			switch {
			case (len(pathParts) == 5) || (pathParts[5] == ""):
				return handleConsumerDrop(app, w, r, pathParts[2], pathParts[4])
			default:
				return makeErrorResponse(http.StatusMethodNotAllowed, "request method not supported", w, r)
			}
		case r.Method == "GET":
			switch {
			case (len(pathParts) == 4) || (pathParts[4] == ""):
				return handleConsumerList(app, w, r, pathParts[2])
			case (len(pathParts) == 5) || (pathParts[5] == ""):
				// Consumer detail - list of consumer streams/hosts? Can be config info later
				return makeErrorResponse(http.StatusNotFound, "unknown API call", w, r)
			case pathParts[5] == "topic":
				switch {
				case (len(pathParts) == 6) || (pathParts[6] == ""):
					return handleConsumerTopicList(app, w, r, pathParts[2], pathParts[4])
				case (len(pathParts) == 7) || (pathParts[7] == ""):
					return handleConsumerTopicDetail(app, w, r, pathParts[2], pathParts[4], pathParts[6])
				}
			case pathParts[5] == "status":
				return handleConsumerStatus(app, w, r, pathParts[2], pathParts[4], false)
			case pathParts[5] == "lag":
				return handleConsumerStatus(app, w, r, pathParts[2], pathParts[4], true)
			}
		default:
			return makeErrorResponse(http.StatusMethodNotAllowed, "request method not supported", w, r)
		}
	case "topic":
		switch {
		case r.Method != "GET":
			return makeErrorResponse(http.StatusMethodNotAllowed, "request method not supported", w, r)
		case (len(pathParts) == 4) || (pathParts[4] == ""):
			return handleBrokerTopicList(app, w, r, pathParts[2])
		case (len(pathParts) == 5) || (pathParts[5] == ""):
			return handleBrokerTopicDetail(app, w, r, pathParts[2], pathParts[4])
		}
	case "offsets":
		// Reserving this endpoint to implement later
		return makeErrorResponse(http.StatusNotFound, "unknown API call", w, r)
	}

	// If we fell through, return a 404
	return makeErrorResponse(http.StatusNotFound, "unknown API call", w, r)
}
*/


func main() {
	http.HandleFunc("/", echo)





	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}


