package main

import (
	"reflect"
	"testing"
)

func TestCommandParsing(t *testing.T) {
	examples := []struct {
		command string
		want    Command
	}{
		{command: "forward 5", want: Command{HORIZONTAL, 5}},
		{command: "backward 15", want: Command{HORIZONTAL, -15}},
		{command: "down 3", want: Command{VERTICAL, 3}},
		{command: "up 4", want: Command{VERTICAL, -4}},
	}

	for _, tt := range examples {
		t.Run(tt.command, func(t *testing.T) {
			got := ParseCommand(tt.command)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestSubTracking(t *testing.T) {
	t.Run("advent examples", func(t *testing.T) {
		commands := []Command{
			ParseCommand("forward 5"),
			ParseCommand("down 5"),
			ParseCommand("forward 8"),
			ParseCommand("up 3"),
			ParseCommand("down 8"),
			ParseCommand("forward 2"),
		}

		sub := NewSub()
		sub.IssueCommands(commands)

		got1 := sub.CurrentLocation
		want1 := Location{15, 60}

		if got1 != want1 {
			t.Fatalf("got %+v, want %+v", got1, want1)
		}

		got2 := sub.Final()
		want2 := 900

		if got2 != want2 {
			t.Fatalf("got %d, want %d", got2, want2)
		}
	})
}
