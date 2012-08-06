package main

import (
	"html/template"
	"net/http"
)

var play = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/play.html",
))

func init() {
	http.HandleFunc("/play", PlayHandler)
}

func PlayHandler(w http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "state")
	state := StateFromSession(session)
	if err := play.Execute(w, state); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
