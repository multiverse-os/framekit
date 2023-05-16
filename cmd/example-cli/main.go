package main

import (
	"fmt"

	framekit "github.com/multiverse-os/framekit"
)

func main() {
	fmt.Println("example-cli")
	fmt.Println("===========")

	bot := framekit.NewBot("mydriasisagent")

	fmt.Printf("bot: %v\n", bot)
}
