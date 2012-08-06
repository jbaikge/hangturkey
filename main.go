package main

import (
	"code.google.com/p/gorilla/sessions"
	"encoding/gob"
	"html/template"
	"log"
	"net/http"
)


var (
	index = template.Must(template.ParseFiles(
		"templates/_base.html",
		"templates/index.html",
	))
	store = sessions.NewCookieStore([]byte("MMMMMMM, TURKEYS"))
)

func init() {
	gob.Register(GameState{})
}

func main() {
	log.Fatal(http.ListenAndServe(":9000", nil))
}
