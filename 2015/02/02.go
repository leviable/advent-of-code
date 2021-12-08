package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalculatePaper(dimensions string) int {
	l, w, h := GetDimensions(dimensions)

	side1 := l * w
	side2 := l * h
	side3 := w * h

	extra, _ := GetSmallest(side1, side2, side3)

	return 2*side1 + 2*side2 + 2*side3 + extra
}

func CalculateRibbon(dimensions string) int {
	l, w, h := GetDimensions(dimensions)
	smallest1, smallest2 := GetSmallest(l, w, h)
	ribbon := 2*smallest1 + 2*smallest2
	bow := l * w * h
	return ribbon + bow
}

func GetDimensions(dimensions string) (int, int, int) {

	d := strings.Split(dimensions, "x")
	toInt := func(num string) int {
		n, _ := strconv.Atoi(num)
		return n
	}

	return toInt(d[0]), toInt(d[1]), toInt(d[2])
}

func GetSmallest(s1, s2, s3 int) (int, int) {
	sides := []int{s1, s2, s3}
	sort.Ints(sides)

	return sides[0], sides[1]
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

	totalPaper, totalRibbon := 0, 0
	for _, d := range input {
		totalPaper += CalculatePaper(d)
		totalRibbon += CalculateRibbon(d)
	}

	fmt.Println("Total paper required: ", totalPaper)
	fmt.Println("Total ribbon required: ", totalRibbon)
}
