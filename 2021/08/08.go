package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type decoder map[string]int

type Pattern struct {
	wires   []string
	decoder decoder
	output  []string
	value   int
}

func LoadInput(input string) []Pattern {
	lines := strings.Split(input, "\n")
	patterns := make([]Pattern, 0)
loop:
	for _, line := range lines {
		s := strings.Split(line, " | ")
		if line == "" || len(s) < 2 {
			continue loop
		}
		wiresRaw, outputRaw := s[0], s[1]
		wires := strings.Split(wiresRaw, " ")
		output := strings.Split(outputRaw, " ")
		decoder := make(decoder)
		patterns = append(patterns, Pattern{wires: wires, decoder: decoder, output: output})
	}
	return patterns
}

func CountUniqueDigits(patterns []Pattern) (sum int) {

	for _, pattern := range patterns {
		for _, val := range pattern.output {
			switch len(val) {
			case 2, 3, 4, 7:
				sum++
			}
		}
	}

	return
}

func Decode(pattern *Pattern) {
	// Sort each wire into a consistent order
	wires := make([]string, 10)
	for idx, w := range pattern.wires {
		wSlice := strings.Split(w, "")
		sort.Strings(wSlice)
		wires[idx] = strings.Join(wSlice, "")
	}

	knownWires := make([]string, 10)
	for _, wire := range wires {
		// Check for numbers with unique segments
		switch len(wire) {
		case 2:
			pattern.decoder[wire] = 1
			knownWires[1] = wire
			continue
		case 3:
			pattern.decoder[wire] = 7
			knownWires[7] = wire
			continue
		case 4:
			pattern.decoder[wire] = 4
			knownWires[4] = wire
			continue
		case 7:
			pattern.decoder[wire] = 8
			knownWires[8] = wire
			continue
		}

	}
	// Check for wires with len(segments) == 6
	// 0, 6, 9
	// 0 cannot contain all of 4, and must contain 1
	// 6 cannot contain all of 1
	// 9 must contain all of 4
	for _, wire := range wires {
		// Skip known wires
		if _, ok := pattern.decoder[wire]; ok {
			continue
		}
		if len(wire) == 6 {
			if SegmentContains(wire, knownWires[4]) {
				pattern.decoder[wire] = 9
				knownWires[9] = wire
			} else if SegmentContains(wire, knownWires[1]) {
				pattern.decoder[wire] = 0
				knownWires[0] = wire
			} else {
				pattern.decoder[wire] = 6
				knownWires[6] = wire
			}
		}

	}
	// Check for wires with len(segments) == 5
	// 2, 3, 5
	// 2 cannot contain 1
	// 3 must contain 1
	// 5 must be within 6
	for _, wire := range wires {
		if _, ok := pattern.decoder[wire]; ok {
			continue
		}
		if SegmentContains(wire, knownWires[1]) {
			pattern.decoder[wire] = 3
			knownWires[3] = wire
		} else if SegmentContains(knownWires[6], wire) {
			pattern.decoder[wire] = 5
			knownWires[5] = wire
		} else {
			pattern.decoder[wire] = 2
			knownWires[2] = wire
		}
	}

	decodedOutput := make([]int, len(pattern.output))
	for idx, output := range pattern.output {
		oSlice := strings.Split(output, "")
		sort.Strings(oSlice)
		outputSorted := strings.Join(oSlice, "")

		decodedOutput[idx] = pattern.decoder[outputSorted]

	}

	outLen := len(decodedOutput) - 1
	for idx, n := range decodedOutput {
		pattern.value += n * int(math.Pow10(outLen-idx))
	}
}

func SegmentContains(segment, sub string) bool {
	for _, c := range strings.Split(sub, "") {
		if !strings.Contains(segment, c) {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprint("Could not open file: ", err))
	}
	defer file.Close()

	wireData := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wireData += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprint("Error scanning file: ", err))
	}

	input := LoadInput(wireData)

	fmt.Println("Total unique output digits: ", CountUniqueDigits(input))

	sum := 0
	for _, pattern := range input {
		// fmt.Println("Pattern is: ", pattern)
		Decode(&pattern)
		sum += pattern.value
		// fmt.Println("Sum is: ", sum)
	}
	fmt.Println("Output value sum is: ", sum)

}
