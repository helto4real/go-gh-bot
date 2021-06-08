package main

import (
	"fmt"

	whook "github.com/go-playground/webhooks/v6/github"
	"github.com/helto4real/ghbot/v0/ghbot"
)

func main() {
	bot := ghbot.NewGibhubBot(8001, "/events", "the_super_secret")
	bot.OnEvent(OnNewEvent, whook.IssuesEvent)
	bot.Run()
}

func OnNewEvent(gibhubEvent interface{}, err error) {
	switch gibhubEvent := gibhubEvent.(type) {
	case whook.IssuesPayload:
		fmt.Printf("%+v", gibhubEvent)
	}
}
