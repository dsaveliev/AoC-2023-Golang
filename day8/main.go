package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const filename = "day8/input.txt"

var re = regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)

func ParseInstructions(row string) []int {
	ss := strings.Split(row, "")
	instructions := make([]int, len(ss))

	for i, s := range ss {
		if s == "L" {
			instructions[i] = 0
		} else {
			instructions[i] = 1
		}
	}

	return instructions
}

func ParseTree(row string) []string {
	sm := re.FindAllStringSubmatch(row, -1)
	return sm[0][1:]
}

func ParseInput() ([]int, map[string][2]string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instructions []int
	tree := map[string][2]string{}
	i := 0
	for scanner.Scan() {
		row := scanner.Text()

		if i == 0 {
			instructions = ParseInstructions(row)
		}

		if i >= 2 {
			t := ParseTree(row)
			tree[t[0]] = [2]string{t[1], t[2]}
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return instructions, tree
}

func Traverse(start string, instructions []int, tree map[string][2]string) int {
	steps := 0
	found := false
	current := start

	for {
		for _, i := range instructions {
			current = tree[current][i]
			steps++
			if current[2] == 'Z' {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	return steps
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func TraverseParallel(instructions []int, tree map[string][2]string) int {
	parallelSteps := []int{}

	for k := range tree {
		if k[2] == 'A' {
			nodeSteps := Traverse(k, instructions, tree)
			parallelSteps = append(parallelSteps, nodeSteps)
		}
	}

	return LCM(parallelSteps...)
}

func main() {
	instructions, tree := ParseInput()

	steps := Traverse("AAA", instructions, tree)
	fmt.Println(steps)

	parallelSteps := TraverseParallel(instructions, tree)
	fmt.Println(parallelSteps)
}

// 17873
// 15746133679061
