package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	board.lastNum = ""
	board.score = 0
	board.newNumCh = newNumCh
	board.bingoCh = bingoCh
	board.rows = make([]row, 5)
	board.columns = make([]column, 5)
	board.mu = new(sync.Mutex)

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
	lastNum  string
	score    int
	wg       *sync.WaitGroup
	newNumCh chan string
	bingoCh  chan int
	rows     []row
	columns  []column
	mu       *sync.Mutex
}

func (b *Board) markRow(idx int, num string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.rows[idx][num] = true
}

func (b *Board) markColumn(idx int, num string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.columns[idx][num] = true
}

func (b *Board) Start() {
	for {
		b.lastNum = <-b.newNumCh
		for idx, r := range b.rows {
			if _, ok := r[b.lastNum]; ok {
				b.markRow(idx, b.lastNum)
			}
		}
		for idx, c := range b.columns {
			if _, ok := c[b.lastNum]; ok {
				b.markColumn(idx, b.lastNum)
			}
		}

	rowLoop:
		for _, r := range b.rows {
			for _, v := range r {
				if v != true {
					continue rowLoop
				}
			}
			b.calculateScore()
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
			b.calculateScore()
			b.bingoCh <- b.id
			break
		}
		b.wg.Done()
	}
}

func (b *Board) calculateScore() {
	b.mu.Lock()
	defer b.mu.Unlock()
	unmarkedSum := 0
	for _, r := range b.rows {
		for k, v := range r {
			if !v {
				kInt, err := strconv.Atoi(k)
				if err != nil {
					panic(err)
				}
				// fmt.Printf("Sum: %d New: %d\n", unmarkedSum, kInt)
				unmarkedSum += kInt
			}
		}
	}
	lastNumInt, err := strconv.Atoi(b.lastNum)
	if err != nil {
		panic(err)
	}

	b.score = unmarkedSum * lastNumInt
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

	fmt.Println(input)
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
			fmt.Printf("Board %d reported a bingo!\n", bingo+1)
			return boards[bingo], nil
		case <-noBingo:
			continue
		}
	}

	// TODO: Return error if we get here, ran out of numbers with no winner
	return &Board{}, errors.New("No board got a bingo")
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprint("Could not open file: ", err))
	}
	defer file.Close()

	bingoData := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bingoData += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprint("Error scanning file: ", err))
	}

	winner, _ := PlayBingo(bingoData)
	fmt.Println("Found a winner: ", winner.id, winner.score)
}
