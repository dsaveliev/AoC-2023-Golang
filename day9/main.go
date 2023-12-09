package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const filename = "day9/input.txt"

func ParseRow(row string) []int {
	rowValues := strings.Split(row, " ")
	values := make([]int, len(rowValues))

	for i, rv := range rowValues {
		v, err := strconv.Atoi(rv)
		if err != nil {
			log.Fatal(err)
		}
		values[i] = v
	}

	return values
}

func ParseInput() [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	history := [][]int{}
	for scanner.Scan() {
		row := scanner.Text()
		history = append(history, ParseRow(row))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return history
}

func GetFirstPrediction(values []int) int {
	l := len(values)
	diffs := make([]int, l-1)
	allZeros := true

	for i := range values {
		if i <= l-2 {
			t := values[i+1] - values[i]
			if t != 0 {
				allZeros = false
			}
			diffs[i] = t
		}
	}

	if allZeros {
		return values[0] - 0
	}

	return values[0] - GetFirstPrediction(diffs)
}

func GetLastPrediction(values []int) int {
	l := len(values)
	diffs := make([]int, l-1)
	allZeros := true

	for i := range values {
		if i <= l-2 {
			t := values[i+1] - values[i]
			if t != 0 {
				allZeros = false
			}
			diffs[i] = t
		}
	}

	if allZeros {
		return values[l-1] + 0
	}

	return values[l-1] + GetLastPrediction(diffs)
}

func GetSums(history [][]int) (int, int) {
	var firstSum, lastSum int

	for _, values := range history {
		lastSum += GetLastPrediction(values)
		firstSum += GetFirstPrediction(values)
	}
	return lastSum, firstSum
}

func main() {
	history := ParseInput()

	fmt.Println(GetSums(history))
}

// 1743490457
// 1053
