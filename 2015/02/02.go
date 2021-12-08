package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculatePaper(dimensions string) (area int) {
	d := strings.Split(dimensions, "x")

	l, w, h := toInt(d[0]), toInt(d[1]), toInt(d[2])

	side1 := l * w
	side2 := l * h
	side3 := w * h

	var extra int
	if side1 <= side2 && side1 <= side3 {
		extra = side1
	} else if side2 <= side1 && side2 <= side3 {
		extra = side2
	} else {
		extra = side3
	}

	return 2*side1 + 2*side2 + 2*side3 + extra
}

func toInt(num string) int {
	n, _ := strconv.Atoi(num)

	return n
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

	total := 0
	for _, d := range input {
		total += CalculatePaper(d)
	}

	fmt.Println("Total paper required: ", total)
}
