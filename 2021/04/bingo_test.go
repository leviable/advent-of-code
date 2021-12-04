package main

import (
	"fmt"
	"testing"
)

func TestNewBoard(t *testing.T) {
	boardRaw := `22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19`

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
