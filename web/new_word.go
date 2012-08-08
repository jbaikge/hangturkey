package web

/*
import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/new-word", NewWordHandler)
}

func NewWordHandler(w http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "state")
	state := StateFromSession(session)
	if state.UpdateCurrent() {
		session.Values["state"] = state
		session.Save(req, w)
		http.Redirect(w, req, "/play", http.StatusTemporaryRedirect)
	} else {
		log.Print("FORWARD TO LEADERBOARD")
	}
}
*/
