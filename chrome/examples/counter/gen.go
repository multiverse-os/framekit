//+build generate

package main

import (
	chomeui "github.com/multiverse-os/chromeui"
)

func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	chomeui.Embed("main", "assets.go", "www")
}
