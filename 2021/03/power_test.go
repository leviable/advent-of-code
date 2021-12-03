package main

import "testing"

func TestConversion(t *testing.T) {
	t.Run("convert gamma", func(t *testing.T) {
		got := ConvertRate("10110")
		want := 22

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("convert epsilon", func(t *testing.T) {
		got := ConvertRate("01001")
		want := 9

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestCrunchDiag(t *testing.T) {
	diagReport := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	t.Run("test crunch", func(t *testing.T) {
		gotGamma, gotEpsilon := CrunchDiag(diagReport)
		wantGamma, wantEpsilon := "10110", "01001"

		if gotGamma != wantGamma {
			t.Errorf("got %q, want %q", gotGamma, wantGamma)
		}

		if gotEpsilon != wantEpsilon {
			t.Errorf("got %q, want %q", gotEpsilon, wantEpsilon)
		}
	})

	t.Run("test final", func(t *testing.T) {
		gamma, epsilon := CrunchDiag(diagReport)
		got, err := GetPower(gamma, epsilon)
		want := 198

		if err != nil {
			t.Fatal("Got an error but didn't expect one: ", err)
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
