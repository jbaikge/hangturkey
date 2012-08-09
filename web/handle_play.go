package web

import (
	"net/http"
)

var play = parseTemplates(
	"web/templates/_base.html",
	"web/templates/play.html",
)

func init() {
	http.Handle("/play", WebHandler(PlayHandler))
}

func PlayHandler(w http.ResponseWriter, req *http.Request, ctx *Context) error {
	return play.Execute(w, ctx.State)
}
