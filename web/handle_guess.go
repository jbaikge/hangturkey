package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type guessMessage struct {
	Complete   bool
	Correct    bool
	Guessed    []string
	Message    string
	TotalScore int
	WordScore  int
	WrongCount int
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
	if err != nil {
		msg.Message = err.Error()
	}
	msg.Complete = ctx.State.CurrentComplete()
	msg.Correct = correct
	msg.Guessed = ctx.State.GuessedLetters()
	msg.TotalScore = ctx.TotalScore()
	msg.WordScore = ctx.WordScore()
	msg.WrongCount = ctx.State.IncorrectGuesses()

	// Resave guesses
	ctx.SaveSession()

	// Send message
	json.NewEncoder(w).Encode(msg)
	return nil
}
