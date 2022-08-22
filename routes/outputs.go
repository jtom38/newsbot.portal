package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/portal/api"
)

func (s *HttpServer) outputsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.ArticleIndex)
	r.Get("/discord/webhooks", s.OutputDiscordIndex)
	s.Router.Route("/discord/webhooks/{ID}", func(r chi.Router) {
		r.Post("/delete", s.DeleteDiscordWebHookById)
		r.Post("/disable", s.DisableDiscordWebHookById)
		r.Post("/enable", s.EnableDiscordWebHookById)
	})

	return r
}

type ListDiscordWebHooksParam struct {
	Title string
	Subtitle string
	Items []ListDiscordWebHooksDetailsParam
}

type ListDiscordWebHooksDetailsParam struct {
	Disabled bool
	Enabled bool
	Item api.Discordwebhook
}

func (s *HttpServer) OutputDiscordIndex(w http.ResponseWriter, r *http.Request) {
	var details []ListDiscordWebHooksDetailsParam
	param := ListDiscordWebHooksParam{
		Title:    "Discord Web Hooks",
		Subtitle: "Here are the known Discord web hooks.",
	}

	items, err := s.api.Outputs.ListDiscordWebHooks()
	if err != nil {
		panic(err)
	}

	for _, item := range *items {
		var i ListDiscordWebHooksDetailsParam

		if !item.Enabled {
			i.Disabled = true
			i.Enabled = false
		} else {
			i.Disabled = false
			i.Enabled = true
		}

		i.Item = item

		details = append(details, i)
	}
	param.Items = details

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "settings.outputs.discordwebhooks.index", param)
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) DeleteDiscordWebHookById(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Form Error",
			Code:  500,
			Error: err,
		})
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	err = s.api.Sources.Delete(uid)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to delete the Source",
			Code:  500,
			Error: err,
		})
		return
	}
}

func (s *HttpServer) EnableDiscordWebHookById(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Form Error",
			Code:  500,
			Error: err,
		})
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	err = s.api.Sources.Enable(uid)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to delete the Source",
			Code:  500,
			Error: err,
		})
		return
	}

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.posted", HttpParam{
		Title:    "Source Enabled",
		Subtitle: "It will be reviewed on the next collection",
	})
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) DisableDiscordWebHookById(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Form Error",
			Code:  500,
			Error: err,
		})
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	err = s.api.Sources.Disable(uid)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to delete the Source",
			Code:  500,
			Error: err,
		})
		return
	}

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.posted", HttpParam{
		Title:    "Source Disabled",
		Subtitle: "It will be skipped on the next collection",
	})
	if err != nil {
		panic(err)
	}
}