package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/portal/api"
)

func (s *HttpServer) articlesRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.ArticleIndex)
	r.Get("/list", s.ListArticles)

	r.Route("/{ID}", func(r chi.Router) {
		r.Get("/", s.DisplayArticleById)
	})
	
	//r.Get("/source", s.ListArticlesBySourceId)

	return r
}

type DisplayArticleParams struct {
	Title string
	Article *api.Article
	Source *api.Source
	IsImage bool
}

func (s *HttpServer) DisplayArticleById(w http.ResponseWriter, r *http.Request) {
	param := DisplayArticleParams{}

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	
	article, err  := s.api.GetArticle(uuid)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	param.Article = article
	param.Title = article.Title

	source, err := s.api.GetSourceById(article.Sourceid)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	param.Source = source

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "articles.display", param)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

type ListArticleParam struct {
	Title string
	Items *[]api.Article
}

// This returns the newest 50 articles to the user
func (s *HttpServer) ListArticles(w http.ResponseWriter, r *http.Request) {
	param := ListArticleParam{}

	items, err  := s.api.ListArticles()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "articles.list", param)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
