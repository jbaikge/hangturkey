package main

import (
	"encoding/gob"
	"net/http"

	_ "github.com/jbaikge/hangturkey/app"
)

func main() {
	log.Fatal(http.ListenAndServe(":9000", nil))
}
