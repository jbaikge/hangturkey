package web

import (
	"fmt"
	"net/http"
)

func init() {
	http.Handle("/guess", WebHandler(GuessHandler))
}

func GuessHandler(w http.ResponseWriter, req *http.Request, ctx *Context) error {
	fmt.Fprintf(w, "GUESS: %s", req.URL.String())
	return nil
}
