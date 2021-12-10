package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"sort"
	"strings"
)

type incompleteType struct {
	remaining string
	missing   string
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

	wrongChars := make(map[rune]int)
	var expectedChars []rune
	var incompleteLines []string

	for _, line := range lines {
		isCorrupted := false

		for _, ch := range line {
			if ch == '(' {
				expectedChars = append(expectedChars, ')')
			} else if ch == '[' {
				expectedChars = append(expectedChars, ']')
			} else if ch == '{' {
				expectedChars = append(expectedChars, '}')
			} else if ch == '<' {
				expectedChars = append(expectedChars, '>')
			} else {
				if ch != expectedChars[len(expectedChars)-1] {
					wrongChars[ch]++
					isCorrupted = true
					break
				} else {
					if len(expectedChars) > 0 {
						expectedChars = expectedChars[:len(expectedChars)-1]
					}
				}
			}
		}

		if !isCorrupted {
			incompleteLines = append(incompleteLines, line)
		}
	}

	for ch, numChars := range wrongChars {
		points := 0
		if ch == ')' {
			points = 3
		} else if ch == ']' {
			points = 57
		} else if ch == '}' {
			points = 1197
		} else if ch == '>' {
			points = 25137
		}
		part1 += numChars * points
	}

	var opened []rune
	var incmpList []incompleteType
	for _, line := range incompleteLines {
		for _, ch := range line {
			if ch == '(' || ch == '[' || ch == '{' || ch == '<' {
				opened = append(opened, ch)
			} else {
				if len(opened) > 0 {
					opened = opened[:len(opened)-1]
				}
			}
		}

		newElem := incompleteType{remaining: string(opened), missing: ""}
		incmpList = append(incmpList, newElem)
		opened = nil
	}

	for i := range incmpList {
		for j := len(incmpList[i].remaining) - 1; j >= 0; j-- {
			if incmpList[i].remaining[j] == '(' {
				incmpList[i].missing += ")"
			} else if incmpList[i].remaining[j] == '[' {
				incmpList[i].missing += "]"
			} else if incmpList[i].remaining[j] == '{' {
				incmpList[i].missing += "}"
			} else if incmpList[i].remaining[j] == '<' {
				incmpList[i].missing += ">"
			}
		}
	}

	var scores []int
	for _, incmp := range incmpList {
		score := 0
		for _, c := range incmp.missing {
			score *= 5
			if c == ')' {
				score += 1
			} else if c == ']' {
				score += 2
			} else if c == '}' {
				score += 3
			} else if c == '>' {
				score += 4
			}
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	part2 = scores[len(scores)/2]

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
