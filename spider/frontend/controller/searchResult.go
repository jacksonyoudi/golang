package controller

import (
	"github.com/olivere/elastic"
	"golang/spider/frontend/view"
	"net/http"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}



//
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	q := strings.TrimSpace(r.FormValue("q"))
	from ,err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		panic(err)
	}


	w.Write(h)


}