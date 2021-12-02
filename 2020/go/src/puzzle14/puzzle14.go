package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"readers"
	"regexp"
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

	part1, part2 := 0, 0
	mem1 := make(map[int]int)
	mem2 := make(map[int]int)
	mask := ""
	numX, xMask := 0, 0.0

	r := regexp.MustCompile("([0-9]+)")

	for _, line := range lines {
		fields := strings.Split(line, "=")
		fields[0] = strings.Trim(fields[0], " ")
		fields[1] = strings.Trim(fields[1], " ")

		if fields[0] == "mask" {
			mask = fields[1]
			numX = strings.Count(mask, "X")
			xMask = math.Pow(2, float64(numX))
		} else {
			matches := r.FindStringSubmatch(fields[0])

			// Part 1
			addr, _ := strconv.Atoi(matches[1])
			v, _ := strconv.Atoi(fields[1])
			for i := 0; i < len(mask); i++ {
				if mask[i] != 'X' {
					if mask[i] == '1' {
						v |= 1 << (len(mask) - 1 - i)
					} else if mask[i] == '0' {
						v &^= 1 << (len(mask) - 1 - i)
					}
				}
			}
			mem1[addr] = v

			// Part 2
			addr, _ = strconv.Atoi(matches[1])
			v, _ = strconv.Atoi(fields[1])
			for j := 0; j < int(xMask); j++ {
				bitNum := int(xMask) / 2

				for i := 0; i < len(mask); i++ {
					if mask[i] == 'X' {
						if j&bitNum == bitNum {
							addr |= 1 << (len(mask) - 1 - i)
						} else if j&bitNum == 0 {
							addr &^= 1 << (len(mask) - 1 - i)
						}
						bitNum /= 2
					} else if mask[i] == '1' {
						addr |= 1 << (len(mask) - 1 - i)
					}
				}
				mem2[addr] = v
			}
		}
	}

	for _, v := range mem1 {
		part1 += v
	}
	for _, v := range mem2 {
		part2 += v
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
