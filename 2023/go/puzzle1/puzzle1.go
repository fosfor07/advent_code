package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/fosfor07/advent_code/pkg/readers"
)

var Digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

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
		for part := 1; part <= 2; part++ {
			firstDigit := getFirstDigit(line, part)
			lastDigit := getLastDigit(line, part)
			firstStr := strconv.Itoa(firstDigit)
			lastStr := strconv.Itoa(lastDigit)
			result, _ := strconv.Atoi(firstStr + lastStr)

			if part == 1 {
				part1 += result
			} else if part == 2 {
				part2 += result
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func getFirstDigit(line string, part int) int {
	threeChars, fourChars, fiveChars := "", "", ""

	for i, ch := range line {
		if unicode.IsDigit(ch) {
			digit, _ := strconv.Atoi(string(ch))
			return digit
		}

		if part == 2 {
			if i > 1 {
				threeChars = fmt.Sprintf("%c%c%c", line[i-2], line[i-1], line[i])
			}
			if i > 2 {
				fourChars = fmt.Sprintf("%c%c%c%c", line[i-3], line[i-2], line[i-1], line[i])
			}
			if i > 3 {
				fiveChars = fmt.Sprintf("%c%c%c%c%c", line[i-4], line[i-3], line[i-2], line[i-1], line[i])
			}

			if _, ok := Digits[threeChars]; ok {
				return Digits[threeChars]
			}
			if _, ok := Digits[fourChars]; ok {
				return Digits[fourChars]
			}
			if _, ok := Digits[fiveChars]; ok {
				return Digits[fiveChars]
			}
		}
	}
	return 0
}

func getLastDigit(line string, part int) int {
	digit := 0
	threeChars, fourChars, fiveChars := "", "", ""

	for i, ch := range line {
		if unicode.IsDigit(ch) {
			digit, _ = strconv.Atoi(string(ch))
		}

		if part == 2 {
			if i > 1 {
				threeChars = fmt.Sprintf("%c%c%c", line[i-2], line[i-1], line[i])
			}
			if i > 2 {
				fourChars = fmt.Sprintf("%c%c%c%c", line[i-3], line[i-2], line[i-1], line[i])
			}
			if i > 3 {
				fiveChars = fmt.Sprintf("%c%c%c%c%c", line[i-4], line[i-3], line[i-2], line[i-1], line[i])
			}

			if _, ok := Digits[threeChars]; ok {
				digit = Digits[threeChars]
			}
			if _, ok := Digits[fourChars]; ok {
				digit = Digits[fourChars]
			}
			if _, ok := Digits[fiveChars]; ok {
				digit = Digits[fiveChars]
			}
		}
	}
	return digit
}
