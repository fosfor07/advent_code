package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

	x1, x2, x3, y1, y2, y3 := 0, 0, 0, 0, 0, 0
	// part 1
	houses1 := make(map[string]int)
	address1 := fmt.Sprintf("x%dy%d", x1, y1)
	houses1[address1] = 1
	// part 2
	houses2 := make(map[string]int)
	address2 := fmt.Sprintf("x%dy%d", x2, y2)
	houses2[address2] = 2
	roboMove := false

	for _, line := range lines {
		for _, ch := range line {
			if ch == '^' {
				y1++
				if roboMove {
					y3++
				} else {
					y2++
				}
			} else if ch == 'v' {
				y1--
				if roboMove {
					y3--
				} else {
					y2--
				}
			} else if ch == '>' {
				x1++
				if roboMove {
					x3++
				} else {
					x2++
				}
			} else if ch == '<' {
				x1--
				if roboMove {
					x3--
				} else {
					x2--
				}
			}

			address1 = fmt.Sprintf("x%dy%d", x1, y1)
			if _, ok := houses1[address1]; ok {
				houses1[address1]++
			} else {
				houses1[address1] = 1
			}

			if roboMove {
				address2 = fmt.Sprintf("x%dy%d", x3, y3)
				roboMove = false
			} else {
				address2 = fmt.Sprintf("x%dy%d", x2, y2)
				roboMove = true
			}

			if _, ok := houses2[address2]; ok {
				houses2[address2]++
			} else {
				houses2[address2] = 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", len(houses1))
	fmt.Printf("Part 2: %d\n", len(houses2))
}
