package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jtom38/newsbot/portal/api"
)

func (s *HttpServer) sourcesRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.ListSources)

	//r.Route("/{ID}", func(r chi.Router) {
	//	r.Get("/", s.GetSourceById)
	//})

	return r
}

type ListSourcesParam struct {
	Title string
	Items *[]api.Source
}

func (s *HttpServer) ListSources(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam{}

	items, err  := s.api.ListSources()
	if err != nil {
		panic(err)
	}
	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) GetSourceById(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam{}

	items, err  := s.api.ListSources()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}
