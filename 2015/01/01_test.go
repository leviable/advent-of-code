package main

import "testing"

func TestDecodeDirections(t *testing.T) {
	examples := []struct {
		directions string
		want       int
	}{
		{directions: "(())", want: 0},
		{directions: "()()", want: 0},
		{directions: "(((", want: 3},
		{directions: "(()(()(", want: 3},
		{directions: "))(((((", want: 3},
		{directions: "())", want: -1},
		{directions: "))(", want: -1},
		{directions: ")))", want: -3},
		{directions: ")())())", want: -3},
	}
	for _, tt := range examples {
		t.Run(tt.directions, func(t *testing.T) {
			got := DecodeDirections(tt.directions, PARTONE)
			want := tt.want

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestEntersBasement(t *testing.T) {
	got := DecodeDirections(")", PARTTWO)
	want := 1

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
