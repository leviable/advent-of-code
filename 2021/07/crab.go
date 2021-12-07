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
	fmt.Printf("Min/Max: %d - %d\n", min, max)
	return
}

func BuildMap(num, min, max int) (NumMap, error) {
	m := make(NumMap)
	if num < min || num > max {
		return m, errors.New(fmt.Sprintf("Number %d not between %d and %d\n", num, min, max))
	}
	for x := min; x <= max; x++ {
		m[x] = abs(x - num)
	}
	return m, nil
}

func GetLeastFuelPosition(numbers []int) (int, error) {
	min, max := FindMinMax(numbers)

	m := make(map[int]NumMap)
	for x := min; x <= max; x++ {
		numMap, err := BuildMap(x, min, max)
		if err != nil {
			return -1, err
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
			fmt.Printf("Least fuel %d at position %d\n", leastFuel, leastFuelPosition)
			continue
		}
	}
	return leastFuelPosition, nil
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

	leastFuelPosition, err := GetLeastFuelPosition(numbers)
	if err != nil {
		panic(fmt.Sprint("got an error and didn't expect one: ", err))
	}

	fmt.Println("Least fuel position is: ", leastFuelPosition)
}
