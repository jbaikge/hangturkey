package web

import (
	"net/http"
	"strings"

	"github.com/jbaikge/hangturkey/app"
)

type playContext struct {
	Alphabet []string
	State    app.GameState
}

var alphabet []string

func init() {
	http.Handle("/play", WebHandler(PlayHandler))
	alphabet = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

}

func PlayHandler(w http.ResponseWriter, req *http.Request, ctx *Context) error {
	var play = parseTemplates(
		"web/templates/_base.html",
		"web/templates/play.html",
	)
	c := playContext{
		State:    ctx.State,
		Alphabet: alphabet,
	}
	return play.Execute(w, c)
}
