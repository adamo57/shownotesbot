package main

import (
	"github.com/nlopes/slack"
	"os"
	"regexp"
	"log"
	"net/http"
)

//IsPodcastRunning is the conditional that tells
// the app whether or not the podcast is running.
var IsPodcastRunning = false

func main() {
	port := os.Getenv("PORT")
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	rtm := api.NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:

			checkPodcastStatus(api, ev.Text)

			match, err := regexp.MatchString("^<", ev.Text)
			if err != nil {
				log.Fatal("error processing string")
			}

			if IsPodcastRunning == true && match == true {
				api.PostMessage("general", ev.Text,
					slack.NewPostMessageParameters())
			}
		}
	}
}

func checkPodcastStatus(api *slack.Client, podcastStatusText string) {
	if podcastStatusText == "start podcast" {
		api.PostMessage("general",
			"I am now listening, type _stop podcast_ " +
			"to tell me to stop listening and give you a list",
				slack.NewPostMessageParameters())
		IsPodcastRunning = true
	} else if podcastStatusText == "stop podcast"{
		api.PostMessage("general", "I am now *not* listening",
			slack.NewPostMessageParameters())
		IsPodcastRunning = false
	}
}