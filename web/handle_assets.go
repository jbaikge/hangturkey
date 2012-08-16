package web

import (
	"net/http"
)

func init() {
	http.Handle("/images/", http.FileServer(http.Dir("web/assets")))
}
