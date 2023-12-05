package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

const filename = "day5/input.txt"

type (
	Seeds struct {
		Chains [][]int
	}

	Range struct {
		Destination int
		Source      int
		Range       int
	}

	RangeMap struct {
		Title          string
		Ranges         []*Range
		DestinationMap *RangeMap
	}
)

func (s *Seeds) AddToTheChain(value, valueRange int) {
	s.Chains = append(s.Chains, []int{value, valueRange})
}

func (m *RangeMap) GetDestination(source int) int {
	d := source
	for _, r := range m.Ranges {
		if r.Source <= source && source < r.Source+r.Range {
			d = r.Destination + (source - r.Source)
		}
	}

	if m.DestinationMap != nil {
		return m.DestinationMap.GetDestination(d)
	}

	return d
}

func ParseSeeds(row string) *Seeds {
	row, found := strings.CutPrefix(row, "seeds: ")
	if !found {
		log.Fatal(errors.New("seeds prefix isn't found"))
	}
	rowSeeds := strings.Split(row, " ")

	seeds := &Seeds{}
	for _, rs := range rowSeeds {
		value, err := strconv.Atoi(rs)
		if err != nil {
			log.Fatal(err)
		}
		seeds.AddToTheChain(value, 1)
	}

	return seeds
}

func ParseAnotherSeeds(row string) *Seeds {
	row, found := strings.CutPrefix(row, "seeds: ")
	if !found {
		log.Fatal(errors.New("seeds prefix isn't found"))
	}
	rowSeeds := strings.Split(row, " ")

	seeds := &Seeds{}
	for i := 0; i < len(rowSeeds); i += 2 {
		value, err := strconv.Atoi(rowSeeds[i])
		if err != nil {
			log.Fatal(err)
		}
		valueRange, err := strconv.Atoi(rowSeeds[i+1])
		if err != nil {
			log.Fatal(err)
		}
		seeds.AddToTheChain(value, valueRange)
	}

	return seeds
}

func ParseRange(row string) *Range {
	values := make([]int, 3)
	for i, rs := range strings.Split(row, " ") {
		s, err := strconv.Atoi(rs)
		if err != nil {
			log.Fatal(err)
		}
		values[i] = s
	}

	return &Range{
		Destination: values[0],
		Source:      values[1],
		Range:       values[2],
	}
}

func ParseInput() (*RangeMap, *Seeds, *Seeds) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var seeds *Seeds
	var anotherSeeds *Seeds
	var headMap *RangeMap
	var currentMap *RangeMap
	for scanner.Scan() {
		row := scanner.Text()

		if strings.Contains(row, "seeds") {
			seeds = ParseSeeds(row)
			anotherSeeds = ParseAnotherSeeds(row)
			continue
		}

		if strings.Contains(row, "map") {
			if currentMap == nil {
				headMap = &RangeMap{Title: row}
				currentMap = headMap
			} else {
				currentMap.DestinationMap = &RangeMap{Title: row}
				currentMap = currentMap.DestinationMap
			}
			continue
		}

		if row == "" {
			continue
		}

		currentMap.Ranges = append(currentMap.Ranges, ParseRange(row))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return headMap, seeds, anotherSeeds
}

func FindLowestLocation(seeds *Seeds, headMap *RangeMap) int {
	locations := make(chan int, len(seeds.Chains))

	var wg sync.WaitGroup
	for _, c := range seeds.Chains {
		wg.Add(1)

		go func(c []int) {
			defer wg.Done()

			min := -1
			for s := c[0]; s < c[0]+c[1]; s++ {
				ll := headMap.GetDestination(s)
				if ll < min || min == -1 {
					min = ll
				}
			}
			locations <- min
		}(c)
	}

	wg.Wait()

	min := -1
	for range seeds.Chains {
		ll := <-locations
		if ll < min || min == -1 {
			min = ll
		}
	}

	return min
}

func PrintOutRangeMaps(m *RangeMap) {
	fmt.Println()
	fmt.Printf("%#v\n", m.Title)
	for _, r := range m.Ranges {
		fmt.Printf("%#v\n", *r)
	}
	if m.DestinationMap != nil {
		PrintOutRangeMaps(m.DestinationMap)
	}
}

func PrintOutChains(seeds *Seeds) {
	fmt.Println()
	for _, c := range seeds.Chains {
		fmt.Println(c)
	}
}

func main() {
	headMap, seeds, anotherSeeds := ParseInput()

	fmt.Println(FindLowestLocation(seeds, headMap))
	fmt.Println(FindLowestLocation(anotherSeeds, headMap))

	// PrintOutRangeMaps(headMap)
	// PrintOutChains(seedsRanged)
}

// 173706076
// 11611182
