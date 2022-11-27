package web

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jtom38/newsbot/portal/api"
)

var (
	index     = parse("templates/index.html")
	errorPage = parse("templates/err.html")
)

type HttpServer struct {
	Router *chi.Mux

	// Links to the class to interface with the API
	//api *api.ApiClient
	api api.CollectorApi

	ctx context.Context
}

func NewServer(ctx context.Context, ApiEndpoint string) *HttpServer {
	s := HttpServer{
		ctx: ctx,
	}

	api := api.New(ApiEndpoint)
	s.api = api

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

	s.Router.Mount("/articles", s.articlesRouter())

	settings := NewSettingsRouter(&s.api)
	s.Router.Mount("/settings", settings.GetRouter())
	
	//s.Router.Mount("/settings/sources", s.sourcesRouter())
	//s.Router.Mount("/settings/outputs", s.outputsRouter())
}

func (s *HttpServer) Index(w http.ResponseWriter, r *http.Request) {
	param := TitlesParam{
		Title:    "Welcome",
		Subtitle: "Your news destination",
	}
	index.Execute(w, param)
}
