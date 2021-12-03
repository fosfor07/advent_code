package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strconv"
	"strings"
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

	// part 1
	// power consumption = gamma rate * epsilon rate
	cntOnes := make([]int, len(lines[0]))
	cntZeros := make([]int, len(lines[0]))

	for _, line := range lines {
		for i, ch := range line {
			if ch == '0' {
				cntZeros[i]++
			} else if ch == '1' {
				cntOnes[i]++
			}
		}
	}

	gamma, epsilon := "", ""

	for i := 0; i < len(cntOnes); i++ {
		if cntOnes[i] > cntZeros[i] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	part1 := g * e

	// part 2
	// life support rating = oxygen generator rating * CO2 scrubber rating
	oxygen := make([]string, len(lines))
	copy(oxygen, lines)
	var oxygenTmp []string

	idx := 0
	for len(oxygen) > 1 {
		cnt0, cnt1 := 0, 0

		for _, record := range oxygen {
			if record[idx] == '0' {
				cnt0++
			} else if record[idx] == '1' {
				cnt1++
			}
		}

		for i := range oxygen {
			if cnt1 >= cnt0 {
				if oxygen[i][idx] == '1' {
					oxygenTmp = append(oxygenTmp, oxygen[i])
				}
			} else {
				if oxygen[i][idx] == '0' {
					oxygenTmp = append(oxygenTmp, oxygen[i])
				}
			}
		}

		oxygen = nil
		oxygen = make([]string, len(oxygenTmp))
		copy(oxygen, oxygenTmp)
		oxygenTmp = nil
		idx++
	}

	co2 := make([]string, len(lines))
	copy(co2, lines)
	var co2Tmp []string

	idx = 0
	for len(co2) > 1 {
		cnt0, cnt1 := 0, 0

		for _, record := range co2 {
			if record[idx] == '0' {
				cnt0++
			} else if record[idx] == '1' {
				cnt1++
			}
		}

		for i := range co2 {
			if cnt1 >= cnt0 {
				if co2[i][idx] == '0' {
					co2Tmp = append(co2Tmp, co2[i])
				}
			} else {
				if co2[i][idx] == '1' {
					co2Tmp = append(co2Tmp, co2[i])
				}
			}
		}

		co2 = nil
		co2 = make([]string, len(co2Tmp))
		copy(co2, co2Tmp)
		co2Tmp = nil
		idx++
	}

	oxRate, _ := strconv.ParseInt(oxygen[0], 2, 64)
	co2Rate, _ := strconv.ParseInt(co2[0], 2, 64)
	part2 := oxRate * co2Rate

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
