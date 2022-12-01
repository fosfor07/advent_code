package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
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

	cals := []int{}
	c := 0

	for _, line := range lines {
		if line == "" {
			cals = append(cals, c)
			c = 0
		} else {
			n, _ := strconv.Atoi(line)
			c += n
		}
	}
	cals = append(cals, c)

	sort.Ints(cals)

	part1 := cals[len(cals)-1]
	part2 := cals[len(cals)-1] + cals[len(cals)-2] + cals[len(cals)-3]

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
