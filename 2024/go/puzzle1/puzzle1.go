package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
	"github.com/fosfor07/advent_code/pkg/utils"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	lines, err := readers.ReadLines(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	part1, part2 := 0, 0
	var list1 []int
	var list2 []int

	for _, line := range lines {
		numbers := strings.Fields(line)
		for i, str := range numbers {
			n, _ := strconv.Atoi(string(str))

			if i == 0 {
				list1 = append(list1, n)
			} else if i == 1 {
				list2 = append(list2, n)
			}
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	for i, l1 := range list1 {
		part1 += utils.Abs(list1[i] - list2[i])

		count := 0
		for _, l2 := range list2 {
			if l1 == l2 {
				count++
			}
		}
		part2 += l1 * count
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
