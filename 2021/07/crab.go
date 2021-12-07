package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumMap map[int]int

type Part int

const (
	PARTONE Part = iota
	PARTTWO
)

func LoadInput(raw string) ([]int, error) {
	input := strings.Split(raw, ",")
	i := make([]int, len(input))

	for idx, numStr := range input {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return i, err
		}
		i[idx] = num
	}

	return i, nil
}

func FindMinMax(numbers []int) (min, max int) {
	min, max = numbers[0], numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

func BuildMap(num, min, max int, part Part) (NumMap, error) {
	partTwoCost := make([]int, max+1)
	c := 0
	for moves := 0; moves <= max; moves++ {
		c += moves
		partTwoCost[moves] += c
	}
	m := make(NumMap)
	if num < min || num > max {
		return m, errors.New(fmt.Sprintf("Number %d not between %d and %d\n", num, min, max))
	}
	for x := min; x <= max; x++ {
		switch part {
		case PARTONE:
			m[x] = abs(x - num)
		case PARTTWO:
			m[x] = partTwoCost[abs(x-num)]
		}
	}
	return m, nil
}

func GetLeastFuelPosition(numbers []int, part Part) (int, int, error) {
	min, max := FindMinMax(numbers)

	m := make(map[int]NumMap)
	for x := min; x <= max; x++ {
		numMap, err := BuildMap(x, min, max, part)
		if err != nil {
			return -1, -1, err
		}
		m[x] = numMap
	}

	leastFuel := -1
	leastFuelPosition := -1
	for x := min; x <= max; x++ {
		fuel := 0
		for _, n := range numbers {
			fuel += m[n][x]
		}
		if leastFuel == -1 || fuel < leastFuel {
			leastFuel = fuel
			leastFuelPosition = x
			continue
		}
	}
	return leastFuelPosition, leastFuel, nil
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprint("Could not open file: ", err))
	}
	defer file.Close()

	crabData := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		crabData += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprint("Error scanning file: ", err))
	}

	numbers, err := LoadInput(crabData)
	if err != nil {
		panic(fmt.Sprint("Got an error while loading input data: ", err))
	}

	fmt.Printf("Loaded %d crabs\n", len(numbers))

	leastFuelPosition, leastFuel, err := GetLeastFuelPosition(numbers, PARTONE)
	if err != nil {
		panic(fmt.Sprint("got an error and didn't expect one: ", err))
	}

	fmt.Printf("Part One Pos/Fuel: %d -> %d\n", leastFuelPosition, leastFuel)

	leastFuelPosition, leastFuel, err = GetLeastFuelPosition(numbers, PARTTWO)
	if err != nil {
		panic(fmt.Sprint("got an error and didn't expect one: ", err))
	}

	fmt.Printf("Part Two Pos/Fuel: %d -> %d\n", leastFuelPosition, leastFuel)
}
