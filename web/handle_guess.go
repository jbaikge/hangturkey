package web

import (
	"errors"
	"fmt"
	"net/http"
)

const guessURL = "/guess/"

func init() {
	http.Handle(guessURL, WebHandler(GuessHandler))
}

func GuessHandler(w http.ResponseWriter, req *http.Request, ctx *Context) error {
	letter := req.URL.String()[len(guessURL):]
	if len(letter) == 0 {
		return errors.New("No letter provided")
	}
	letter = letter[0:1]
	fmt.Fprintf(w, "GUESS:  %s\n", req.URL.String())
	fmt.Fprintf(w, "LETTER: %s\n", letter)
	return nil
}
