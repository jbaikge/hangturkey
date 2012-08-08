package web

import (
	"code.google.com/p/gorilla/sessions"
	"encoding/gob"
	"log"
	"math/rand"
	"strings"
)

type GameState struct {
	Scores      map[string]int
	CurrentWord string
}

func init() {
	gob.Register(GameState{})
}



func (s GameState) HasWon() (won bool) {
	won = true
	for _, score := range s.Scores {
		if score == 0 {
			won = false
		}
	}
	return
}

func (s GameState) Letters() []string {
	return strings.Split(s.CurrentWord, "")
}

func (s *GameState) UpdateCurrent() bool {
	var (
		newWord string
		unseen  = []string{}
	)
	// Initialize scores if need-be
	if s.Scores == nil {
		s.Scores = make(map[string]int, len(words))
		for _, w := range words {
			s.Scores[w] = 0
		}
	}
	// Check if we've already won - no need to pick random out of nothing
	if s.HasWon() {
		return false
	}
	// Gather unseen
	for w, score := range s.Scores {
		if score == 0 {
			unseen = append(unseen, w)
		}
	}
	// Pick a random unseen
	log.Printf("UNSEEN LEN %d", len(unseen))
	newWord = unseen[rand.Intn(len(unseen))]
	log.Printf("STATE      %+v", s)
	log.Printf("NEW WORD   %s", newWord)
	s.CurrentWord = newWord
	return true
}
