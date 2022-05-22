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

	boxIds, err := readers.ReadLines(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	// Part 1
	cnt2, cnt3 := 0, 0

	for _, id := range boxIds {
		found2, found3 := false, false

		for _, curChar := range id {
			numOfChar := strings.Count(id, string(curChar))

			if numOfChar == 2 && found2 == false {
				found2 = true
				cnt2++
			}

			if numOfChar == 3 && found3 == false {
				found3 = true
				cnt3++
			}

			if found2 == true && found3 == true {
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", cnt2*cnt3)

	// Part 2
	for _, boxID1 := range boxIds {
		cntDiff, diffPos := 0, 0

		for _, boxID2 := range boxIds {
			cntDiff = 0

			for pos := range boxID1 {
				if boxID1[pos] != boxID2[pos] {
					cntDiff++

					if cntDiff > 1 {
						break
					} else {
						diffPos = pos
					}
				}
			}

			if cntDiff == 1 {
				break
			}
		}

		if cntDiff == 1 {
			fmt.Printf("Part 2: %s\n", boxID1[:diffPos]+boxID1[diffPos+1:])
			break
		}
	}
}
