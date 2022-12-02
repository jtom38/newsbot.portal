package web

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/portal/api"
)

const (
	FFXIVSourceName   = "ffxiv"
	RedditSourceName  = "reddit"
	TwitchSourceName  = "twitch"
	YoutubeSourceName = "youtube"
)

var (
	pageError           = parse("templates/err.html")
	pageSettingsUpdated = parseSettings("templates/settings/posted.html")

	pageSettingIndex           = parseSettings("templates/settings/index.html")
	pageSourceUpdated          = parseSettings("templates/settings/posted.html")
	pageSettingSourcesList     = parseSettings("templates/settings/sources/list.html")
	pageSettingsNewRedditForm  = parseSettings("templates/settings/sources/new-reddit.html")
	pageSettingsNewTwitchForm  = parseSettings("templates/settings/sources/new-twitch.html")
	pageSettingsNewYouTubeForm = parseSettings("templates/settings/sources/new-youtube.html")

	pageSettingsDiscordWebhooksList = parseSettings("templates/settings/outputs/discordwebhooks/list.html")
	pageSettingsDiscordWebhooksForm = parseSettings("templates/settings/outputs/discordwebhooks/new.html")

	pageSettingsSubscriptionsList = parseSettings("templates/settings/subscriptions/list.html")
	pageSettingsSubscriptionsForm = parseSettings("templates/settings/subscriptions/form.html")
)

type SettingsRouter struct {
	_api api.CollectorApi
}

func NewSettingsRouter(api *api.CollectorApi) SettingsRouter {
	c := SettingsRouter{
		_api: *api,
	}
	return c
}

func (s *SettingsRouter) GetRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.SettingsIndex)

	r.Post("/sources/disable", s.DisableSourceById)
	r.Post("/sources/enable", s.EnableSourceById)

	r.Get("/sources/reddit", s.ListReddit)
	r.Get("/sources/reddit/new", s.NewRedditForm)
	r.Post("/sources/reddit/new", s.NewRedditPost)

	r.Get("/sources/youtube", s.ListYoutube)
	r.Get("/sources/youtube/new", s.NewYouTubeForm)
	r.Post("/sources/youtube/new", s.NewYouTubePost)

	r.Get("/sources/twitch", s.ListTwitch)
	r.Get("/sources/twitch/new", s.NewTwitchForm)
	r.Post("/sources/twitch/new", s.NewTwitchPost)

	r.Get("/sources/ffxiv", s.ListFfxiv)

	r.Get("/outputs/discord/webhooks", s.ListDiscordWebHooks)
	r.Get("/outputs/discord/webhooks/new", s.NewDiscordWebHooksForm)
	r.Post("/outputs/discord/webhooks/new", s.NewDiscordWebhookPost)

	r.Get("/subscriptions/discord/webhooks", s.ListDiscordWebHookSubscriptions)
	r.Get("/subscriptions/discord/webhooks/new", s.NewDiscordWebHookSubscriptionForm)
	r.Post("/subscriptions/discord/webhooks/new", s.NewDiscordWebHookSubscriptionPost)
	r.Post("/subscriptions/discord/webhooks/delete", s.DeleteDiscordWebHookSubscription)

	return r
}

func (s SettingsRouter) SettingsIndex(w http.ResponseWriter, r *http.Request) {
	param := TitlesParam{
		Title:    "Configuration",
		Subtitle: "It doesn't do anything on its own",
	}
	pageSettingIndex.Execute(w, param)
}

type UpdateSourceParam struct {
	Title    string
	Subtitle string
	Errors   []string
}

// /settings/sources/enable?id
func (s *SettingsRouter) EnableSourceById(w http.ResponseWriter, r *http.Request) {
	param := UpdateSourceParam{
		Title:    "Source was not enabled",
		Subtitle: "See error for details.",
	}

	err := r.ParseForm()
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageSettingsUpdated.Execute(w, param)
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		param.Errors = append(param.Errors, "The Source ID is missing")
		pageSettingsUpdated.Execute(w, param)
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageSettingsUpdated.Execute(w, param)
		return
	}

	err = s._api.Sources().Enable(uid)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageSettingsUpdated.Execute(w, param)
		return
	}

	param = UpdateSourceParam{
		Title:    "Source was enabled",
		Subtitle: "Head on back to see the change.",
	}
	if pageSettingsUpdated.Execute(w, param); err != nil {
		log.Print(err)
	}
}

