package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Part int

const (
	PARTONE Part = iota
	PARTTWO
)

var UpOrDown = map[string]int{"(": 1, ")": -1}

func DecodeDirections(directions string, part Part) (floor int) {
	for idx, d := range strings.Split(directions, "") {
		floor += UpOrDown[d]
		if part == PARTTWO && floor < 0 {
			return idx + 1
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

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprint("Error scanning file: ", err))
	}

	fmt.Println("Last floor is: ", DecodeDirections(input[0], PARTONE))
	fmt.Println("First time to basement is: ", DecodeDirections(input[0], PARTTWO))
}
