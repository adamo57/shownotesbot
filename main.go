package main

import (
	"github.com/nlopes/slack"
	"os"
	"fmt"
	//"regexp"
	"regexp"
)

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// api.SetDebug(true)

	rtm := api.NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			// Replace #general with your Channel ID
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "#general"))

		case *slack.MessageEvent:
			match, err := regexp.MatchString("^<", ev.Text)
			fmt.Println(match)
			if match == true && err == nil {
				fmt.Println(ev.Text)
			}

		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:

		}
	}
}