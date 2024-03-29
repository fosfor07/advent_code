package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

type seatType struct {
	r, c int
}

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
	seats := make(map[int]seatType)

	for _, line := range lines {
		rL, rU := 0, 127  // rows range
		cL, cU := 0, 7    // columns range
		rN := rU - rL + 1 // number of rows in range
		cN := cU - cL + 1 // number of columns in range

		for _, ch := range line {
			if ch == 'F' { // front
				rU = rU - rN/2
				rN = rU - rL + 1
			} else if ch == 'B' { // back
				rL = rL + rN/2
				rN = rU - rL + 1
			} else if ch == 'L' { // left
				cU = cU - cN/2
				cN = cU - cL + 1
			} else if ch == 'R' { //right
				cL = cL + cN/2
				cN = cU - cL + 1
			}
		}

		// seat ID: multiply the row by 8, then add the column
		sID := rL*8 + cL
		if sID > part1 {
			part1 = sID
		}

		seats[sID] = seatType{r: rL, c: cL}
	}

	// find missing seat
	for r := 0; r <= 127; r++ {
		for c := 0; c <= 7; c++ {
			sID := r*8 + c
			if _, ok := seats[sID]; !ok {
				// check seats with adjacent IDs
				if _, ok1 := seats[sID-1]; ok1 {
					if _, ok2 := seats[sID+1]; ok2 {
						part2 = sID
						break
					}
				}
			}
		}
		if part2 != 0 {
			break
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
