package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	lines, err := readers.ReadLines(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	part1, part2, idx := 0, 0, 1

	for _, line := range lines {
		for _, ch := range line {
			if ch == '(' {
				part1++
			} else if ch == ')' {
				part1--
			}

			if part1 == -1 && part2 == 0 {
				part2 = idx
			}
			idx++
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
