package web

import (
	"mime"
	"net/http"
)

func init() {
	mime.AddExtensionType(".svg", "image/svg+xml")
	for _, path := range []string{"css", "images", "js"} {
		http.Handle("/"+path+"/", http.FileServer(http.Dir("web/assets")))
	}
}
