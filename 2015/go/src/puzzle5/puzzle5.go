package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"regexp"
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

	part1, part2 := 0, 0

	r := regexp.MustCompile("ab|cd|pq|xy")

	// part 1
	for _, line := range lines {
		if r.MatchString(line) {
			continue
		}

		cond1, cond2 := false, false
		vowels := 0

		for i, ch := range line {
			if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
				vowels++
				if vowels >= 3 {
					cond1 = true
				}
			}

			if i > 0 && line[i] == line[i-1] {
				cond2 = true
			}

			if cond1 && cond2 {
				part1++
				break
			}
		}
	}

	// part 2
	for _, line := range lines {
		cond3, cond4 := false, false

		for i := range line {
			for j := range line {
				if i > 0 && j > 0 && i != j && i != j-1 && j != i-1 && line[i] == line[j] && line[i-1] == line[j-1] {
					cond3 = true
				}
			}

			if i > 1 && line[i] == line[i-2] {
				cond4 = true
			}

			if cond3 && cond4 {
				part2++
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
