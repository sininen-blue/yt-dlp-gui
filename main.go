package main

import (
	"context"
	// "log"
	// "net/http"
	"os"

	"yt-dlp-gui/views"
)

func main() {
	component := views.Hello("John")
	component.Render(context.Background(), os.Stdout)
}

// func index(w http.ResponseWriter, r *http.Request) {
// }
