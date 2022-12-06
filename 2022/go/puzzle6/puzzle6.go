package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

func verify(marker []byte) bool {
	for i := 0; i < len(marker)-1; i++ {
		for j := i + 1; j < len(marker); j++ {
			if marker[i] == marker[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := readers.ReadLines(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	part1, part2 := 0, 0

	// part 1
	for i := 4; i < len(lines[0])-1; i++ {
		marker := []byte{}
		for j := i - 3; j <= i; j++ {
			marker = append(marker, lines[0][j])
		}

		if verify(marker) {
			part1 = i + 1
			break
		}
	}

	// part 2
	for i := 14; i < len(lines[0])-1; i++ {
		marker := []byte{}
		for j := i - 13; j <= i; j++ {
			marker = append(marker, lines[0][j])
		}

		if verify(marker) {
			part2 = i + 1
			break
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
