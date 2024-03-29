package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

func bagColors(bag string, bags map[string]string, seen map[string]bool) int {
	numColors := 0

	for iBag, contain := range bags {
		matched, _ := regexp.MatchString(bag, contain)
		if matched {
			// Count each of bag colors only one time.
			if _, ok := seen[iBag]; !ok {
				numColors++
				seen[iBag] = true
				numColors += bagColors(iBag, bags, seen)
			}
		}
	}

	return numColors
}

func neededBags(bag string, bags map[string]string, outBags int) int {
	numBags := 0

	for iBag, contain := range bags {
		if iBag == bag {
			fields := strings.Split(contain, ",")
			for _, field := range fields {
				field = strings.Trim(field, " ")

				toks := strings.Split(field, " ")
				num, _ := strconv.Atoi(toks[0])
				numBags += num * outBags
				if num != 0 {
					outBags *= num
				}
				numBags += neededBags(toks[1]+" "+toks[2], bags, outBags)
				if num != 0 {
					outBags /= num
				}
			}
		}
	}

	return numBags
}

const myBag string = "shiny gold"

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
	bags := make(map[string]string)

	for _, line := range lines {
		fields := strings.Split(line, "contain")
		fields[0] = strings.Trim(fields[0], " ")
		fields[1] = strings.Trim(fields[1], " ")

		words := strings.Split(fields[0], " ")
		bags[words[0]+" "+words[1]] = fields[1]
	}

	seen := make(map[string]bool)
	part1 = bagColors(myBag, bags, seen)
	part2 = neededBags(myBag, bags, 1)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
