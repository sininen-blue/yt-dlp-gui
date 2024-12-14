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
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("GET /", indexView)
	http.HandleFunc("POST /download", downloadHandler)
	http.HandleFunc("POST /arg/{argument}", argumentHandler)

	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexView(w http.ResponseWriter, r *http.Request) {
	component := views.Base()
	component.Render(r.Context(), w)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	component := views.Base()

	services.AddArgument(r.FormValue("url"))
	err := services.Download()
	if err != nil {
		log.Fatal(err)
	}

	component.Render(r.Context(), w)
}

func argumentHandler(w http.ResponseWriter, r *http.Request) {
	component := views.Base()

	arg := r.PathValue("argument")
	log.Println(r.FormValue(arg))

	// switch statement here about all the possible arguments

	component.Render(r.Context(), w)
}
