package web

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const guessURL = "/guess/"

func init() {
	http.Handle(guessURL, WebHandler(GuessHandler))
}

func GuessHandler(w http.ResponseWriter, req *http.Request, ctx *Context) error {
	letter := strings.Replace(req.URL.String(), guessURL, "", 1)
	if len(letter) == 0 {
		return errors.New("No letter provided")
	}
	letter = letter[0:1]
	fmt.Fprintf(w, "GUESS:  %s\n", req.URL.String())
	fmt.Fprintf(w, "LETTER: %s\n", letter)
	return nil
}
