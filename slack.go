package main

import (
	"net/http"
	"encoding/json"
	"bytes"
	"os"
)

type SlackMessage struct {
	Text string `json:"text,omitempty"`
}

func PostMessage(content SlackMessage) error {
	url := os.Getenv("SLACK_URL")

	data, err := json.Marshal(content)

	if err != nil {
		return err
	}

	_, err2 := http.Post(url, "application/json", bytes.NewReader(data))


	return err2
}
