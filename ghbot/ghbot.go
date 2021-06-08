package ghbot

import (
	"fmt"
	"net/http"

	"github.com/go-playground/webhooks/v6/github"
)

type GitHubCallBackFunction func(event interface{}, err error)

type GithubBot struct {
	port   int
	path   string
	secret string
	hook   *github.Webhook
}

func NewGibhubBot(port int, webhookPath string, secret string) *GithubBot {
	hook, _ := github.New(github.Options.Secret(secret))
	return &GithubBot{
		path:   webhookPath,
		secret: secret,
		hook:   hook,
		port:   port,
	}
}

func (gb *GithubBot) OnEvent(f GitHubCallBackFunction, events ...github.Event) {
	http.HandleFunc(gb.path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := gb.hook.Parse(r, events...)

		go f(payload, err)
	})
}

func (gb *GithubBot) Run() {
	http.ListenAndServe(fmt.Sprintf(":%d", gb.port), nil)
}
