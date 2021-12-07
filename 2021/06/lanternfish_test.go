package main

import (
	"fmt"
	"reflect"
	"testing"
)

const input = "3,4,3,1,2"

func TestLoadInput(t *testing.T) {
	got, err := LoadInput(input)
	want := []*LanternFish{
		&LanternFish{3},
		&LanternFish{4},
		&LanternFish{3},
		&LanternFish{1},
		&LanternFish{2},
	}

	if err != nil {
		t.Error("Got an error and didn't expect one: ", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFishAging(t *testing.T) {
	fish := []struct {
		name      string
		age, want int
	}{
		{age: 5, want: 4},
		{age: 4, want: 3},
		{age: 3, want: 2},
		{age: 2, want: 1},
		{age: 1, want: 0},
		{age: 0, want: FISHTIMERRESET},
	}

	for _, tt := range fish {
		t.Run(fmt.Sprint(tt.age), func(t *testing.T) {
			f := &LanternFish{tt.age}
			f.GrowOlder()

			got := f.timer
			want := tt.want

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestSpawning(t *testing.T) {
	fish := []struct {
		name string
		age  int
		want spawnSignal
	}{
		{name: "about to spawn", age: 0, want: NEWSPAWN},
		{name: "won't spawn", age: 1, want: NOSPAWN},
	}
	for _, tt := range fish {
		t.Run(tt.name, func(t *testing.T) {
			f := &LanternFish{tt.age}

			got := f.GrowOlder()

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNewSchool(t *testing.T) {
	daysToAge := 10
	fish, _ := LoadInput(input)

	got := NewSchool(fish, daysToAge).Size()
	want := 5

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSchoolAging(t *testing.T) {

	days := []struct {
		name string
		days int
		size int
	}{
		{name: "after 1 day", days: 1, size: 5},
		{name: "after 2 days", days: 2, size: 6},
		{name: "after 3 days", days: 3, size: 7},
		{name: "after 4 days", days: 4, size: 9},
		{name: "after 11 days", days: 11, size: 15},
		{name: "after 18 days", days: 18, size: 26},
		{name: "after 80 days", days: 80, size: 5934},
		{name: "after 170 days", days: 170, size: 5934},
		{name: "after 256 days", days: 256, size: 26984457539},
	}

	for _, tt := range days {
		t.Run(tt.name, func(t *testing.T) {
			fish, _ := LoadInput(input)
			school := NewSchool(fish, tt.days)
			school.Grow()
			got := school.Size()
			want := tt.size

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
