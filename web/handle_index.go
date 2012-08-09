package web

import (
	"net/http"
)

var index = parseTemplates(
	"web/templates/_base.html",
	"web/templates/index.html",
)

func init() {
	http.Handle("/", WebHandler(IndexHandler))
}

func IndexHandler(w http.ResponseWriter, req *http.Request, ctx *Context) error {
	return index.Execute(w, nil)
}
