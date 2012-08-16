package web

import (
	"mime"
	"net/http"
)

func init() {
	mime.AddExtensionType(".svg", "image/svg+xml")
	http.Handle("/images/", http.FileServer(http.Dir("web/assets")))
}
