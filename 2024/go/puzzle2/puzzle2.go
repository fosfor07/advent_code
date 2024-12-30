package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
	"github.com/fosfor07/advent_code/pkg/utils"
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
		report := strings.Fields(line)
		if isReportSafe(report) {
			part1++
			part2++
			continue
		}

		for i := 0; i < len(report); i++ {
			newReport := make([]string, len(report))
			copy(newReport, report)

			newReport = append(newReport[:i], newReport[i+1:]...)
			if isReportSafe(newReport) {
				part2++
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func isReportSafe(report []string) bool {
	trend := 0

	for i := 0; i < len(report)-1; i++ {
		n1, _ := strconv.Atoi(string(report[i]))
		n2, _ := strconv.Atoi(string(report[i+1]))

		if n1 == n2 || utils.Abs(n1-n2) > 3 {
			return false
		}

		if trend == 0 {
			if n2 < n1 {
				trend = -1
			} else if n2 > n1 {
				trend = 1
			}
		} else if trend == -1 && n2 > n1 {
			return false
		} else if trend == 1 && n2 < n1 {
			return false
		}
	}

	return true
}
