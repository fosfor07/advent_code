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

type rangeType struct {
	min int
	max int
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

	part1 := 0
	your, nearby := false, false

	rules := make(map[string][]rangeType)

	var myticket []int
	numFlds := 0
	var validTickets [][]int
	numValTickets := 0

	for _, line := range lines {
		if line == "" {
			continue
		} else if line == "your ticket:" {
			your = true
			continue
		} else if line == "nearby tickets:" {
			nearby = true
			continue
		}

		if !your && !nearby {
			names := strings.Split(line, ":")
			name := names[0]

			ranges := strings.Split(line, " ")
			for _, rStr := range ranges {
				matched, _ := regexp.MatchString("^[0-9]+-[0-9]+$", rStr)
				if matched {
					limits := strings.Split(rStr, "-")
					var r rangeType
					r.min, _ = strconv.Atoi(limits[0])
					r.max, _ = strconv.Atoi(limits[1])
					rules[name] = append(rules[name], r)
				}
			}
		} else if your && !nearby {
			vals := strings.Split(line, ",")
			for _, vStr := range vals {
				v, _ := strconv.Atoi(vStr)
				myticket = append(myticket, v)
			}
			numFlds = len(myticket)
		} else if nearby {
			vals := strings.Split(line, ",")
			tValid := true
			var flds []int
			for _, vStr := range vals {
				fValid := false
				v, _ := strconv.Atoi(vStr)
				for _, rule := range rules {
					for _, rng := range rule {
						if v >= rng.min && v <= rng.max {
							fValid = true
							break
						}
					}
					if fValid {
						break
					}
				}
				if !fValid {
					tValid = false
					part1 += v
				}

				flds = append(flds, v)
			}
			if tValid {
				validTickets = append(validTickets, flds)
				numValTickets++
			}
		}
	}

	columnCnts := make(map[string][]int)

	for _, ticket := range validTickets {
		for i, v := range ticket {

			// For each rule count number of matching values in each column.
			for name, rule := range rules {
				for _, rng := range rule {
					if v >= rng.min && v <= rng.max {
						if _, ok := columnCnts[name]; !ok {
							columnCnts[name] = make([]int, numFlds)
							columnCnts[name][i] = 1
						} else {
							columnCnts[name][i]++
						}
						break
					}
				}
			}

		}
	}

	name2idx := make(map[string]int)
	skipIdx := make(map[int]bool)

	for f := 0; f < numFlds; f++ {
		// Find rule with exactly one matching column.
		for name, column := range columnCnts {
			idx := 0
			numMatchCols := 0
			for i, val := range column {
				if _, ok := skipIdx[i]; !ok {
					if val == numValTickets {
						numMatchCols++
						if numMatchCols > 1 {
							break
						} else {
							idx = i
						}
					}
				}
			}
			if numMatchCols == 1 {
				name2idx[name] = idx
				// Skip the idx column in the next iteration.
				skipIdx[idx] = true
				break
			}
		}
	}

	part2 := 1
	for name, idx := range name2idx {
		if strings.HasPrefix(name, "departure") {
			part2 *= myticket[idx]
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