// /settings/sources/disable?id
func (s *SettingsRouter) DisableSourceById(w http.ResponseWriter, r *http.Request) {
	param := UpdateSourceParam{
		Title:    "Source was not disabled",
		Subtitle: "See error for details.",
	}

	err := r.ParseForm()
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageSettingsUpdated.Execute(w, param)
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		param.Errors = append(param.Errors, "ID value was missing")
		pageSettingsUpdated.Execute(w, param)
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageSettingsUpdated.Execute(w, param)
		return
	}

	err = s._api.Sources().Disable(uid)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageSettingsUpdated.Execute(w, param)
		return
	}

	param = UpdateSourceParam{
		Title:    "Source was disabled",
		Subtitle: "Head on back to see the change",
	}
	if pageSettingsUpdated.Execute(w, param); err != nil {
		log.Print(err)
	}
}

type ListSettingsParam struct {
	Title      string
	Subtitle   string
	Items      *[]api.Source
	SourceName string
	IsError    bool
	Errors     []string
}

func (s SettingsRouter) ListReddit(w http.ResponseWriter, r *http.Request) {
	param := ListSettingsParam{
		Title:      "Known Subreddits",
		Subtitle:   "Here you can see the available sources to pick from ",
		SourceName: RedditSourceName,
	}

	items, err := s._api.Sources().ListBySource(RedditSourceName)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageSettingSourcesList.Execute(w, param)
		return
	}

	param.Items = items

	if pageSettingSourcesList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) ListYoutube(w http.ResponseWriter, r *http.Request) {
	param := ListSettingsParam{
		Title:      "Known YouTube Channels",
		Subtitle:   "Here you can see the available sources to pick from ",
		SourceName: YoutubeSourceName,
	}

	items, err := s._api.Sources().ListBySource(YoutubeSourceName)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageSettingSourcesList.Execute(w, param)
		return
	}

	param.Items = items

	if pageSettingSourcesList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) ListTwitch(w http.ResponseWriter, r *http.Request) {
	param := ListSettingsParam{
		Title:      "Known Twitch Streamers",
		Subtitle:   "Here you can see the available sources to pick from ",
		SourceName: TwitchSourceName,
	}

	items, err := s._api.Sources().ListBySource(TwitchSourceName)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageSettingSourcesList.Execute(w, param)
		return
	}

	param.Items = items

	if pageSettingSourcesList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) ListFfxiv(w http.ResponseWriter, r *http.Request) {
	param := ListSettingsParam{
		Title:      "Known Final Fantasy XIV regions",
		Subtitle:   "Here you can see the available sources to pick from",
		SourceName: FFXIVSourceName,
	}

	items, err := s._api.Sources().ListBySource(FFXIVSourceName)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageSettingSourcesList.Execute(w, param)
		return
	}

	param.Items = items

	if pageSettingSourcesList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

type NewSourceParam struct {
	Title      string
	Subtitle   string
	SourceName string
	Errors     []string
}

func (s SettingsRouter) NewRedditForm(w http.ResponseWriter, r *http.Request) {
	var err error
	param := NewSourceParam{
		Title:      "Create a new Reddit monitor",
		Subtitle:   "",
		SourceName: RedditSourceName,
	}
	if pageSettingsNewRedditForm.Execute(w, param); err != nil {
		log.Print(err)
	}
}

