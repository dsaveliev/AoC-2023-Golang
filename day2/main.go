package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	filename = "day2/input.txt"
)

var (
	r     = regexp.MustCompile(`^Game (\d+): (.*)$`)
	total = Attempt{Red: 12, Green: 13, Blue: 14}
)

type (
	Attempt struct {
		Red   int
		Green int
		Blue  int
	}

	Game struct {
		ID       int
		Attempts []Attempt
	}
)

func (a Attempt) Sum() int {
	return a.Red + a.Green + a.Blue
}

func (a Attempt) IsValid() bool {
	if a.Sum() > total.Sum() ||
		a.Red > total.Red ||
		a.Green > total.Green ||
		a.Blue > total.Blue {
		return false
	}
	return true
}

func (g Game) IsValid() bool {
	for _, a := range g.Attempts {
		if !a.IsValid() {
			return false
		}
	}
	return true
}

func (g Game) Power() int {
	var power Attempt

	for _, a := range g.Attempts {
		if a.Red > power.Red {
			power.Red = a.Red
		}
		if a.Green > power.Green {
			power.Green = a.Green
		}
		if a.Blue > power.Blue {
			power.Blue = a.Blue
		}
	}

	return power.Red * power.Green * power.Blue
}

func ParseAttempt(row string) Attempt {
	var red, green, blue int
	var err error

	colors := strings.Split(row, ", ")
	for _, c := range colors {
		cubes := strings.Split(c, " ")
		switch cubes[1] {
		case "red":
			red, err = strconv.Atoi(cubes[0])
			if err != nil {
				log.Fatal(err)
			}
		case "green":
			green, err = strconv.Atoi(cubes[0])
			if err != nil {
				log.Fatal(err)
			}
		case "blue":
			blue, err = strconv.Atoi(cubes[0])
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return Attempt{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func ParseGame(row string) Game {
	sm := r.FindAllStringSubmatch(row, -1)

	id, err := strconv.Atoi(sm[0][1])
	if err != nil {
		log.Fatal(err)
	}
	input := sm[0][2]

	rawAttempts := strings.Split(input, "; ")
	attempts := make([]Attempt, len(rawAttempts))
	for i, ra := range rawAttempts {
		attempts[i] = ParseAttempt(ra)
	}

	return Game{
		ID:       id,
		Attempts: attempts,
	}
}

func ParseInput() []Game {
	games := []Game{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		games = append(games, ParseGame(row))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return games
}

func CalculateCheckSum(games []Game) int {
	var sum int

	for _, g := range games {
		if g.IsValid() {
			sum += g.ID
		}
	}

	return sum
}

func CalculatePower(games []Game) int {
	var sum int

	for _, g := range games {
		sum += g.Power()
	}

	return sum
}

func main() {
	games := ParseInput()

	sum := CalculateCheckSum(games)
	power := CalculatePower(games)

	fmt.Println(sum)
	fmt.Println(power)
}

// 2593
// 54699
