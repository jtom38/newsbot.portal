package routes

import (
	"context"
	"html/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jtom38/newsbot/portal/api"
)

type HttpParam struct {
	Title string
}

type HttpServer struct {
	Router *chi.Mux

	ctx *context.Context

	// Links to the class to interface with the API
	api *api.ApiClient

	// Contains where to find all Templates
	templates *template.Template
}

func NewServer(ctx *context.Context, ApiEndpoint string) *HttpServer {
	s := HttpServer{
		ctx: ctx,
	}

	api := api.New(ApiEndpoint)
	s.api = api

	tmpl := NewTmpl()
	err := tmpl.Load("./web/templates", ".html")
	if err != nil {
		panic(err)
	}
	s.templates = tmpl.Template

	s.Router = chi.NewRouter()
	s.MountMiddleware()
	s.MountRoutes()

	return &s
}

func (s *HttpServer) MountMiddleware() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
}

func (s *HttpServer) MountRoutes() {
	s.Router.Get("/", s.Index)
	s.Router.Get("/articles", s.ArticleIndex)
	s.Router.Get("/articles/list", s.ListArticle)
}
