package main

import (
	"fmt"
	"reflect"
	"testing"
)

const input = "16,1,2,0,4,2,7,1,2,14"

func TestLoadInput(t *testing.T) {
	got, err := LoadInput(input)
	want := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	if err != nil {
		t.Fatal("Got an error but didn't expect one: ", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinMax(t *testing.T) {
	numbers, _ := LoadInput(input)
	gotMin, gotMax := FindMinMax(numbers)
	wantMin, wantMax := 0, 16

	if gotMin != wantMin {
		t.Errorf("got %d, want %d", gotMin, wantMin)
	}

	if gotMax != wantMax {
		t.Errorf("got %d, want %d", gotMax, wantMax)
	}
}

func TestBuildMaps(t *testing.T) {
	min, max := 0, 4

	examples := []struct {
		num  int
		want NumMap
	}{
		{num: 0, want: NumMap{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}},
		{num: 1, want: NumMap{0: 1, 1: 0, 2: 1, 3: 2, 4: 3}},
		{num: 2, want: NumMap{0: 2, 1: 1, 2: 0, 3: 1, 4: 2}},
		{num: 3, want: NumMap{0: 3, 1: 2, 2: 1, 3: 0, 4: 1}},
		{num: 4, want: NumMap{0: 4, 1: 3, 2: 2, 3: 1, 4: 0}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprint(tt.num), func(t *testing.T) {
			got, err := BuildMap(tt.num, min, max)

			if err != nil {
				t.Fatal("Got an error and didn't expect one: ", err)
			}

			if fmt.Sprint(got) != fmt.Sprint(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLeastFuelPosition(t *testing.T) {
	numbers, _ := LoadInput(input)

	got, err := GetLeastFuelPosition(numbers)
	want := 2

	if err != nil {
		t.Fatal("Got error and didn't expect one: ", err)
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
