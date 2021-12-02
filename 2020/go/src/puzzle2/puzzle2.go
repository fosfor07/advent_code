package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

	valid1, valid2 := 0, 0

	r := regexp.MustCompile("([0-9]+)-([0-9]+)(.+):(.+)")

	for _, line := range lines {
		matches := r.FindStringSubmatch(line)

		minV, _ := strconv.Atoi(matches[1])
		maxV, _ := strconv.Atoi(matches[2])
		letter := strings.Trim(matches[3], " ")
		password := strings.Trim(matches[4], " ")

		chNum := 0
		fmin, fmax := false, false
		for i, ch := range password {
			// part 1
			if letter == string(ch) {
				chNum++
			}

			// part 2
			if (i == minV-1) && (letter == string(ch)) {
				fmin = true
			}
			if (i == maxV-1) && (letter == string(ch)) {
				fmax = true
			}
		}

		// part 1
		if (chNum >= minV) && (chNum <= maxV) {
			valid1++
		}

		// part 2
		if (fmin && !fmax) || (fmax && !fmin) {
			valid2++
		}
	}

	fmt.Printf("Part 1 = %d\n", valid1)
	fmt.Printf("Part 2 = %d\n", valid2)
}
