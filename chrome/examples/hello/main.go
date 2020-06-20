package main

import (
	"log"
	"net/url"

	chromeui "github.com/multiverse-os/chromeui"
)

func main() {
	// Create UI with basic HTML passed via data URI

	ui, err := chromeui.New("data:text/html,"+url.PathEscape(`
	<html>
		<head>
			<title>Hello</title>
			<script>alert(1)</script>
		</head>
		<body><h1>Hello, world!</h1></body>
	</html>
	`), "", 480, 320)

	if err != nil {
		log.Fatal(err)
	}

	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
