package view

import (
	"github.com/olivere/elastic"
	"golang/spider/frontend/controller"
	"golang/spider/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(
	filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

func (s SearchResultView) Render(
	w io.Writer,data model.SearchResult) error {
		return s.template.Execute(w, data)
}

func CreateearchResultHandler(template string) controller.SearchResultHandler {
	client ,err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return controller.SearchResultHandler{
		CreateSearchResultView(template),
		client,
	}

}