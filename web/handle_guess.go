package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type guessMessage struct {
	Guessed []string
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
	msg := &guessMessage{}

	correct, err := ctx.State.Guess(l)
	switch {
	case err != nil:
		msg.Message = err.Error()
	case correct:
		msg.Score = 5
	case !correct:
		msg.Score = -1
	}
	msg.Guessed = ctx.State.GuessedLetters()

	// Resave guesses
	ctx.SaveSession()

	// Send message
	json.NewEncoder(w).Encode(msg)
	return nil
}
