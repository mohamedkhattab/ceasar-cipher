package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/index.html"))
	templates.ExecuteTemplate(w, "layout", nil)
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/public/", files)
	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
