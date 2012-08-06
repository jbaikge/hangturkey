package main

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/new-word", NewWordHandler)
}

func NewWordHandler(w http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "state")
	var state GameState
	if s, ok := session.Values["state"]; ok {
		state = s.(GameState)
	} else {
		state = GameState{}
	}
	if state.UpdateCurrent() {
		session.Values["state"] = state
		session.Save(req, w)
	} else {
		log.Print("FORWARD TO LEADERBOARD")
	}
}
