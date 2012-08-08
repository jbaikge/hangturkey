package main

import (
	"log"
	"net/http"

	_ "github.com/jbaikge/hangturkey/web"
)

func main() {
	log.Fatal(http.ListenAndServe(":9000", nil))
}
