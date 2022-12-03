package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"

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

	part1, part2 := 0, 0

	// part 1
	for _, line := range lines {
		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]

		fnd := false
		for _, item1 := range compartment1 {
			for _, item2 := range compartment2 {
				if item1 == item2 {
					priority := int(item1)

					if unicode.IsLower(item1) {
						priority = priority - int('a') + 1
					} else {
						priority = priority - int('A') + 27
					}

					part1 += priority
					fnd = true
					break
				}
			}
			if fnd {
				break
			}
		}
	}

	// part 2
	cnt := 0
	group := []string{}
	for _, line := range lines {
		group = append(group, line)
		cnt++

		fnd := false
		if cnt == 3 {
			for _, item1 := range group[0] {
				for _, item2 := range group[1] {
					for _, item3 := range group[2] {
						if item1 == item2 && item1 == item3 {
							priority := int(item1)

							if unicode.IsLower(item1) {
								priority = priority - int('a') + 1
							} else {
								priority = priority - int('A') + 27
							}
							part2 += priority
							fnd = true
							break
						}
					}
					if fnd {
						break
					}
				}
				if fnd {
					break
				}
			}

			group = []string{}
			cnt = 0
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
