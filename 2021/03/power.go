package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ConvertRate(raw string) int {
	i, err := strconv.ParseInt(raw, 2, 64)
	if err != nil {
		panic(fmt.Sprint("Error converting to integer: ", raw))
	}
	return int(i)
}

type Counter struct {
	one, zero int
}

func (c *Counter) Add(val string) {
	switch val {
	case "1":
		c.one++
	case "0":
		c.zero++
	default:
		panic(fmt.Sprint("Something went wrong: ", val))
	}
}

func (c *Counter) GetVals() (string, string) {
	if c.one > c.zero {
		return "1", "0"
	}
	return "0", "1"
}

func NewReport() *Report {
	return &Report{make(map[int]*Counter), 0}
}

type Report struct {
	Count  map[int]*Counter
	digits int
}

func (r *Report) Add(entry string) {
	var counter *Counter
	var ok bool
	for i, e := range entry {
		if counter, ok = r.Count[i]; !ok {
			r.digits++
			r.Count[i] = &Counter{}
			counter = r.Count[i]
		}
		counter.Add(string(e))
	}
}

func (r *Report) Crunch() (string, string) {
	gamma, epsilon := "", ""
	for i := 0; i < r.digits; i++ {
		g, e := r.Count[i].GetVals()
		gamma = gamma + g
		epsilon = epsilon + e
	}
	return gamma, epsilon
}

func CrunchDiag(diagnostic []string) (string, string) {
	report := NewReport()

	for _, entry := range diagnostic {
		report.Add(entry)
	}
	return report.Crunch()
}

func GetPower(gamma, epsilon string) (int, error) {

	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return 0, err
	}

	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return 0, err
	}

	return int(g * e), nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprint("Could not open file: ", err))
	}
	defer file.Close()

	diagReport := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		diagReport = append(diagReport, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprint("Error scanning file: ", err))
	}

	gamma, epsilon := CrunchDiag(diagReport)
	power, _ := GetPower(gamma, epsilon)

	fmt.Println("Power is: ", power)
}
