package events

import (
	"fmt"

	gdk "github.com/mattn/go-gtk/gdk"
)

func ButtonHandler(event chan interface{}) {
	for {
		e := <-event
		switch ev := e.(type) {
		case *gdk.EventButton:
			fmt.Printf("[DEBUG] button-press-event: %d %d\n",
				int(ev.X), int(ev.Y))
			if ev.Button == 2 {
				fmt.Printf("[DEBUG] Open link on new tab")
			}
			// TODO : Button 1 + Ctrl
			break
		default:
			fmt.Printf("[DEBUG] event: %v\n", ev)
		}
	}
}
