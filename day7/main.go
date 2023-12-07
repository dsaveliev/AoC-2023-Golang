package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filename = "day7/input.txt"

var (
	shiftMap = map[string]int{
		"A": 12, "K": 11, "Q": 10, "J": 9, "T": 8, "9": 7, "8": 6, "7": 5, "6": 4, "5": 3, "4": 2, "3": 1, "2": 0}
	shiftMapWithJocker = map[string]int{
		"A": 12, "K": 11, "Q": 10, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1, "J": 0}
)

type Hand struct {
	Cards              []string
	Bid                int
	Strength           int
	StrengthWithJocker int
}

func NewHand(cards []string, bid int) *Hand {
	hand := &Hand{
		Cards: cards,
		Bid:   bid,
	}
	hand.BuildStrength()

	return hand
}

func (h *Hand) BuildStrength() {
	h.Strength = h.GetType() * int(math.Pow10(10))
	h.StrengthWithJocker = h.GetTypeWithJocker() * int(math.Pow10(10))

	for i, c := range h.Cards {
		h.Strength += shiftMap[c] * int(math.Pow10(8-2*i))
		h.StrengthWithJocker += shiftMapWithJocker[c] * int(math.Pow10(8-2*i))
	}
}

func (h *Hand) GetType() int {
	m := map[string]int{}
	for _, c := range h.Cards {
		m[c]++
	}

	values := []int{}
	for _, v := range m {
		values = append(values, v)
	}
	sort.Ints(values)

	pattern := ""
	for _, v := range values {
		pattern += strconv.Itoa(v)
	}

	return ScorePattern(pattern)
}

func (h *Hand) GetTypeWithJocker() int {
	m := map[string]int{}
	j := 0
	for _, c := range h.Cards {
		if c == "J" {
			j++
			continue
		}
		m[c]++
	}

	values := []int{}
	if len(m) != 0 {
		for _, v := range m {
			values = append(values, v)
		}
		sort.Ints(values)
		values[len(values)-1] += j
	} else {
		values = append(values, j)
	}

	pattern := ""
	for _, v := range values {
		pattern += strconv.Itoa(v)
	}

	return ScorePattern(pattern)
}

func ScorePattern(pattern string) int {
	switch pattern {
	case "5":
		return 6
	case "14":
		return 5
	case "23":
		return 4
	case "113":
		return 3
	case "122":
		return 2
	case "1112":
		return 1
	case "11111":
		return 0
	default:
		return 0
	}
}

func ParseHand(row string) *Hand {
	values := strings.Split(row, " ")
	cards := strings.Split(values[0], "")

	bid, err := strconv.Atoi(values[1])
	if err != nil {
		log.Fatal(err)
	}

	return NewHand(cards, bid)
}

func ParseInput() []*Hand {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := []*Hand{}
	for scanner.Scan() {
		row := scanner.Text()

		hands = append(hands, ParseHand(row))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return hands
}

func SumBids(hands []*Hand) int {
	var bidSum int

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].Strength < hands[j].Strength
	})

	for i, h := range hands {
		bidSum += h.Bid * (i + 1)
	}

	return bidSum
}

func SumBidsWithJocker(hands []*Hand) int {
	var bidSum int

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].StrengthWithJocker < hands[j].StrengthWithJocker
	})

	for i, h := range hands {
		bidSum += h.Bid * (i + 1)
	}

	return bidSum
}

func main() {
	hands := ParseInput()
	for _, h := range hands {
		h.BuildStrength()
	}
	bids := SumBids(hands)
	bidsWithJocker := SumBidsWithJocker(hands)

	fmt.Println(bids)
	fmt.Println(bidsWithJocker)
}

// 251216224
// 250825971
