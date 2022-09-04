package routes

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/portal/api"
)

func (s *HttpServer) outputsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/discord/webhooks", s.OutputDiscordIndex)
	r.Get("/discord/webhooks/new", s.NewDiscordWebhookDisplay)
	r.Post("/discord/webhooks/new", s.NewDiscordWebhookFormPost)
	r.Route("/discord/webhooks/{ID}", func(r chi.Router) {
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
		Subtitle: "Here are the known Discord Webhooks.",
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


/* Discord Web Hooks */
func (s *HttpServer) DeleteDiscordWebHookById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	if id == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Source ID",
			Code:  500,
			Error: errors.New("ID was not found"),
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

	err = s.api.Outputs.DeleteDiscordWebHook(uid)
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
	id := chi.URLParam(r, "ID")
	if id == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Source ID",
			Code:  500,
			Error: errors.New("ID was not found"),
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

	err = s.api.Outputs.EnableDiscordWebHook(uid)
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
	id := chi.URLParam(r, "ID")
	if id == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Source ID",
			Code:  500,
			Error: errors.New("ID was not found"),
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

	err = s.api.Outputs.DisableDiscordWebHook(uid)
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

func (s *HttpServer) NewDiscordWebhookDisplay(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam {
		Title: "New Discord Webhook",
		Subtitle: "Where should messages go?",
	}

	w.Header().Add("Content Type", "text/html")
	err := s.templates.ExecuteTemplate(w, "settings.outputs.discord.webhooks.new", param)
	if err != nil {
		panic(err)
	}
}

// This validates the infomation sent from the form and passes it to the API.
func (s *HttpServer) NewDiscordWebhookFormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Form Error",
			Code:  500,
			Error: err,
		})
		return
	}

	server := r.Form.Get("server")
	if server == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Name value",
			Code:  500,
			Error: err,
		})
		return
	}

	channel := r.Form.Get("channel")
	if channel == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Name value",
			Code:  500,
			Error: err,
		})
		return
	}

	url := r.Form.Get("url")
	if url == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Name value",
			Code:  500,
			Error: err,
		})
		return
	}

	err = s.api.Outputs.NewDiscordWebhook(server, channel, url)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to add new Discord Webhook",
			Code:  500,
			Error: err,
		})
		return
	}

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.posted", HttpParam{
		Title:    "Source Added",
		Subtitle: "It will be reviewed on the next collection",
	})
	if err != nil {
		panic(err)
	}
}