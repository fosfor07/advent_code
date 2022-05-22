package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	ints, err := readers.ReadInts(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	part1 := 0
	for _, num := range ints {
		part1 += num
	}
	fmt.Printf("Part 1: %d\n", part1)

	part2 := 0
	freqs := make(map[int]bool)
	found := false

	for !found {
		for _, num := range ints {
			part2 += num

			if freqs[part2] {
				found = true
				break
			} else {
				freqs[part2] = true
			}
		}
	}
	fmt.Printf("Part 2: %d\n", part2)
}
