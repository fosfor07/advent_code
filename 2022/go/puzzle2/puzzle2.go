package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

const (
	Lose int = 0
	Draw     = 3
	Win      = 6
)

const (
	Rock     int = 1
	Paper        = 2
	Scissors     = 3
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
		plays := strings.Split(line, " ")
		if plays[0] == "A" {
			if plays[1] == "X" {
				part1 += Draw + Rock
			} else if plays[1] == "Y" {
				part1 += Win + Paper
			} else if plays[1] == "Z" {
				part1 += Lose + Scissors
			}
		} else if plays[0] == "B" {
			if plays[1] == "X" {
				part1 += Lose + Rock
			} else if plays[1] == "Y" {
				part1 += Draw + Paper
			} else if plays[1] == "Z" {
				part1 += Win + Scissors
			}
		} else if plays[0] == "C" {
			if plays[1] == "X" {
				part1 += Win + Rock
			} else if plays[1] == "Y" {
				part1 += Lose + Paper
			} else if plays[1] == "Z" {
				part1 += Draw + Scissors
			}
		}
	}

	// part 2
	for _, line := range lines {
		plays := strings.Split(line, " ")
		if plays[0] == "A" {
			if plays[1] == "X" {
				part2 += Lose + Scissors
			} else if plays[1] == "Y" {
				part2 += Draw + Rock
			} else if plays[1] == "Z" {
				part2 += Win + Paper
			}
		} else if plays[0] == "B" {
			if plays[1] == "X" {
				part2 += Lose + Rock
			} else if plays[1] == "Y" {
				part2 += Draw + Paper
			} else if plays[1] == "Z" {
				part2 += Win + Scissors
			}
		} else if plays[0] == "C" {
			if plays[1] == "X" {
				part2 += Lose + Paper
			} else if plays[1] == "Y" {
				part2 += Draw + Scissors
			} else if plays[1] == "Z" {
				part2 += Win + Rock
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
