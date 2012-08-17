package web

import (
	"net/http"
)

func init() {
	http.Handle("/new-word", WebHandler(NewWordHandler))
}

func NewWordHandler(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
	switch {
	case !ctx.CanGuess():
		http.Redirect(w, req, "/leaderboard", http.StatusTemporaryRedirect)
	case ctx.State.UpdateCurrent():
		ctx.SaveSession()
		http.Redirect(w, req, "/play", http.StatusTemporaryRedirect)
	default:
		http.Redirect(w, req, "/leaderboard", http.StatusTemporaryRedirect)
	}
	return
}
