package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	lines, err := readers.ReadLines(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	part1, part2 := 0, 0

	program := strings.Join(lines, "")
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	part1 = calculate(program, r)

	subprograms := strings.Split(program, `don't()`)
	for i, subprogram := range subprograms {
		if i == 0 {
			part2 += calculate(subprogram, r)
		} else {
			commands := strings.Split(subprogram, `do()`)
			for j, cmd := range commands {
				if j > 0 {
					part2 += calculate(cmd, r)
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func calculate(program string, r *regexp.Regexp) int {
	total := 0
	instructions := r.FindAllString(program, -1)

	for _, ins := range instructions {
		multipliedValue := 1

		substrings := strings.Split(ins, ",")
		for _, s := range substrings {
			strFactor := strings.Trim(s, "mul()")
			factor, _ := strconv.Atoi(strFactor)
			multipliedValue *= factor
		}

		total += multipliedValue
	}
	return total
}
