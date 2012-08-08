package main

import (
	"html/template"
	"net/http"
)

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func init() {
	http.HandleFunc("/", IndexHandler)
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	if err := index.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
