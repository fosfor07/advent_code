package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

type Operator string

const (
	OperatorAdd         Operator = "+"
	OperatorMultiply    Operator = "*"
	OperatorConcatenate Operator = "||"
)

var operatorsPart1 = []Operator{OperatorAdd, OperatorMultiply}
var operatorsPart2 = []Operator{OperatorAdd, OperatorMultiply, OperatorConcatenate}

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
		lineParts := strings.Split(line, ":")
		result, _ := strconv.Atoi(lineParts[0])
		numbers := strings.Fields(lineParts[1])

		equations := createEquations(numbers, operatorsPart1)
		if verifyEquations(equations, result) {
			part1 += result
		}

		equations = createEquations(numbers, operatorsPart2)
		if verifyEquations(equations, result) {
			part2 += result
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func createEquations(numbers []string, operators []Operator) []string {
	var equations []string

	for {
		if len(numbers) == 1 {
			for i := range equations {
				equations[i] += " " + numbers[0]
			}
			return equations
		} else {
			n := numbers[0]
			numbers = numbers[1:]

			var updatedEquations []string

			for _, op := range operators {
				if len(equations) == 0 {
					equation := n + " " + string(op)
					updatedEquations = append(updatedEquations, equation)
				} else {
					for _, equation := range equations {
						equation += " " + n + " " + string(op)
						updatedEquations = append(updatedEquations, equation)
					}
				}
			}
			equations = updatedEquations
		}
	}
}

func verifyEquations(equations []string, result int) bool {
	for _, equation := range equations {
		if calculate(equation) == result {
			return true
		}
	}
	return false
}

func calculate(equation string) int {
	result := 0

	fields := strings.Fields(equation)
	if len(fields) == 1 {
		result, _ = strconv.Atoi(fields[0])
		return result
	}

	for i := 0; i < len(fields)-1; i++ {
		if i == 0 {
			result, _ = strconv.Atoi(fields[i])
		} else {
			_, err := strconv.Atoi(fields[i])
			if err != nil {
				n, _ := strconv.Atoi(fields[i+1])
				if fields[i] == string(OperatorAdd) {
					result += n
				} else if fields[i] == string(OperatorMultiply) {
					result *= n
				} else if fields[i] == string(OperatorConcatenate) {
					result, _ = strconv.Atoi(strconv.Itoa(result) + fields[i+1])
				}
			}
		}
	}
	return result
}
