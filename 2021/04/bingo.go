package main

import (
	"errors"
	"fmt"
	"strings"
)

type row map[string]bool

func newRow(raw string) (r row) {
	r = make(row)
	for _, c := range strings.Split(raw, " ") {
		if c == "" {
			continue
		}
		r[c] = false
	}

	return
}

func NewBoard(raw string) *Board {
	board := &Board{}
	for _, r := range strings.Split(raw, "\n") {
		board.rows = append(board.rows, newRow(r))
	}

	return board
}

type Board struct {
	rows []row
}

func (b *Board) GetRow(n int) (row, error) {
	if len(b.rows) < n {
		return make(row), errors.New(fmt.Sprintf("Row does not exist: %d", n))
	}
	return b.rows[n], nil
}
