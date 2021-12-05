package main

import (
	"errors"
	"fmt"
	"strings"
)

type (
	row    map[string]bool
	column map[string]bool
)

func NewBoard(raw []string) *Board {
	board := &Board{}
	board.rows = make([]row, 5)
	board.columns = make([]column, 5)
	for i := 0; i < 5; i++ {
		board.rows[i] = make(row)
		board.columns[i] = make(column)
	}
	for i, r := range raw {
		board.rows[i] = make(row)
		// Remove double spaces and leading spaces
		r = strings.TrimLeft(r, " ")
		r = strings.ReplaceAll(r, "  ", " ")
		for j, val := range strings.Split(r, " ") {
			board.rows[i][val] = false
			board.columns[j][val] = false
		}
	}

	return board
}

type Board struct {
	rows    []row
	columns []column
}

func (b *Board) GetRow(n int) (row, error) {
	if len(b.rows) < n {
		return make(row), errors.New(fmt.Sprintf("Row does not exist: %d", n))
	}
	return b.rows[n], nil
}

func (b *Board) GetColumn(n int) (column, error) {
	if len(b.columns) < n {
		return make(column), errors.New(fmt.Sprintf("Column does not exist: %d", n))
	}
	return b.columns[n], nil
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
