package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp/syntax"
	"strconv"
)

const (
	filename  = "day3/input.txt"
	dimension = 140
)

type LookUpMap map[int]map[int]struct{}

type Gear struct {
	X        int
	Y        int
	Adjacent []*Number
}

func (g *Gear) GetLookUp() LookUpMap {
	lookUp := NewLookUpMap()
	return AddToLookUp(g.X, g.Y, lookUp)
}

func (g *Gear) AddAdjacent(numbers []*Number) {
	lookUp := g.GetLookUp()
	for _, n := range numbers {
		if n.IsMatch(lookUp) {
			g.Adjacent = append(g.Adjacent, n)
		}
	}
}

func (g *Gear) IsCorrect() bool {
	return len(g.Adjacent) == 2
}

func (g *Gear) GetRatio() int {
	if !g.IsCorrect() {
		return 0
	}

	ratio := 1
	for _, n := range g.Adjacent {
		ratio *= n.Value
	}

	return ratio
}

type Number struct {
	Value int
	X     int
	Y     []int
}

func (n *Number) IsMatch(lookUp LookUpMap) bool {
	for _, y := range n.Y {
		if _, exists := lookUp[n.X][y]; exists {
			return true
		}
	}
	return false
}

type Digits struct {
	Value []rune
	X     int
	Y     []int
}

func (d *Digits) IsEmpty() bool {
	return len(d.Value) == 0
}

func (d *Digits) AddDigit(j int, r rune) {
	d.Y = append(d.Y, j)
	d.Value = append(d.Value, r)
}

func (d *Digits) Reset() {
	d.Y = []int{}
	d.Value = []rune{}
}

func (d *Digits) BuildNumber() *Number {
	value, err := strconv.Atoi(string(d.Value))
	if err != nil {
		log.Fatal(err)
	}
	return &Number{
		Value: value,
		X:     d.X,
		Y:     d.Y,
	}
}

func NewLookUpMap() LookUpMap {
	lookUp := make(LookUpMap)
	for i := 0; i < dimension; i++ {
		lookUp[i] = make(map[int]struct{})
	}
	return lookUp
}

func AddToLookUp(i, j int, lookUp LookUpMap) LookUpMap {
	for di := i - 1; di <= i+1; di++ {
		if di < 0 || di >= dimension {
			continue
		}
		for dj := j - 1; dj <= j+1; dj++ {
			if dj < 0 || dj >= dimension {
				continue
			}
			lookUp[di][dj] = struct{}{}
		}
	}
	return lookUp
}

func ParseRow(i int, row string, numbers []*Number, gears []*Gear, lookUp LookUpMap) ([]*Number, []*Gear, LookUpMap) {
	ds := &Digits{X: i}

	for j, r := range row {
		if syntax.IsWordChar(r) {
			ds.AddDigit(j, r)
			continue
		}

		if !ds.IsEmpty() {
			numbers = append(numbers, ds.BuildNumber())
			ds.Reset()
		}

		if r != '.' {
			if r == '*' {
				gears = append(gears, &Gear{X: i, Y: j})
			}
			lookUp = AddToLookUp(i, j, lookUp)
		}
	}

	if !ds.IsEmpty() {
		numbers = append(numbers, ds.BuildNumber())
	}

	return numbers, gears, lookUp
}

func ParseInput() ([]*Number, []*Gear, LookUpMap) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := []*Number{}
	gears := []*Gear{}
	lookUp := NewLookUpMap()

	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		numbers, gears, lookUp = ParseRow(i, row, numbers, gears, lookUp)
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers, gears, lookUp
}

func SumNumbers(numbers []*Number, lookUp LookUpMap) int {
	var sum int

	for _, n := range numbers {
		if n.IsMatch(lookUp) {
			sum += n.Value
		}
	}

	return sum
}

func SumGears(numbers []*Number, gears []*Gear) int {
	var sum int

	for _, g := range gears {
		g.AddAdjacent(numbers)
	}

	for _, g := range gears {
		sum += g.GetRatio()
	}

	return sum
}

func main() {
	numbers, gears, globalLookUp := ParseInput()

	sum := SumNumbers(numbers, globalLookUp)
	ratio := SumGears(numbers, gears)

	fmt.Println(sum)
	fmt.Println(ratio)
}

// 530495
// 80253814
