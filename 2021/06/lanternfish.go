package main

import (
	"strconv"
	"strings"
)

type spawnSignal int

const (
	NEWFISHTIMER   = 8
	FISHTIMERRESET = 6

	NEWSPAWN spawnSignal = iota
	NOSPAWN
)

type LanternFish struct {
	timer int
}

func (l *LanternFish) GrowOlder() (*LanternFish, spawnSignal) {
	newFish := &LanternFish{NEWFISHTIMER}
	signal := NOSPAWN
	l.timer--
	if l.timer < 0 {
		l.timer = FISHTIMERRESET
		signal = NEWSPAWN
	}

	return newFish, signal
}

func LoadInput(input string) ([]*LanternFish, error) {
	fish := strings.Split(input, ",")
	lf := make([]*LanternFish, len(fish))

	for idx, f := range fish {
		timer, err := strconv.Atoi(f)
		if err != nil {
			return []*LanternFish{}, err
		}
		lf[idx] = &LanternFish{timer}
	}

	return lf, nil
}

func NewSchool(fish []*LanternFish) *School {
	return &School{fish}
}

type School struct {
	fish []*LanternFish
}

func (s *School) Size() int {
	return len(s.fish)
}

func (s *School) NextDay() {
	for _, fish := range s.fish[:] {
		newFish, signal := fish.GrowOlder()
		if signal == NEWSPAWN {
			s.fish = append(s.fish, newFish)
		}
	}
}
