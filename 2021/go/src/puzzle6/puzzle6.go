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

	input = bytes.TrimSpace(input)
	ints, err := readers.ReadCsvInts(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	part1, part2 := 0, 0

	var fishes [9]int

	for _, n := range ints {
		fishes[n]++
	}

	for d := 0; d < 256; d++ {
		value0 := fishes[0]

		for i := 0; i < 8; i++ {
			fishes[i] = fishes[i+1]
		}
		fishes[6] += value0
		fishes[8] = value0

		if d == 79 {
			for _, f := range fishes {
				part1 += f
			}
		}
	}

	for _, f := range fishes {
		part2 += f
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
