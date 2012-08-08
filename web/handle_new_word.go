package web

import (
	"log"
	"net/http"
)

func init() {
	http.Handle("/new-word", WebHandler(NewWordHandler))
}

func NewWordHandler(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
	if ctx.State.UpdateCurrent() {
		ctx.Session.Values["state"] = ctx.State
		ctx.Session.Save(req, w)
		http.Redirect(w, req, "/play", http.StatusTemporaryRedirect)
	} else {
		log.Print("FORWARD TO LEADERBOARD")
	}
	return
}
