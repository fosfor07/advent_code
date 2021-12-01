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

	part1, part2, prevDepth := 0, 0, 0

	for _, depth := range depths {
		if prevDepth < depth && prevDepth != 0 {
			part1++
		}
		prevDepth = depth
	}

	for i := 0; i < len(depths)-3; i++ {
		if depths[i]+depths[i+1]+depths[i+2] < depths[i+1]+depths[i+2]+depths[i+3] {
			part2++
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
