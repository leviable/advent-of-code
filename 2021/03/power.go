package main

import (
	"fmt"
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
	return &Report{make(map[int]*Counter)}
}

type Report struct {
	Count map[int]*Counter
}

func (r *Report) Add(entry string) {
	var counter *Counter
	var ok bool
	for i, e := range entry {
		if counter, ok = r.Count[i]; !ok {
			r.Count[i] = &Counter{}
			counter = r.Count[i]
		}
		counter.Add(string(e))
	}
}

func (r *Report) Crunch() (string, string) {
	gamma, epsilon := "", ""
	for i := 0; i < 5; i++ {
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
