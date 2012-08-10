package app

import (
	"encoding/gob"
	"errors"
	"log"
	"math/rand"
	"strings"
)

type Guess struct {
	Complete  bool
	Correct   string
	Incorrect string
}

type GameState struct {
	CurrentWord string
	Guesses     map[string]Guess
}

func init() {
	gob.Register(GameState{})
}

func (s GameState) CurrentLetters() []string {
	return strings.Split(s.CurrentWord, "")
}

func (s GameState) CurrentWordScore(correct, incorrect int) int {
	return s.wordScore(s.CurrentWord, correct, incorrect)
}

func (s *GameState) Guess(letter string) (correct bool, err error) {
	correct = strings.Contains(s.CurrentWord, letter)
	g := s.Guesses[s.CurrentWord]
	switch {
	case strings.Contains(g.Correct, letter) || strings.Contains(g.Incorrect, letter):
		err = errors.New("You have already guessed " + letter)
	case correct:
		g.Correct += letter
		g.Complete = true
		for i := 0; i < len(s.CurrentWord); i++ {
			if !strings.Contains(g.Correct, s.CurrentWord[i:i+1]) {
				g.Complete = false
				break
			}
		}
	case !correct:
		g.Incorrect += letter
	}
	s.Guesses[s.CurrentWord] = g
	return
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
	for _, g := range s.Guesses {
		won = won && g.Complete
	}
	return
}

func (s GameState) IsSpace(idx int) bool {
	return s.CurrentLetters()[idx] == " "
}

func (s GameState) TotalScore(correct, incorrect int) (total int) {
	for w := range s.Guesses {
		total += s.wordScore(w, correct, incorrect)
	}
	return
}

func (s *GameState) UpdateCurrent() bool {
	var (
		newWord string
		unseen  = []string{}
	)
	// Initialize guesses if need-be
	log.Printf("GUESSES: nil: %+v len: %d", s.Guesses == nil, len(s.Guesses))
	if s.Guesses == nil {
		s.Guesses = make(map[string]Guess, len(words))
		for _, w := range words {
			s.Guesses[w] = Guess{}
		}
	}
	// Check if we've already won - no need to pick random out of nothing
	if s.HasWon() {
		return false
	}
	// Gather unseen
	for w, guess := range s.Guesses {
		if !guess.Complete {
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

func (s GameState) wordScore(w string, correct, incorrect int) int {
	return len(s.Guesses[w].Correct)*correct + len(s.Guesses[s.CurrentWord].Incorrect)*incorrect
}
