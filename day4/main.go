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

const filename = "day4/input.txt"

var r = regexp.MustCompile(`^Card[ ]+(\d+):([ \d]+) \| ([ \d]+)`)

type (
	Card struct {
		ID      int
		Winning map[int]struct{}
		Actual  []int
	}

	Deck []*Card
)

func (c *Card) GetPoints() int {
	points := 0
	for _, n := range c.Actual {
		if _, exists := c.Winning[n]; exists {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func (c *Card) GetMatches() int {
	matches := 0
	for _, n := range c.Actual {
		if _, exists := c.Winning[n]; exists {
			matches += 1
		}
	}
	return matches
}

func ParseRow(row string) *Card {
	sm := r.FindAllStringSubmatch(row, -1)

	id, err := strconv.Atoi(sm[0][1])
	if err != nil {
		log.Fatal(err)
	}

	winning := map[int]struct{}{}
	for _, s := range strings.Split(sm[0][2], " ") {
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		winning[n] = struct{}{}
	}

	actual := []int{}
	for _, s := range strings.Split(sm[0][3], " ") {
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		actual = append(actual, n)
	}

	return &Card{
		ID:      id,
		Winning: winning,
		Actual:  actual,
	}
}

func ParseInput() Deck {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := Deck{}
	for scanner.Scan() {
		row := scanner.Text()
		cards = append(cards, ParseRow(row))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cards
}

func SumPoints(cards Deck) int {
	sum := 0
	for _, c := range cards {
		sum += c.GetPoints()
	}
	return sum
}

func GenerateCopies(matches int, cards Deck, result *Deck) {
	for k := 0; k < matches; k++ {
		*result = append(*result, cards[k])
		m := cards[k].GetMatches()
		GenerateCopies(m, cards[k+1:], result)
	}
}

func main() {
	cards := ParseInput()
	sum := SumPoints(cards)

	result := Deck{}
	GenerateCopies(len(cards), cards, &result)

	fmt.Println(sum)
	fmt.Println(len(result))
}

// 25004
// 14427616
