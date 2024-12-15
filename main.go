package main

import (
	"fmt"
	"log"
	"net/http"

	"yt-dlp-gui/services"
	"yt-dlp-gui/views"
)

var availableOptions services.Options
var chosenOptions map[string]string

func main() {
	chosenOptions = make(map[string]string)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("GET /", indexView)
	http.HandleFunc("POST /download", downloadHandler)
	http.HandleFunc("POST /updateOpts", optionsHandler)
	http.HandleFunc("POST /arg/{arg_name}", argumentHandler)

	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexView(w http.ResponseWriter, r *http.Request) {
	component := views.Base(availableOptions)
	component.Render(r.Context(), w)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	component := views.Base(availableOptions)

	services.SetArgument("url", r.FormValue("url"))
	err := services.Download()
	if err != nil {
		log.Fatal(err)
	}

	component.Render(r.Context(), w)
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	// lmao
	availableOptions, _ = services.GetOptions(url)

	component := views.Base(availableOptions)
	component.Render(r.Context(), w)
}

func argumentHandler(w http.ResponseWriter, r *http.Request) {
	component := views.Base(availableOptions)

	argName := r.PathValue("arg_name")
	argValue := r.FormValue(argName)

	switch argName {
	case "video_resolution":
		chosenOptions["resolution"] = argValue
		format := ""

		if argValue == "Best" {
			format += "bestvideo*"
		}
	default:
		log.Println("prop error here, you somehow added an invalid thing")
	}
	// switch statement here about all the possible arguments

	component.Render(r.Context(), w)
}
