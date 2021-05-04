package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strconv"
	"strings"
	"utils"
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

	minDist, minSteps := 0, 0
	wire1 := make(map[string]int)

	for wireNum, line := range lines {
		steps, x, y := 0, 0, 0

		values, err := readers.ReadCsvWords(strings.NewReader(string(line)))
		if err != nil {
			panic(err)
		}

		for _, value := range values {
			dir, distStr := value[0], value[1:]
			dist, _ := strconv.Atoi(string(distStr))

			for i := 0; i < dist; i++ {
				if dir == 'R' {
					y++
				} else if dir == 'L' {
					y--
				} else if dir == 'U' {
					x--
				} else if dir == 'D' {
					x++
				}
				steps++

				key := fmt.Sprintf("x%dy%d", x, y)
				if wireNum == 0 {
					if _, ok := wire1[key]; !ok {
						wire1[key] = steps
					}
				} else {
					if _, ok := wire1[key]; ok {
						iDist := utils.Abs(x) + utils.Abs(y)

						if iDist < minDist || minDist == 0 {
							minDist = iDist
						}
						if (wire1[key]+steps) < minSteps || minSteps == 0 {
							minSteps = wire1[key] + steps
						}
					}
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", minDist)
	fmt.Printf("Part 2: %d\n", minSteps)
}
