package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pattern struct {
	wires  []string
	output []string
}

func LoadInput(input string) []Pattern {
	lines := strings.Split(input, "\n")
	patterns := make([]Pattern, len(lines))
	for idx, line := range lines {
		s := strings.Split(line, " | ")
		if len(s) < 2 {
			continue
		}
		wiresRaw, outputRaw := s[0], s[1]
		wires := strings.Split(wiresRaw, " ")
		output := strings.Split(outputRaw, " ")
		patterns[idx] = Pattern{wires: wires, output: output}
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
}
