package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const filename = "day1/input.txt"

var (
	rNumeric      = regexp.MustCompile(`(\d)`)
	rAlphanumeric = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	mAlphanumeric = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
)

type Parser struct {
	regexp *regexp.Regexp
}

func NewParser(regexp *regexp.Regexp) Parser {
	return Parser{regexp: regexp}
}

func (p Parser) FirstDigit(row string) int {
	sm := p.regexp.FindStringSubmatch(row)
	return ParseDigit(sm[0])
}

func (p Parser) LastDigit(row string) int {
	length := len(row)
	for i := 1; i <= length; i++ {
		sm := p.regexp.FindStringSubmatch(row[length-i : length])
		if len(sm) == 0 {
			continue
		}
		return ParseDigit(sm[0])
	}
	return 0
}

func (p Parser) Sum() int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		row := scanner.Text()
		sum += 10*p.FirstDigit(row) + p.LastDigit(row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func ParseDigit(s string) int {
	if len(s) == 1 {
		d, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return d
	}
	return mAlphanumeric[s]
}

func main() {
	fmt.Println(NewParser(rNumeric).Sum())
	fmt.Println(NewParser(rAlphanumeric).Sum())
}

// 54916
// 54728
