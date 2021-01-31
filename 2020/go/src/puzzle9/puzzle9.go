package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
	"utils"
)

const pLen int = 25

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

	part1, part2 := 0, 0

	for s := pLen; s < len(ints); s++ {
		fnd := false
		preamble := ints[s-pLen : s]
		for i := 0; i < pLen; i++ {
			for j := 0; j < pLen; j++ {
				if i != j {
					if ints[s] == (preamble[i] + preamble[j]) {
						fnd = true
						break
					}
				}
			}
			if fnd {
				break
			}
		}
		if !fnd {
			part1 = ints[s]
			break
		}
	}

	total, firstIdx, lastIdx := 0, 0, 0
	for i, num1 := range ints {
		total = num1
		for j := i + 1; j < len(ints); j++ {
			total += ints[j]
			if total == part1 {
				firstIdx = i
				lastIdx = j
				break
			} else if total > part1 {
				break
			}
		}
		if total == part1 {
			break
		}
	}

	part2 = utils.Min(ints[firstIdx:lastIdx+1]) + utils.Max(ints[firstIdx:lastIdx+1])

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
