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

	depths, err := readers.ReadInts(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	part1, part2 := 0, 0

	for i := range depths {
		if i > 0 && depths[i-1] < depths[i] {
			part1++
		}
	}

	for j := 0; j < len(depths)-3; j++ {
		if depths[j]+depths[j+1]+depths[j+2] < depths[j+1]+depths[j+2]+depths[j+3] {
			part2++
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
