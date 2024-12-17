package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"text/template"
)

var tmpl template.Template

type VideoInfo struct {
	Title    string   `json:"title"`
	Formats  []Format `json:"formats"`
	Playlist string   `json:"playlist"`
}

type Format struct {
	Id         string `json:"format_id"`
	Ext        string `json:"ext"`
	Resolution string `json:"resolution"`
	AudioExt   string `json:"audio_ext"`
	VideoExt   string `json:"video_ext"`
}

func main() {
	tmpl = *template.Must(template.ParseGlob("views/*.html"))

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("GET /", indexView)
	http.HandleFunc("POST /check/", optionsView)

	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexView(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "base", "")
}

func optionsView(w http.ResponseWriter, r *http.Request) {
	// assumes valid url
	url := r.FormValue("url")

	cmd := exec.Command("yt-dlp", "-j", url)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("Error running program:", err)
	}

	var videoInfo VideoInfo
	err = json.Unmarshal(output, &videoInfo)
	if err != nil {
		log.Fatal("Error unmarshaling data:", err)
	}

	resolutions := make(map[string]string)
	for _, format := range videoInfo.Formats {
		if format.VideoExt == "none" {
			continue
		}
		if _, ok := resolutions[format.Resolution]; ok == false {
			resolutions[format.Resolution] = format.Id
		}
	}

	data := map[string]map[string]string{
		"resolutions": resolutions,
	}

	tmpl.ExecuteTemplate(w, "options", data)
}
