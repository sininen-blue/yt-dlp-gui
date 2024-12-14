package main

import (
	"fmt"
	"log"
	"net/http"

	"yt-dlp-gui/services"
	"yt-dlp-gui/views"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexView)

	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexView(w http.ResponseWriter, r *http.Request) {
	component := views.Base()
	switch r.Method {
	case http.MethodGet:
		log.Println("in get")
		component.Render(r.Context(), w)
	case http.MethodPost:
		log.Println("in post")
		services.AddArgument(r.FormValue("url"))
		err := services.Download()
		if err != nil {
			log.Fatal(err)
		}

		component.Render(r.Context(), w)
	}
}
