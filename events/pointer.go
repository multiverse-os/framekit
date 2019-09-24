package events

import (
	"fmt"

	gdk "github.com/mattn/go-gtk/gdk"
)

// MotionHandler handle events from mouse
func PointerHandler(event chan interface{}) {
	for {
		e := <-event
		switch ev := e.(type) {
		case *gdk.EventMotion:
			fmt.Println("[DEBUG] pointer-notify-event:",
				int(ev.X), int(ev.Y))
			break
		default:
			fmt.Printf("[DEBUG] event: %v\n", ev)
		}
	}
}
