package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
)

type seatType struct {
	r int
	c int
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
	rL, rU := 0, 127  // rows range
	cL, cU := 0, 7    // columns range
	rD := rU - rL + 1 // number of rows in range
	cD := cU - cL + 1 // number of columns in range

	for _, line := range lines {
		for _, ch := range line {
			if ch == 'F' { // front
				rU = rU - rD/2
				rD = rU - rL + 1
			}
			if ch == 'B' { // back
				rL = rL + rD/2
				rD = rU - rL + 1
			}
			if ch == 'L' { // left
				cU = cU - cD/2
				cD = cU - cL + 1
			}
			if ch == 'R' { //right
				cL = cL + cD/2
				cD = cU - cL + 1
			}
		}

		// seat ID: multiply the row by 8, then add the column
		sID := rL*8 + cL

		var seat seatType
		seat.r = rL
		seat.c = cL
		seats[sID] = seat

		if sID > part1 {
			part1 = sID
		}

		rL, rU = 0, 127
		cL, cU = 0, 7
		rD = rU - rL + 1
		cD = cU - cL + 1
	}

	// find missing seat
	for r := 0; r < 128; r++ {
		for c := 0; c < 8; c++ {
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
