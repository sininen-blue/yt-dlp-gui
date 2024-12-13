package main

import (
	"fmt"
	"log"
	"net/http"

	"yt-dlp-gui/views"

	"github.com/a-h/templ"
)

func main() {
	component := views.Base()

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
