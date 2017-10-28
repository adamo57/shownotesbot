package main

import (
	"github.com/nlopes/slack"
	"os"
	"fmt"
	"regexp"
)

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	rtm := api.NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			match, err := regexp.MatchString("^<", ev.Text)

			if match == true && err == nil {
				fmt.Println(ev.Text)
			}
		}
	}
}