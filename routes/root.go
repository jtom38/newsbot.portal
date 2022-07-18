package routes

import (
	"net/http"
)

func (s *HttpServer) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	err := s.templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		panic(err)
	}
}


