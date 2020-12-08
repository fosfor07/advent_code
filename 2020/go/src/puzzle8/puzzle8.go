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

	linesInit, err := readers.ReadLines(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	acc1, acc2, idx := 0, 0, 0
	lines := make([]string, len(linesInit))
	executed := make(map[int]bool)
	checked := make(map[int]bool)
	part2, end2 := false, false

	for {
		numCodes := copy(lines, linesInit)
		if numCodes != len(linesInit) {
			panic(fmt.Sprintf("Error: Copy failed! Copied elements: %d\n", numCodes))
		}

		// skip for part 1
		if part2 {
			for i := range lines {
				if _, ok := checked[i]; !ok {
					f := strings.Split(lines[i], " ")
					if f[0] == "jmp" {
						lines[i] = "nop " + f[1]
						checked[i] = true
						break
					} else if f[0] == "nop" {
						lines[i] = "jmp " + f[1]
						checked[i] = true
						break
					}
				}
			}
		}

		for {
			fields := strings.Split(lines[idx], " ")
			op := fields[0]
			arg, _ := strconv.Atoi(fields[1])

			if _, ok := executed[idx]; ok {
				if !part2 {
					acc1 = acc2
					part2 = true
				}
				acc2 = 0
				idx = 0
				executed = make(map[int]bool)
				break
			}
			executed[idx] = true

			if op == "acc" {
				acc2 += arg
				idx++
			} else if op == "jmp" {
				idx += arg
			} else if op == "nop" {
				idx++
			}

			if idx >= len(lines) {
				end2 = true
				break
			}
		}
		if end2 {
			break
		}
	}

	fmt.Printf("Part 1: %d\n", acc1)
	fmt.Printf("Part 2: %d\n", acc2)
}
