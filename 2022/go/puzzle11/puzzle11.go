package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

type monkeyType struct {
	items               []int
	op                  string
	opValue, testValue  int
	trueDest, falseDest int
	inspections         int
}

func monkeyBusiness(monkeys []monkeyType, roundLimit int, part int) {
	manageLvl := 1
	for _, m := range monkeys {
		manageLvl = manageLvl * m.testValue
	}

	round := 0
	for round < roundLimit {
		for mId, m := range monkeys {
			for _, item := range m.items {
				monkeys[mId].inspections++

				wLevel := item
				if m.op == "*" {
					if m.opValue != 0 {
						wLevel = wLevel * m.opValue
					} else {
						wLevel = wLevel * wLevel
					}
				} else if m.op == "+" {
					if m.opValue != 0 {
						wLevel = wLevel + m.opValue
					} else {
						wLevel = wLevel + wLevel
					}
				}

				if part == 1 {
					wLevel = wLevel / 3
				} else if part == 2 {
					if wLevel >= manageLvl {
						wLevel = wLevel % manageLvl
					}
				}

				if wLevel%m.testValue == 0 {
					monkeys[m.trueDest].items = append(monkeys[m.trueDest].items, wLevel)
				} else {
					monkeys[m.falseDest].items = append(monkeys[m.falseDest].items, wLevel)
				}
			}
			monkeys[mId].items = nil
		}

		round++
	}
}

func monkeyLevel(monkeys []monkeyType) int {
	sorted := []int{}
	for _, m := range monkeys {
		sorted = append(sorted, m.inspections)
	}
	sort.Ints(sorted)

	return sorted[len(sorted)-1] * sorted[len(sorted)-2]
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

	monkeys1 := []monkeyType{}
	monkey := monkeyType{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		tkns := strings.Split(line, " ")

		if tkns[0] == "Monkey" && tkns[1] != "0:" {
			monkeys1 = append(monkeys1, monkey)
			monkey = monkeyType{}
		} else if tkns[1] == "items:" {
			for i := 2; i < len(tkns); i++ {
				itemStr := strings.Trim(tkns[i], ",")
				item, _ := strconv.Atoi(itemStr)
				monkey.items = append(monkey.items, item)
			}
		} else if tkns[0] == "Operation:" {
			monkey.op = tkns[4]
			monkey.opValue, _ = strconv.Atoi(tkns[5])
		} else if tkns[0] == "Test:" {
			monkey.testValue, _ = strconv.Atoi(tkns[3])
		} else if tkns[0] == "If" && tkns[1] == "true:" {
			monkey.trueDest, _ = strconv.Atoi(tkns[5])
		} else if tkns[0] == "If" && tkns[1] == "false:" {
			monkey.falseDest, _ = strconv.Atoi(tkns[5])
		}
	}
	monkeys1 = append(monkeys1, monkey)
	monkeys2 := make([]monkeyType, len(monkeys1))
	copy(monkeys2, monkeys1)

	// part 1
	monkeyBusiness(monkeys1, 20, 1)
	part1 := monkeyLevel(monkeys1)

	// part 2
	monkeyBusiness(monkeys2, 10000, 2)
	part2 := monkeyLevel(monkeys2)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
