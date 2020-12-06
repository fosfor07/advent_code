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
	// Append an empty line to read answers of the last group.
	lines = append(lines, "")

	part1, part2 := 0, 0

	answers := make(map[rune]int)
	ppl := 0

	for _, line := range lines {
		if len(line) > 0 {
			ppl++
			for _, ch := range line {
				if _, ok := answers[ch]; !ok {
					answers[ch] = 1
				} else {
					answers[ch]++
				}
			}
		} else {
			part1 += len(answers)

			for _, answer := range answers {
				if answer == ppl {
					part2++
				}
			}

			answers = make(map[rune]int)
			ppl = 0
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
