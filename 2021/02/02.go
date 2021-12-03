package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type plane int

const (
	FORWARD  = "forward"
	BACKWARD = "backward"
	DOWN     = "down"
	UP       = "up"

	HORIZONTAL plane = iota
	VERTICAL
)

type Command struct {
	Plane  plane
	Amount int
}

type Location struct {
	x, y int
}

func NewSub() Sub {
	return Sub{CurrentLocation: Location{0, 0}, Aim: 0}
}

type Sub struct {
	CurrentLocation Location
	Aim             int
}

func (s *Sub) Final() int {
	return s.CurrentLocation.x * s.CurrentLocation.y
}

func (s *Sub) MoveHorizontal(amount int) {
	s.CurrentLocation.x += amount
}

func (s *Sub) MoveVertical(amount int) {
	s.CurrentLocation.y += amount
}

func (s *Sub) IssueCommands(commands []Command) {
	for _, c := range commands {
		switch c.Plane {
		case HORIZONTAL:
			s.MoveHorizontal(c.Amount)
			if c.Amount > 0 {
				s.MoveVertical(s.Aim * c.Amount)
			}
		case VERTICAL:
			// s.MoveVertical(c.Amount)
			s.Aim += c.Amount
		}
	}
}

func ParseCommand(command string) Command {
	cmd := strings.Split(command, " ")

	var p plane
	switch cmd[0] {
	case FORWARD, BACKWARD:
		p = HORIZONTAL
	case DOWN, UP:
		p = VERTICAL
	}

	amount, _ := strconv.Atoi(cmd[1])

	switch cmd[0] {
	case BACKWARD, UP:
		amount *= -1
	}

	return Command{Plane: p, Amount: amount}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprint("Failed to open file: ", err))
	}
	defer file.Close()

	commands := make([]Command, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, ParseCommand(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprint("Error scanning file: ", err))
	}

	sub := NewSub()
	sub.IssueCommands(commands)

	fmt.Println("Sub final position is: ", sub.Final())
}
