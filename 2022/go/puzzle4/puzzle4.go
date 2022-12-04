package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

	part1, part2 := 0, 0

	for _, line := range lines {
		pairs := strings.Split(line, ",")
		sections1 := strings.Split(pairs[0], "-")
		sections2 := strings.Split(pairs[1], "-")

		s1, _ := strconv.Atoi(sections1[0])
		e1, _ := strconv.Atoi(sections1[1])
		s2, _ := strconv.Atoi(sections2[0])
		e2, _ := strconv.Atoi(sections2[1])

		if s1 <= s2 && e1 >= e2 {
			part1++
		} else if s2 <= s1 && e2 >= e1 {
			part1++
		}

		if s1 <= s2 && e1 >= s2 {
			part2++
		} else if s2 <= s1 && e2 >= s1 {
			part2++
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
