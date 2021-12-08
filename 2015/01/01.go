package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var UpOrDown = map[string]int{"(": 1, ")": -1}

func DecodeDirections(directions string) (floor int) {
	for _, d := range strings.Split(directions, "") {
		floor += UpOrDown[d]
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

	fmt.Println("Last floor is: ", DecodeDirections(input[0]))
}
