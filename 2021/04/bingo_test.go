package main

import (
	"fmt"
	"reflect"
	"testing"
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
		board := NewBoard(boardRaw)
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
		got, _ := LoadInput(input)
		want := []string{"7", "4", "9", "5", "11", "17", "23", "2", "0", "14", "21", "24", "10", "16", "13", "6", "15", "25", "12", "22", "18", "20", "8", "19", "3", "26", "1"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("load boards", func(t *testing.T) {
		_, got := LoadInput(input)
		want := NewBoard([]string{
			"22 13 17 11  0",
			" 8  2 23  4 24",
			"21  9 14 16  7",
			" 6 10  3 18  5",
			" 1 12 20 15 19",
		})

		if !reflect.DeepEqual(got[0], want) {
			t.Errorf("got %v, want %v", got[0], want)
		}
	})
}
