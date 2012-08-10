package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type guessMessage struct {
	Correct    bool
	Guessed    []string
	Message    string
	TotalScore int
	WordScore  int
}

const (
	guessURL       = "/guess/"
	correctScore   = 5
	incorrectScore = -1
)

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
	msg.Correct = correct
	msg.Guessed = ctx.State.GuessedLetters()
	msg.TotalScore = ctx.State.TotalScore(correctScore, incorrectScore)
	msg.WordScore = ctx.State.CurrentWordScore(correctScore, incorrectScore)

	// Resave guesses
	ctx.SaveSession()

	// Send message
	json.NewEncoder(w).Encode(msg)
	return nil
}
