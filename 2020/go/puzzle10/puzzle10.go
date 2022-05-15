package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
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

	ints = append(ints, 0)
	sort.Ints(ints)
	ints = append(ints, ints[len(ints)-1]+3)

	part1, part2 := 0, 0
	d1, d3 := 0, 0
	ways := []int{1}

	for i := 1; i < len(ints); i++ {
		d := ints[i] - ints[i-1]
		if d == 1 {
			d1++
		} else if d == 3 {
			d3++
		}

		ways = append(ways, 0)
		for j := 0; j < i; j++ {
			if ints[i]-ints[j] <= 3 {
				ways[i] += ways[j]
			}
		}
	}

	part1 = d1 * d3
	part2 = ways[len(ways)-1]
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
