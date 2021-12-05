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

func NewBoard(raw []string) *Board {
	board := &Board{}
	for _, r := range raw {
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

func LoadInput(input string) ([]string, []*Board) {
	splitInput := strings.Split(input, "\n")

	bingoNumbers := strings.Split(splitInput[0], ",")

	boards := make([]*Board, 0)
	boardsRaw := splitInput[2:]

	for len(boardsRaw) > 0 {
		b := boardsRaw[:5]
		if len(boardsRaw) > 5 {
			boardsRaw = boardsRaw[6:]
		}

		boards = append(boards, NewBoard(b))
	}
	return bingoNumbers, boards
}
