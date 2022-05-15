package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input string = "19,0,5,1,10,13"
const part1Limit int = 2020
const part2Limit int = 30000000

func main() {
	part1, part2 := 0, 0
	indices := make(map[int]int)
	idx, n := 0, 0

	nums := strings.Split(input, ",")
	for i, nStr := range nums {
		n, _ = strconv.Atoi(nStr)
		indices[n] = i
	}
	idx = indices[n]

	for idx < part2Limit-1 {
		if _, ok := indices[n]; ok {
			prevIdx := indices[n]
			indices[n] = idx
			n = idx - prevIdx
		} else {
			indices[n] = idx
			n = 0
		}

		if idx < part1Limit-1 {
			part1 = n
		}
		idx++
	}
	part2 = n

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
