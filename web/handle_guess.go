package web

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

type guessMessage struct {
	Letter  string
	Message string
	Score   int
}

const guessURL = "/guess/"

func init() {
	http.Handle(guessURL, WebHandler(GuessHandler))
}

func GuessHandler(w http.ResponseWriter, req *http.Request, ctx *Context) error {
	l := req.URL.String()[len(guessURL):]
	if len(l) == 0 {
		return errors.New("No letter provided")
	}
	l = strings.ToLower(l[:1])
	current := ctx.State.CurrentWord
	guesses := ctx.State.Guesses[current]
	msg := &guessMessage{}

	// See if the word was already guessed
	if strings.Contains(current, l) {
		if strings.Contains(guesses.Correct, l) {
			msg.Message = "You have already guessed " + l
		} else {
			guesses.Correct += l
			msg.Letter = l
			msg.Score = 5
		}
	} else {
		if strings.Contains(guesses.Incorrect, l) {
			msg.Message = "You have already guessed " + l
		} else {
			guesses.Incorrect += l
			msg.Letter = l
			msg.Score = -1
		}
	}

	// Resave guesses
	ctx.State.Guesses[current] = guesses
	ctx.SaveSession()

	// Send message
	json.NewEncoder(w).Encode(msg)
	return nil
}
