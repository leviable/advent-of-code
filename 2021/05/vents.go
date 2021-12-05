package main

import (
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func NewPoint(raw string) (p Point, err error) {
	// Example raw point: 8,0
	var x, y int
	point := strings.Split(raw, ",")
	x, err = strconv.Atoi(point[0])
	if err != nil {
		return p, err
	}
	y, err = strconv.Atoi(point[1])
	if err != nil {
		return p, err
	}

	p.x, p.y = x, y
	return p, nil
}

type Line struct {
	begin, end Point
}

func NewLine(raw string) (Line, error) {
	// Example raw line: 8,0 -> 0,8
	points := strings.Split(raw, " -> ")
	begin, err := NewPoint(points[0])
	if err != nil {
		return Line{}, err
	}
	end, err := NewPoint(points[1])
	if err != nil {
		return Line{}, err
	}
	return Line{begin, end}, nil
}

func LoadInput(input string) ([]Line, error) {
	lines := make([]Line, 0)
	for _, line := range strings.Split(input, "\n") {
		l, err := NewLine(line)
		if err != nil {
			return []Line{}, err
		}
		lines = append(lines, l)
	}
	return lines, nil
}

type Grid struct {
	origin, end Point
}

func GetGrid(lines []Line) Grid {
	origin := Point{100_000, 100_000}

	end := Point{0, 0}

	g := Grid{origin, end}
	for _, line := range lines {
		if line.begin.x < g.origin.x {
			g.origin.x = line.begin.x
		}
		if line.end.x < g.origin.x {
			g.origin.x = line.end.x
		}
		if line.begin.y < g.origin.y {
			g.origin.y = line.begin.y
		}
		if line.end.y < g.origin.y {
			g.origin.y = line.end.y
		}
		if line.end.x > g.end.x {
			g.end.x = line.end.x
		}
		if line.begin.x > g.end.x {
			g.end.x = line.begin.x
		}
		if line.end.y > g.end.y {
			g.end.y = line.end.y
		}
		if line.begin.y > g.end.y {
			g.end.y = line.begin.y
		}
	}
	return g
}
