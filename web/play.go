package web

import (
	"html/template"
	"net/http"
	"strings"
)

var play = template.Must(template.New("_base.html").Funcs(template.FuncMap{
	"TrimSpace": strings.TrimSpace,
}).ParseFiles(
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
