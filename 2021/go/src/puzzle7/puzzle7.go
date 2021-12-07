package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
	"utils"
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

	var fuelSpent1 []int
	var fuelSpent2 []int

	for _, p1 := range ints {
		fuel1, fuel2 := 0, 0
		for _, p2 := range ints {
			fuel1 += utils.Abs(p1 - p2)
			for d := 1; d <= utils.Abs(p1-p2); d++ {
				fuel2 += d
			}
		}

		fuelSpent1 = append(fuelSpent1, fuel1)
		fuelSpent2 = append(fuelSpent2, fuel2)
	}

	part1, part2 := 0, 0

	for _, f := range fuelSpent1 {
		if f < part1 || part1 == 0 {
			part1 = f
		}
	}

	for _, f := range fuelSpent2 {
		if f < part2 || part2 == 0 {
			part2 = f
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
