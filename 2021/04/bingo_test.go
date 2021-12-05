package main

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

var (
	dummyWg      *sync.WaitGroup
	dummyNumCh   = make(chan string)
	dummyBingoCh = make(chan int)
)

func TestNewBoard(t *testing.T) {
	boardRaw := []string{
		"22 13 17 11 0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
	}

	t.Run("Rows", func(t *testing.T) {
		board := NewBoard(0, dummyWg, dummyNumCh, dummyBingoCh, boardRaw)
		want := []map[string]bool{
			map[string]bool{"22": false, "13": false, "17": false, "11": false, "0": false},
			map[string]bool{"8": false, "2": false, "23": false, "4": false, "24": false},
			map[string]bool{"21": false, "9": false, "14": false, "16": false, "7": false},
			map[string]bool{"6": false, "10": false, "3": false, "18": false, "5": false},
			map[string]bool{"1": false, "12": false, "20": false, "15": false, "19": false},
		}

		for i := 0; i < 5; i++ {
			got, err := board.GetRow(i)

			if err != nil {
				t.Fatalf("Got an error and didn't expect one: %v", err)
			}

			if fmt.Sprint(got) != fmt.Sprint(want[i]) {
				t.Errorf("got %v, want %v", got, want[i])
			}
		}
	})

	t.Run("Columns", func(t *testing.T) {
		board := NewBoard(0, dummyWg, dummyNumCh, dummyBingoCh, boardRaw)
		want := []map[string]bool{
			map[string]bool{"22": false, "8": false, "21": false, "6": false, "1": false},
			map[string]bool{"13": false, "2": false, "9": false, "10": false, "12": false},
			map[string]bool{"17": false, "23": false, "14": false, "3": false, "20": false},
			map[string]bool{"11": false, "4": false, "16": false, "18": false, "15": false},
			map[string]bool{"0": false, "24": false, "7": false, "5": false, "19": false},
		}

		for i := 0; i < 5; i++ {
			got, err := board.GetColumn(i)

			if err != nil {
				t.Fatalf("Got an error and didn't expect one: %v", err)
			}

			if fmt.Sprint(got) != fmt.Sprint(want[i]) {
				t.Errorf("got %v, want %v", got, want[i])
			}
		}
	})
}

func TestLoadInput(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19
`

	t.Run("load numbers", func(t *testing.T) {
		_, _, got, _ := LoadInput(input)
		want := []string{"7", "4", "9", "5", "11", "17", "23", "2", "0", "14", "21", "24", "10", "16", "13", "6", "15", "25", "12", "22", "18", "20", "8", "19", "3", "26", "1"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("load boards", func(t *testing.T) {
		_, _, _, got := LoadInput(input)
		want := NewBoard(0, dummyWg, dummyNumCh, dummyBingoCh, []string{
			"22 13 17 11  0",
			" 8  2 23  4 24",
			"21  9 14 16  7",
			" 6 10  3 18  5",
			" 1 12 20 15 19",
		})

		if !reflect.DeepEqual(got[0].rows, want.rows) {
			t.Errorf("got %v, want %v", got[0].rows, want.columns)
		}

		if !reflect.DeepEqual(got[0].columns, want.columns) {
			t.Errorf("got %v, want %v", got[0].rows, want.columns)
		}
	})
}

func TestBingo(t *testing.T) {
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

	got, err := PlayBingo(input)
	want := NewBoard(2, dummyWg, dummyNumCh, dummyBingoCh, []string{
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	})

	if err != nil {
		t.Fatal("Got an unexpected error: ", err)
	}

	if fmt.Sprint(got.raw) != fmt.Sprint(want.raw) {
		t.Errorf("got %v, want %v", got.raw, want.raw)
	}
}
