package main

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/multiverse-os/framekit/webkit2"

	"log"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Webkit2 Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new  user content manager so we can alter the web view.
	ucm, err := webkit.UserContentManagerNew()
	if err != nil {
		log.Fatal("Unable to create user content manager:", err)
	}

	wv, err := webkit.WebViewNewWithUserContentManager(ucm)
	if err != nil {
		log.Fatal("Unable to create webview:", err)
	}

	// Add the label to the window.
	win.Add(wv)

	// Set the default window size.
	win.SetDefaultSize(1000, 800)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Add a callback to the webview for when the load operation changes.
	wv.Connect("load-changed", loadChanged)

	// Load a URI.
	wv.LoadURI("https://www.google.co.nz/#q=example.com")

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

func loadChanged(webView *webkit.WebView, event webkit.LoadEvent) {
	if event == webkit.LOAD_COMMITTED {
		// Print the final URI for the webpage (the URI after all redirects are done).
		log.Println("Final URI for page load:", webView.GetURI())
	}
}
