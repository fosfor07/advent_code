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

	ints, err := readers.ReadInts(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	found1, found2 := false, false
	part1, part2 := 0, 0

	for _, expense1 := range ints {
		for _, expense2 := range ints {
			if expense1+expense2 == 2020 {
				found1 = true
				part1 = expense1 * expense2
			}

			if !found2 {
				for _, expense3 := range ints {
					if expense1+expense2+expense3 == 2020 {
						found2 = true
						part2 = expense1 * expense2 * expense3
						break
					}
				}
			}
		}

		if found1 && found2 {
			break
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
