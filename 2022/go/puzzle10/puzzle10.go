package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

const width = 40
const height = 6

func nextCycle(cycle int, x int, signals []int, crt *[height][width]rune) (int, []int) {
	// part 1
	for i := 20; i <= 220; i += width {
		if cycle+1 == i {
			signals = append(signals, (cycle+1)*x)
		}
	}

	// part 2
	x = x + (cycle/width)*width
	if cycle == x-1 || cycle == x || cycle == x+1 {
		r := (cycle) / width
		c := (cycle) % width
		if r >= 0 && r < height && c >= 0 && c < width {
			crt[r][c] = '#'
		}
	}

	cycle++

	return cycle, signals
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

	cycle, x := 0, 1
	signals := []int{}

	var crt [height][width]rune
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			crt[r][c] = ' '
		}
	}

	for _, line := range lines {
		tkns := strings.Split(line, " ")
		if tkns[0] == "noop" {
			cycle, signals = nextCycle(cycle, x, signals, &crt)
		} else if tkns[0] == "addx" {
			cycle, signals = nextCycle(cycle, x, signals, &crt)
			cycle, signals = nextCycle(cycle, x, signals, &crt)

			v, _ := strconv.Atoi(tkns[1])
			x += v
		}
	}

	part1 := 0
	for _, signal := range signals {
		part1 += signal
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2:\n")
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			fmt.Printf("%c", crt[r][c])
		}
		fmt.Printf("\n")
	}
}
