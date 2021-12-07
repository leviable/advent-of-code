package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type spawnSignal int

const (
	NEWFISHTIMER   = 9
	FISHTIMERRESET = 6

	NEWSPAWN spawnSignal = iota
	NOSPAWN
)

func NewLanternFish() *LanternFish {
	return &LanternFish{NEWFISHTIMER}
}

type LanternFish struct {
	timer int
}

func (l *LanternFish) String() string {
	return fmt.Sprintf("fish[%d]", l.timer)
}

func (l *LanternFish) GrowOlder() spawnSignal {
	signal := NOSPAWN
	l.timer--
	if l.timer < 0 {
		l.timer = FISHTIMERRESET
		signal = NEWSPAWN
	}

	return signal
}

func LoadInput(input string) ([]*LanternFish, error) {
	fish := strings.Split(input, ",")
	lf := make([]*LanternFish, len(fish))

	for idx, f := range fish {
		timer, err := strconv.Atoi(f)
		if err != nil {
			return []*LanternFish{}, err
		}
		lf[idx] = &LanternFish{timer}
	}

	return lf, nil
}

func NewSchool(fish []*LanternFish, daysToAge int) *School {
	fMap := make(map[int][]*LanternFish)
	for x := 2; x <= daysToAge; x++ {
		fMap[x] = make([]*LanternFish, 0)
	}
	fMap[1] = fish
	return &School{fMap, daysToAge, len(fish)}
}

type School struct {
	fish map[int][]*LanternFish
	days int
	size int
}

func (s *School) Size() int {
	return s.size
}

func (s *School) track() {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("%s: %d\n", t, s.size)
		}
	}
}

func (s *School) Grow() {
	var (
		fish *LanternFish
		age  int
	)

	go s.track()
	// loop:
	//   Pull the youngest fish from the fish map
	//   Grow it till s.days
	//   Add any new fish to the map and s.size++
	for s.ungrownFish() {
		fish, age = s.getYoungestFish()
		for x := age; x <= s.days; x++ {
			if fish.GrowOlder() == NEWSPAWN {
				s.fish[x] = append(s.fish[x], NewLanternFish())
				s.size++
			}
		}
	}
}

func (s *School) ungrownFish() bool {
	// fmt.Println("In ungrownfish")
	// fmt.Println("Map is: ", s.fish)
	for day := s.days; day >= 1; day-- {
		if len(s.fish[day]) != 0 {
			return true
		}
	}

	return false
}

func (s *School) getYoungestFish() (*LanternFish, int) {
	var (
		fish *LanternFish
		day  int
	)
	for day = s.days; day >= 1; day-- {
		if len(s.fish[day]) == 0 {
			continue
		}
		fish, s.fish[day] = s.fish[day][0], s.fish[day][1:]
		// fmt.Println("*****************")
		// fmt.Printf("On day %d pulled fish\n", day)
		// fmt.Println(s.fish)
		// fmt.Println("*****************")
		break
	}
	return fish, day
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		panic(fmt.Sprint("Could not open file: ", err))
	}
	defer file.Close()

	fishData := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fishData += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprint("Error scanning file: ", err))
	}

	fish, err := LoadInput(fishData)
	if err != nil {
		panic(fmt.Sprint("Got an error and didn't expect it: ", err))
	}

	// school := NewSchool(fish, 80)
	// school.Grow()

	// fmt.Println("After 80 days the school size is: ", school.Size())

	school := NewSchool(fish, 256)
	school.Grow()

	fmt.Println("After 256 days the school size is: ", school.Size())
}
