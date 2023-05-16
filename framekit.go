package framekit

import (
	"fmt"

	chrome "github.com/multiverse-os/framekit/chrome"
)

type Bot struct {
	Username string
	Browser  *chrome.Browser
}

func NewBot(username string) *Bot {
	return &Bot{
		Username: username,
		Browser:  &chrome.Browser{},
	}
}

func TestFunction() {
	fmt.Println("vim-go")
}
