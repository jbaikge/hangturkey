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
	http.HandleFunc("/play", WebHandler(PlayHandler))
}

func PlayHandler(w http.ResponseWriter, req *http.Request, ctx *Context) {
	if err := play.Execute(w, ctx.State); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