// This handles the data from the form and sends it to the API
func (s SettingsRouter) NewRedditPost(w http.ResponseWriter, r *http.Request) {
	param := ErrorParam{
		Title:    "Failed to add a new Reddit source",
		Subtitle: "See the error for details.",
		Code:     500,
	}
	err := r.ParseForm()
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		param.Error = "Subreddit name was missing from the form"
		pageError.Execute(w, param)
		return
	}

	uri := fmt.Sprintf("https://reddit.com/r/%v", name)
	err = s._api.Sources().NewReddit(name, uri)
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	p := TitlesParam{
		Title:    "New Reddit source was added",
		Subtitle: "Head on back to see the update",
	}

	if pageSourceUpdated.Execute(w, p); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewTwitchForm(w http.ResponseWriter, r *http.Request) {
	var err error
	param := NewSourceParam{
		Title:      "Create a new Twitch monitor",
		Subtitle:   "",
		SourceName: TwitchSourceName,
	}
	if pageSettingsNewTwitchForm.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewTwitchPost(w http.ResponseWriter, r *http.Request) {
	param := ErrorParam{
		Title:    "Failed to add a new Twitch source",
		Subtitle: "See the error for details.",
		Code:     500,
	}
	err := r.ParseForm()
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		param.Error = "Subreddit name was missing from the form"
		pageError.Execute(w, param)
		return
	}

	err = s._api.Sources().NewTwitch(name)
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	p := TitlesParam{
		Title:    "New Twitch source was added",
		Subtitle: "Head on back to see the update",
	}

	if pageSourceUpdated.Execute(w, p); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewYouTubeForm(w http.ResponseWriter, r *http.Request) {
	var err error
	param := NewSourceParam{
		Title:      "Create a new YouTube monitor",
		Subtitle:   "",
		SourceName: YoutubeSourceName,
	}
	if pageSettingsNewYouTubeForm.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewYouTubePost(w http.ResponseWriter, r *http.Request) {
	param := ErrorParam{
		Title:    "Failed to add a new YouTube source",
		Subtitle: "See the error for details.",
		Code:     500,
	}
	err := r.ParseForm()
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		param.Error = "Channel name was missing from the form."
		pageError.Execute(w, param)
		return
	}

	url := r.Form.Get("url")
	if url == "" {
		param.Error = "URL name was missing from the form."
		pageError.Execute(w, param)
		return
	}

	err = s._api.Sources().NewYouTube(name, url)
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	p := TitlesParam{
		Title:    "New YouTube source was added",
		Subtitle: "Head on back to see the update",
	}

	if pageSourceUpdated.Execute(w, p); err != nil {
		log.Print(err)
	}
}

type ListOutputDiscordWebHooks struct {
	Title    string
	Subtitle string
	Errors   []string
	Items    *[]api.Discordwebhook
}

func (s SettingsRouter) ListDiscordWebHooks(w http.ResponseWriter, r *http.Request) {
	param := ListOutputDiscordWebHooks{
		Title:    "Discord WebHooks",
		Subtitle: "Here you can see the available sources to pick from",
	}

	items, err := s._api.Outputs().DiscordWebHook().List()
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
	}

	param.Items = items

	if pageSettingsDiscordWebhooksList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewDiscordWebHooksForm(w http.ResponseWriter, r *http.Request) {
	var err error
	param := TitlesParam{
		Title: "New Discord Webhook",
	}
	if pageSettingsDiscordWebhooksForm.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewDiscordWebhookPost(w http.ResponseWriter, r *http.Request) {
	param := ErrorParam{
		Title:    "Failed to add a new YouTube source",
		Subtitle: "See the error for details.",
		Code:     500,
	}
	err := r.ParseForm()
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	server := r.Form.Get("server")
	if server == "" {
		param.Error = "Server name was missing from the form."
		pageError.Execute(w, param)
		return
	}

	url := r.Form.Get("url")
	if url == "" {
		param.Error = "URL name was missing from the form."
		pageError.Execute(w, param)
		return
	}

	channel := r.Form.Get("channel")
	if channel == "" {
		param.Error = "Channel was missing from the form."
		pageError.Execute(w, param)
		return
	}

	err = s._api.Outputs().DiscordWebHook().New(server, channel, url)
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	p := TitlesParam{
		Title:    "New Discord Web Hook",
		Subtitle: "Head on back to see the update",
	}

	if pageSourceUpdated.Execute(w, p); err != nil {
		log.Print(err)
	}
}

type ListSubscriptionsParam struct {
	Title    string
	Subtitle string
	Errors   []string
	Items    []ListSubscriptionsDetailsParam
	NewHref  string
}

type ListSubscriptionsDetailsParam struct {
	Subscription api.Subscription
	Source       api.Source
	Output       api.Discordwebhook
}

func (s SettingsRouter) ListDiscordWebHookSubscriptions(w http.ResponseWriter, r *http.Request) {
	param := ListSubscriptionsParam{
		Title:    "Subscriptions",
		Subtitle: "Links between Sources and Discord Web Hooks",
		NewHref:  "/settings/subscriptions/discord/webhooks/new",
	}

	subs, err := s._api.Subscriptions().List()
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
	}

	var details []ListSubscriptionsDetailsParam

	for index, sub := range *subs {
		sourceDetails, err := s._api.Sources().GetById(sub.Sourceid)
		if err != nil {
			msg := fmt.Sprintf("Failed to get source details on ID '%v' at index '%v'", sub.ID.String(), index)
			param.Errors = append(param.Errors, msg)
		}

		outputDetails, err := s._api.Outputs().DiscordWebHook().Get(sub.Discordwebhookid)
		if err != nil {
			msg := fmt.Sprintf("Failed to get output details on ID '%v' at index '%v'", sub.ID.String(), index)
			param.Errors = append(param.Errors, msg)
		}

		d := ListSubscriptionsDetailsParam{
			Subscription: sub,
			Source:       *sourceDetails,
			Output:       *outputDetails,
		}

		details = append(details, d)

	}

	param.Items = details

	if pageSettingsSubscriptionsList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

type NewDiscordWebHookSubscriptionFormParam struct {
	Title    string
	Subtitle string
	Errors   []string
	Outputs  []api.Discordwebhook
	Sources  []api.Source
}

func (s SettingsRouter) NewDiscordWebHookSubscriptionForm(w http.ResponseWriter, r *http.Request) {
	var err error
	param := NewDiscordWebHookSubscriptionFormParam{
		Title: "New Discord Webhook Subscription",
	}

	outputs, err := s._api.Outputs().DiscordWebHook().List()
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
	}

	sources, err := s._api.Sources().List()
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
	}

	param.Outputs = *outputs
	param.Sources = *sources

	if pageSettingsSubscriptionsForm.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewDiscordWebHookSubscriptionPost(w http.ResponseWriter, r *http.Request) {
	param := TitlesParam{
		Title:    "New Discord Web Hook",
		Subtitle: "Head on back to see the update",
	}
	err := r.ParseForm()
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageError.Execute(w, param)
		return
	}

	source := r.Form.Get("sourceName")
	if source == "" {
		msg := "Source was missing from the form."
		param.Errors = append(param.Errors, msg)
		pageError.Execute(w, param)
		return
	}

	stringSplit := strings.Split(source, "-")
	sourceRecord, err := s._api.Sources().GetBySourceAndName(stringSplit[0], stringSplit[1])
	if err != nil {
		param.Errors = append(param.Errors, "The ID value is missing.")
		pageError.Execute(w, param)
		return
	}

	DiscordWebHook := r.Form.Get("DiscordWebHook")
	if DiscordWebHook == "" {
		msg := "DiscordWebHook name was missing from the form."
		param.Errors = append(param.Errors, msg)
		pageError.Execute(w, param)
		return
	}

	outputSplit := strings.Split(DiscordWebHook, "-")
	outputRecord, err := s._api.Outputs().DiscordWebHook().GetByServerAndChannel(outputSplit[0], outputSplit[1])
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageError.Execute(w, param)
		return
	}

	err = s._api.Subscriptions().New(outputRecord[0].ID, sourceRecord.ID)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageError.Execute(w, param)
		return
	}

	if pageSourceUpdated.Execute(w, param); err != nil {
		log.Print(err)
	}
}

// This will query for a ID value to find the requested subscription to delete.
func (s SettingsRouter) DeleteDiscordWebHookSubscription(w http.ResponseWriter, r *http.Request) {
	param := UpdateSourceParam{
		Title:    "Discord Webhook Subscription",
		Subtitle: "See error for details.",
	}

	err := r.ParseForm()
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageError.Execute(w, param)
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		param.Errors = append(param.Errors, "The ID value is missing.")
		pageError.Execute(w, param)
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageError.Execute(w, param)
		return
	}

	err = s._api.Subscriptions().Delete(uid)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageError.Execute(w, param)
		return
	}

	param = UpdateSourceParam{
		Title:    "Subscription was deleted",
		Subtitle: "Head on back to see the change",
	}
	if pageSettingsUpdated.Execute(w, param); err != nil {
		log.Print(err)
	}
}
