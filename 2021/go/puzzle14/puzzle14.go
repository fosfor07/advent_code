package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

const stepsPart1 int = 10
const stepsPart2 int = 40

type rulesType struct {
	pair, newElem string
}

func subtractMinFromMax(pairs map[string]int, template string) int {
	elements := make(map[byte]int)

	for pair, num := range pairs {
		elements[pair[0]] += num
	}
	elements[template[len(template)-1]] += 1

	minNum, maxNum := 0, 0
	for _, num := range elements {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum || minNum == 0 {
			minNum = num
		}
	}

	return maxNum - minNum
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

	template := ""
	pairs := make(map[string]int)
	var rules []rulesType

	for n, line := range lines {
		if line == "" {
			continue
		} else if n == 0 {
			for _, ch := range line {
				template += string(ch)
			}
		} else {
			tokens := strings.Split(line, "->")
			tokens[0] = strings.TrimSpace(tokens[0])
			tokens[1] = strings.TrimSpace(tokens[1])

			r := rulesType{pair: tokens[0], newElem: string(tokens[1][0])}
			rules = append(rules, r)
		}
	}

	for i := 0; i < len(template)-1; i++ {
		pairs[string(template[i])+string(template[i+1])]++
	}

	for s := 0; s < stepsPart2; s++ {
		pairs2 := make(map[string]int)
		for pair, num := range pairs {
			pairs2[pair] = num
		}

		for pair, num := range pairs {
			if num > 0 {
				for _, r := range rules {
					if pair == r.pair {
						key1 := string(pair[0]) + r.newElem
						key2 := r.newElem + string(pair[1])
						pairs2[key1] += pairs[pair]
						pairs2[key2] += pairs[pair]
						pairs2[pair] -= pairs[pair]
					}
				}
			}
		}

		for pair, num := range pairs2 {
			pairs[pair] = num
		}

		if s == stepsPart1-1 {
			part1 := subtractMinFromMax(pairs, template)
			fmt.Printf("Part 1: %d\n", part1)
		}
	}

	part2 := subtractMinFromMax(pairs, template)
	fmt.Printf("Part 2: %d\n", part2)
}
