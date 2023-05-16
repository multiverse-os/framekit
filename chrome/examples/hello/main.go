package main

import (
	"log"

	chrome "github.com/multiverse-os/framekit/chrome"
)

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := chrome.New("https://twitch.tv/mydriasisagent", "--automation=false", 480, 320)
	if err != nil {
		log.Fatal(err)
	}

	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
