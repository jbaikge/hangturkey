package main

import (
	"encoding/json"
	"log"
	"os"
)

var words []string

func init() {
	var err error
	if words, err = GetWordList("words.json"); err != nil {
		log.Fatal(err)
	}
}

func GetWordList(filename string) (words []string, err error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return
	}
	err = json.NewDecoder(f).Decode(&words)
	if err != nil {
		return
	}
	return
}
