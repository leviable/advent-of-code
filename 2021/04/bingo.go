package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type (
	row    map[string]bool
	column map[string]bool
)

func NewBoard(id int, wg *sync.WaitGroup, newNumCh chan string, bingoCh chan int, raw []string) *Board {
	board := new(Board)
	board.id = id
	board.wg = wg
	board.raw = raw
	board.newNumCh = newNumCh
	board.bingoCh = bingoCh
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

	go board.Start()

	return board
}

type Board struct {
	id       int
	raw      []string
	wg       *sync.WaitGroup
	newNumCh chan string
	bingoCh  chan int
	rows     []row
	columns  []column
}

func (b *Board) Start() {
	for {
		newNum := <-b.newNumCh
		for _, r := range b.rows {
			if _, ok := r[newNum]; ok {
				r[newNum] = true
			}
		}
		for _, c := range b.columns {
			if _, ok := c[newNum]; ok {
				c[newNum] = true
			}
		}

	rowLoop:
		for _, r := range b.rows {
			for _, v := range r {
				if v != true {
					continue rowLoop
				}
			}
			b.bingoCh <- b.id
			break
		}

	columnLoop:
		for _, c := range b.columns {
			for _, v := range c {
				if !v {
					continue columnLoop
				}
			}
			b.bingoCh <- b.id
			break
		}
		b.wg.Done()
	}
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

func LoadInput(input string) (*sync.WaitGroup, chan int, []string, map[int]*Board) {

	wg := new(sync.WaitGroup)

	newNumCh := make(chan string)
	bingoCh := make(chan int)

	splitInput := strings.Split(input, "\n")

	bingoNumbers := strings.Split(splitInput[0], ",")

	boardsRaw := splitInput[2:]

	boards := make(map[int]*Board)

	boardId := 0
	for len(boardsRaw) >= 5 {
		b := boardsRaw[:5]
		if len(boardsRaw) > 5 {
			boardsRaw = boardsRaw[6:]
		} else {
			boardsRaw = boardsRaw[5:]
		}

		boards[boardId] = NewBoard(boardId, wg, newNumCh, bingoCh, b)

		boardId++
	}
	return wg, bingoCh, bingoNumbers, boards
}

func PlayBingo(input string) (*Board, error) {
	wg, bingoCh, numbers, boards := LoadInput(input)

	boardCount := len(boards)
	fmt.Printf("Loaded %d boards\n", boardCount)

	for _, num := range numbers {
		// Do broadcast here

		wg.Add(boardCount)
		noBingo := make(chan struct{})
		go func() {
			wg.Wait()
			close(noBingo)
		}()

		for _, b := range boards {
			b.newNumCh <- num
		}

		select {
		case bingo := <-bingoCh:
			fmt.Printf("Board %d reported a bingo!\n", bingo)
			return boards[bingo], nil
		case <-noBingo:
			continue
		}
	}

	// TODO: Return error if we get here, ran out of numbers with no winner
	return &Board{}, errors.New("No board got a bingo")
}

func main() {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

	PlayBingo(input)
}
