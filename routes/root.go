package routes

import (
	"net/http"

	"github.com/jtom38/newsbot/portal/api"
)

func (s *HttpServer) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	err := s.templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) ArticleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	err := s.templates.ExecuteTemplate(w, "articles.index", nil)
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) GetArticleById(w http.ResponseWriter, r *http.Request) {
	items, err  := s.api.ListArticles()
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "articles.index", items)
	if err != nil {
		panic(err)
	}
}

type ListArticleParam struct {
	Items *[]api.Article
}

func (s *HttpServer) ListArticle(w http.ResponseWriter, r *http.Request) {
	param := ListArticleParam{}

	items, err  := s.api.ListArticles()
	if err != nil {
		panic(err)
	}
	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "articles.list", param)
	if err != nil {
		panic(err)
	}
}

