package web

import "net/http"

type WebHandler func(http.ResponseWriter, *http.Request, *Context) error

func (h WebHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx, err := NewContext(w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer ctx.Close()

	err = h(w, req, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
