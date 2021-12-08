package main

import (
	"fmt"
	"reflect"
	"testing"
)

const input = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

func TestLoadInput(t *testing.T) {
	got := LoadInput(input)
	want := []Pattern{
		Pattern{
			wires:   []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"},
			decoder: decoder{},
			output:  []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
		},
	}

	if !reflect.DeepEqual(got[:1], want) {
		t.Errorf("got %v, want %v", got[0], want)
	}

}

func TestCountUniqueDigits(t *testing.T) {
	patterns := LoadInput(input)

	got := CountUniqueDigits(patterns)
	want := 26

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSegmentContains(t *testing.T) {
	got := SegmentContains("abcdef", "abef")
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestDecoder(t *testing.T) {
	t.Run("decoder", func(t *testing.T) {
		pattern := LoadInput("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")[0]
		Decode(&pattern)

		got := pattern.decoder
		want := decoder{
			"abcdefg": 8,
			"bcdef":   5,
			"acdfg":   2,
			"abcdf":   3,
			"abd":     7,
			"abcdef":  9,
			"bcdefg":  6,
			"abef":    4,
			"abcdeg":  0,
			"ab":      1,
		}

		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("got %v, want %v", got, want)
		}

		gotValue := pattern.value
		wantValue := 5353
		if gotValue != wantValue {
			t.Errorf("got %v, want %v", gotValue, wantValue)
		}
	})

	t.Run("single line decode", func(t *testing.T) {
		pattern := LoadInput("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe")[0]
		Decode(&pattern)

		got := pattern.value
		want := 8394

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})

}

func TestSum(t *testing.T) {
	got := 0
	for _, pattern := range LoadInput(input) {
		Decode(&pattern)
		got += pattern.value
	}

	want := 61229

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
