package main

import (
	//"code.google.com/p/gorilla/securecookie"
	"code.google.com/p/gorilla/sessions"
	"encoding/gob"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

type GameState struct {
	Scores      map[string]int
	CurrentWord string
}

var (
	index = template.Must(template.ParseFiles(
		"templates/_base.html",
		"templates/index.html",
	))
	store = sessions.NewCookieStore([]byte("1234567890123456"))
)

func init() {
	gob.Register(GameState{})
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	if err := index.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewWordHandler(w http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "state")
	if err != nil {
		log.Println(err)
	}
	var (
		newWord string
		state   GameState
		unseen  = []string{}
	)
	if s, ok := session.Values["state"]; ok {
		state = s.(GameState)
	} else {
		state.Scores = make(map[string]int, len(words))
		for _, w := range words {
			state.Scores[w] = 0
		}
	}
	// Gather unseen
	for w, score := range state.Scores {
		if score == 0 {
			unseen = append(unseen, w)
		}
	}
	// Pick a random unseen
	log.Printf("UNSEEN LEN %d", len(unseen))
	newWord = unseen[rand.Intn(len(unseen))]
	log.Printf("STATE      %+v", state)
	log.Printf("NEW WORD   %s", newWord)
	state.CurrentWord = newWord
	// Set session stuff
	session.Values["state"] = state
	session.Save(req, w)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/new-word", NewWordHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
