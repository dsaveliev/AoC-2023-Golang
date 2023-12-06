package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const filename = "day6/input.txt"

type Race struct {
	Time     int
	Distance int
}

func (r *Race) BuildSolutions() []int {
	solutions := []int{}

	for t := 0; t <= r.Time; t++ {
		s := r.Time*t - t*t
		if s > r.Distance {
			solutions = append(solutions, s)
		}
	}

	return solutions
}

func MultiplySolutions(races []*Race) int {
	result := 1
	for _, r := range races {
		result *= len(r.BuildSolutions())
	}
	return result
}

func ParseRow(prefix, row string) []int {
	result := []int{}

	row, found := strings.CutPrefix(row, prefix)
	if !found {
		log.Fatal(errors.New("prefix isn't found"))
	}
	rowSeeds := strings.Split(strings.Trim(row, " "), " ")

	for _, rs := range rowSeeds {
		if rs == "" {
			continue
		}
		value, err := strconv.Atoi(rs)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, value)
	}

	return result
}

func ParseWholeRow(prefix, row string) int {
	row, found := strings.CutPrefix(row, prefix)
	if !found {
		log.Fatal(errors.New("prefix isn't found"))
	}

	row = strings.ReplaceAll(row, " ", "")
	value, err := strconv.Atoi(row)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func BuildRaces(times, distances []int) []*Race {
	races := []*Race{}
	for i := range times {
		races = append(races, &Race{
			Time:     times[i],
			Distance: distances[i],
		})
	}
	return races
}

func ParseInput() ([]*Race, *Race) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var times, distances []int
	var wholeTime, wholeDistance int
	for scanner.Scan() {
		row := scanner.Text()

		if strings.Contains(row, "Time:") {
			times = ParseRow("Time:", row)
			wholeTime = ParseWholeRow("Time:", row)
		}

		if strings.Contains(row, "Distance:") {
			distances = ParseRow("Distance:", row)
			wholeDistance = ParseWholeRow("Distance:", row)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return BuildRaces(times, distances), &Race{Time: wholeTime, Distance: wholeDistance}
}

func main() {
	races, wholeRace := ParseInput()

	fmt.Println(MultiplySolutions(races))
	fmt.Println(len(wholeRace.BuildSolutions()))
}

// 1084752
// 28228952
