package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	lines, err := readers.ReadLines(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	part1, part2 := 0, 0

	for _, line := range lines {
		record := strings.Split(line, ":")
		gameId := strings.Split(record[0], " ")
		gameNum, _ := strconv.Atoi(gameId[1])
		draws := strings.Split(record[1], ";")

		red, green, blue := 0, 0, 0
		redMax, blueMax, greenMax := 0, 0, 0
		possible := true

		for _, draw := range draws {
			draw = strings.TrimSpace(draw)
			colors := strings.Split(draw, ",")

			for _, color := range colors {
				color = strings.TrimSpace(color)
				cubes := strings.Split(color, " ")

				cubeNum, _ := strconv.Atoi(cubes[0])
				if cubes[1] == "red" {
					red = cubeNum
					if red > redMax {
						redMax = red
					}
				} else if cubes[1] == "green" {
					green = cubeNum
					if green > greenMax {
						greenMax = green
					}
				} else if cubes[1] == "blue" {
					blue = cubeNum
					if blue > blueMax {
						blueMax = blue
					}
				}
			}

			if red > 12 || green > 13 || blue > 14 {
				possible = false
			}
		}

		if possible {
			part1 += gameNum
		}
		part2 += redMax * blueMax * greenMax
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
