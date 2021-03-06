package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction int
type problemPart int

const (
	DIAGONAL direction = iota
	HORIZONTAL
	VERTICAL

	partOne problemPart = iota
	partTwo
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
		if line == "" {
			continue
		}
		l, err := NewLine(line)
		if err != nil {
			return []Line{}, err
		}
		lines = append(lines, l)
	}
	return lines, nil
}

func NewGrid(lines []Line, part problemPart) *Grid {
	origin := Point{100_000, 100_000}
	end := Point{0, 0}
	grid := make(map[Point]int)
	return &Grid{origin, end, lines, grid, part}
}

type Grid struct {
	origin, end Point
	lines       []Line
	grid        map[Point]int
	part        problemPart
}

func (g *Grid) findBoundaries() {
	for _, line := range g.lines {
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
}

func (g *Grid) traceLines() {
	for _, line := range g.lines {
		for _, point := range getAllPoints(line, g.part) {
			if _, ok := g.grid[point]; !ok {
				g.grid[point] = 1
			} else {
				g.grid[point] += 1
			}
		}
	}
}

func getDirection(line Line) direction {
	if line.begin.x == line.end.x {
		return VERTICAL
	} else if line.begin.y == line.end.y {
		return HORIZONTAL
	} else {
		return DIAGONAL
	}

}

func getAllPoints(line Line, part problemPart) []Point {
	var start, end int
	points := make([]Point, 0)
	switch getDirection(line) {
	case HORIZONTAL:
		if line.begin.x < line.end.x {
			start = line.begin.x
			end = line.end.x
		} else {
			start = line.end.x
			end = line.begin.x
		}
		for x := start; x < end+1; x++ {
			points = append(points, Point{x, line.begin.y})
		}
	case VERTICAL:
		if line.begin.y < line.end.y {
			start = line.begin.y
			end = line.end.y
		} else {
			start = line.end.y
			end = line.begin.y
		}
		for y := start; y < end+1; y++ {
			points = append(points, Point{line.begin.x, y})
		}
	case DIAGONAL:
		if part != partTwo {
			break
		}
		// Need to see if it's at exactly 45 degrees
		// If the horizontal length == vertical length, ????
		xDelta := line.end.x - line.begin.x
		yDelta := line.end.y - line.begin.y
		if abs(xDelta) != abs(yDelta) {
			return points
		}

		points = append(points, Point{line.begin.x, line.begin.y})
		nextX := next(line.begin.x, xDelta)
		nextY := next(line.begin.y, yDelta)

		for len(points) < abs(xDelta)+1 {
			points = append(points, Point{nextX(), nextY()})
		}
	}
	return points
}

func next(start, change int) func() int {
	val := start
	if change < 0 {
		change = -1
	} else {
		change = 1
	}
	return func() int {
		val += change
		return val
	}

}

func (g *Grid) DangerScore() (score int) {
	for _, v := range g.grid {
		if v > 1 {
			score++
		}
	}
	return
}

func GetGrid(lines []Line, part problemPart) *Grid {

	g := NewGrid(lines, part)
	g.findBoundaries()
	g.traceLines()

	return g
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprint("Could not open file: ", err))
	}
	defer file.Close()

	ventData := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ventData += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprint("Error scanning file: ", err))
	}

	lines, err := LoadInput(ventData)
	if err != nil {
		panic(fmt.Sprint("Got an error while loading input data: ", err))
	}

	grid := GetGrid(lines, partOne)
	fmt.Println("The danger score for part one is: ", grid.DangerScore())

	grid = GetGrid(lines, partTwo)
	fmt.Println("The danger score for part two is: ", grid.DangerScore())
}
