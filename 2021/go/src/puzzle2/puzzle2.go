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

	part1, part2 := 0, 0
	x, y1, aim, y2 := 0, 0, 0, 0

	for _, line := range lines {
		command := strings.Split(line, " ")
		v, _ := strconv.Atoi(command[1])

		if command[0] == "forward" {
			x += v
			y2 += aim * v
		} else if command[0] == "down" {
			y1 += v
			aim += v
		} else if command[0] == "up" {
			y1 -= v
			aim -= v
		}
	}
	part1 = x * y1
	part2 = x * y2

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
