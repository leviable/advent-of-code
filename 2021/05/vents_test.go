package main

import (
	"reflect"
	"testing"
)

const input = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestLoadInput(t *testing.T) {

	got, err := LoadInput(input)
	want := []Line{
		Line{Point{0, 9}, Point{5, 9}},
		Line{Point{8, 0}, Point{0, 8}},
		Line{Point{9, 4}, Point{3, 4}},
		Line{Point{2, 2}, Point{2, 1}},
		Line{Point{7, 0}, Point{7, 4}},
		Line{Point{6, 4}, Point{2, 0}},
		Line{Point{0, 9}, Point{2, 9}},
		Line{Point{3, 4}, Point{1, 4}},
		Line{Point{0, 0}, Point{8, 8}},
		Line{Point{5, 5}, Point{8, 2}},
	}

	if err != nil {
		t.Fatal("Got an error and didn't expect one: ", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetGrid(t *testing.T) {
	lines, _ := LoadInput(input)

	t.Run("grid creation", func(t *testing.T) {
		got := GetGrid(lines)
		want := &Grid{Point{0, 0}, Point{9, 9}, lines, make(map[Point]int)}

		if !reflect.DeepEqual(got.origin, want.origin) {
			t.Errorf("got %v, want %v", got.origin, want.origin)
		}

		if !reflect.DeepEqual(got.end, want.end) {
			t.Errorf("got %v, want %v", got.end, want.end)
		}
	})

	t.Run("danger score", func(t *testing.T) {
		grid := GetGrid(lines)
		got := grid.DangerScore()
		want := 5

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})
}
