package web

import (
	"net/http"
)

func init() {
	http.Handle("/play", WebHandler(PlayHandler))
}

func PlayHandler(w http.ResponseWriter, req *http.Request, ctx *Context) error {
	var play = parseTemplates(
		"web/templates/_base.html",
		"web/templates/play.html",
	)
	return play.Execute(w, ctx.State)
}
