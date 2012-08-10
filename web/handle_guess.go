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
	current := ctx.State.CurrentWord
	g := ctx.State.Guesses[current]
	msg := &guessMessage{}

	// See if the word was already guessed
	switch {
	case strings.Contains(g.Correct, l) || strings.Contains(g.Incorrect, l):
		msg.Message = "You have already guessed " + l
	case strings.Contains(current, l):
		g.Correct += l
		msg.Score = 5
	case !strings.Contains(current, l):
		g.Incorrect += l
		msg.Score = -1
	}
	msg.Guessed = ctx.State.GuessedLetters()

	// Resave guesses
	ctx.State.Guesses[current] = g
	ctx.SaveSession()

	// Send message
	json.NewEncoder(w).Encode(msg)
	return nil
}
