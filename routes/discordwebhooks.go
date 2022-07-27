package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *HttpServer) discordWebHooksRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.ArticleIndex)
	//r.Get("/list", s.ListArticles)
	//r.Route("/{ID}", func(r chi.Router) {
	//	r.Get("/", s.DisplayArticleById)
	//})

	return r
}