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
	return &Report{
		Count:   make(map[int]*Counter),
		Entries: make([]string, 0),
		digits:  0,
	}
}

type Report struct {
	Count   map[int]*Counter
	Entries []string
	digits  int
}

func (r *Report) Add(entry string) {
	var counter *Counter
	var ok bool

	r.Entries = append(r.Entries, entry)
	for i, e := range entry {
		if counter, ok = r.Count[i]; !ok {
			r.digits++
			r.Count[i] = &Counter{}
			counter = r.Count[i]
		}
		counter.Add(string(e))
	}
}

func (r *Report) Crunch() (gamma string, epsilon string, o2 string, co2 string) {
	for i := 0; i < r.digits; i++ {
		g, e := r.Count[i].GetVals()
		gamma = gamma + g
		epsilon = epsilon + e
	}

	o2Selector := func(one, zero int) (use string) {
		use = "1"
		if zero > one {
			use = "0"
		}
		return
	}
	o2 = r.calcRating(o2Selector, r.Entries[:])

	co2Selector := func(one, zero int) (use string) {
		use = "0"
		if one < zero {
			use = "1"
		}
		return
	}
	co2 = r.calcRating(co2Selector, r.Entries[:])

	return
}

func (r *Report) calcRating(selector func(int, int) string, entries []string) string {

	for i := 0; i < r.digits; i++ {
		if len(entries) <= 1 {
			break
		}
		one, zero := 0, 0
		for _, e := range entries {
			if string(e[i]) == "1" {
				one++
			} else {
				zero++
			}
		}

		use := selector(one, zero)

		eCopy := entries[:]
		entries = make([]string, 0)
		for _, e := range eCopy {
			if string(e[i]) == use {
				entries = append(entries, e)
			}
		}

	}
	return entries[0]
}

func CrunchDiag(diagnostic []string) (string, string, string, string) {
	report := NewReport()

	for _, entry := range diagnostic {
		report.Add(entry)
	}
	return report.Crunch()
}

func GetTotal(aRaw, bRaw string) (int, error) {

	a, err := strconv.ParseInt(aRaw, 2, 64)
	if err != nil {
		return 0, err
	}

	b, err := strconv.ParseInt(bRaw, 2, 64)
	if err != nil {
		return 0, err
	}

	return int(a * b), nil
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

	gamma, epsilon, o2, co2 := CrunchDiag(diagReport)
	power, _ := GetTotal(gamma, epsilon)
	rating, _ := GetTotal(o2, co2)

	fmt.Println("Power is: ", power)
	fmt.Println("Rating is: ", rating)
}
