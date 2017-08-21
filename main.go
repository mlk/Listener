package main

import (
	"os"
	"fmt"
	"bufio"
)

func userJoined(user string) {
	PostMessage(SlackMessage{Text: user + " joined."})
}

func userLeft(user string) {
	PostMessage(SlackMessage{Text: user + " left."})
}

func main() {
	if os.Getenv("SLACK_URL") == "" {
		fmt.Print("Environment var SLACK_URL must be set.")
		return
	}

	PostMessage(SlackMessage{Text: "Factorio server has started up!"})

	factorio := CreateFactorio(userJoined, userLeft)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := Parse(scanner.Text())
		if line.EventType == JOINED {
			factorio.AddUser(line.User)
		} else if line.EventType == LEFT {
			factorio.DeleteUser(line.User)
		}
		fmt.Println(line.Raw)
	}

	PostMessage(SlackMessage{Text: "Factorio server has shut down :wave:"})
}
