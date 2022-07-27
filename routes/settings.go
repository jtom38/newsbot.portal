package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *HttpServer) settingsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.SettingsIndex)

	//r.Route("/{ID}", func(r chi.Router) {
	//	r.Get("/", s.GetSourceById)
	//})

	return r
}

type SettingsIndexParam struct {
	Title string
	Subtitle string
}

func (s *HttpServer) SettingsIndex(w http.ResponseWriter, r *http.Request) {
	param := SettingsIndexParam{
		Title: "Config your news!",
		Subtitle: "Select your config!",
	}

	w.Header().Add("Content Type", "text/html")
	err := s.templates.ExecuteTemplate(w, "settings.index", param)
	if err != nil {
		panic(err)
	}
}
