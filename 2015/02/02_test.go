package main

import "testing"

func TestCalculatePaper(t *testing.T) {
	examples := []struct {
		dimensions string
		want       int
	}{
		{dimensions: "2x3x4", want: 58},
		{dimensions: "1x1x10", want: 42},
	}

	for _, tt := range examples {
		t.Run(tt.dimensions, func(t *testing.T) {
			got := CalculatePaper("2x3x4")
			want := 58

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
