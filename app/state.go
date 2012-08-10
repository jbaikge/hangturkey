package app

import (
	"encoding/gob"
	"log"
	"math/rand"
	"strings"
)

type Guess struct {
	Correct   string
	Incorrect string
}

type GameState struct {
	CurrentWord string
	Guesses     map[string]Guess
	Scores      map[string]int
}

func init() {
	gob.Register(GameState{})
}

func (s GameState) CurrentLetters() []string {
	return strings.Split(s.CurrentWord, "")
}

func (s GameState) GuessedLetters() []string {
	guessed := make([]string, len(s.CurrentWord))
	for i, l := range s.CurrentLetters() {
		if strings.Contains(s.Guesses[s.CurrentWord].Correct, l) {
			guessed[i] = l
		}
	}
	return guessed
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

func (s GameState) IsSpace(idx int) bool {
	return s.CurrentLetters()[idx] == " "
}

func (s *GameState) UpdateCurrent() bool {
	var (
		newWord string
		unseen  = []string{}
	)
	// Initialize scores if need-be
	if s.Scores == nil {
		s.Scores = make(map[string]int, len(words))
		s.Guesses = make(map[string]Guess, len(words))
		for _, w := range words {
			s.Scores[w] = 0
			s.Guesses[w] = Guess{}
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
