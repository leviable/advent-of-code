package main

import "testing"

func TestCalculatePaper(t *testing.T) {
	examples := []struct {
		dimensions string
		want       int
	}{
		{dimensions: "2x3x4", want: 58},
		{dimensions: "1x1x10", want: 43},
	}

	for _, tt := range examples {
		t.Run(tt.dimensions, func(t *testing.T) {
			got := CalculatePaper(tt.dimensions)
			want := tt.want

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestCalculateRibbon(t *testing.T) {
	examples := []struct {
		dimensions string
		want       int
	}{
		{dimensions: "2x3x4", want: 34},
		{dimensions: "1x1x10", want: 14},
	}

	for _, tt := range examples {
		t.Run(tt.dimensions, func(t *testing.T) {
			got := CalculateRibbon(tt.dimensions)
			want := tt.want

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
