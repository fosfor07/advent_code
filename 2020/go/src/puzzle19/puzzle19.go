package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
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
	var messages []string
	rules := make(map[string][]string)
	dRules := make(map[string][]string)

	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}

		if line[0] == 'a' || line[0] == 'b' {
			messages = append(messages, line)
		} else {
			fields := strings.Split(line, ":")
			fields[0] = strings.TrimSpace(fields[0])
			fields[1] = strings.TrimSpace(fields[1])

			if fields[1][0] == '"' {
				dRules[fields[0]] = append(dRules[fields[0]], string(fields[1][1]))
			} else {
				subRules := strings.Split(fields[1], "|")
				for _, sr := range subRules {
					sr = strings.TrimSpace(sr)
					rules[fields[0]] = append(rules[fields[0]], sr)
				}
			}
		}
	}

	// part 1
	for len(rules) > 0 {
		var toDelete []string

		for key, rule := range rules {
			ready := true
			for _, sr := range rule {
				fields := strings.Split(sr, " ")
				for _, f := range fields {
					if _, ok := dRules[f]; !ok {
						ready = false
						break
					}
				}
				if !ready {
					break
				}
			}

			if ready {
				for _, subrule := range rule {
					var tmp1 []string
					var tmp2 []string
					fields := strings.Split(subrule, " ")
					for _, f := range fields {
						for rID, dRule := range dRules {
							if string(f) == rID {
								if len(tmp1) == 0 {
									for _, dSubrule := range dRule {
										tmp1 = append(tmp1, dSubrule)
									}
								} else {
									for i := 0; i < len(dRule); i++ {
										for j := 0; j < len(tmp1); j++ {
											tmp2 = append(tmp2, tmp1[j]+dRule[i])
										}
									}
								}
								break
							}
						}
					}

					if len(tmp2) > 0 {
						for _, t := range tmp2 {
							dRules[key] = append(dRules[key], t)
						}
					} else {
						for _, t := range tmp1 {
							dRules[key] = append(dRules[key], t)
						}
					}
				}

				toDelete = append(toDelete, key)
			}
		}

		for _, ruleID := range toDelete {
			delete(rules, ruleID)
		}
	}

	for _, str := range messages {
		for _, rule := range dRules["0"] {
			if str == rule {
				part1++
			}
		}
	}

	// part 2
	len42 := len(dRules["42"][0])
	len31 := len(dRules["31"][0])

	for _, str := range messages {
		done := false
		leftFor8 := str
		// apply rule number 8
		for {
			matched8 := false
			cmp42for8 := leftFor8[0:len42]
			leftFor8 = leftFor8[len42:]

			for _, rule := range dRules["42"] {
				if cmp42for8 == rule {
					matched8 = true
					leftFor11 := leftFor8
					// apply rule number 11
					for {
						matched11 := false
						for _, r42 := range dRules["42"] {
							cmp42for11 := leftFor11[0:len42]
							if cmp42for11 == r42 {
								for _, r31 := range dRules["31"] {
									cmp31for11 := leftFor11[len(leftFor11)-len31:]
									if cmp31for11 == r31 {
										leftFor11 = leftFor11[len42:]
										if len(leftFor11) == len(dRules["31"][0]) {
											part2++
											done = true
										} else {
											leftFor11 = leftFor11[0 : len(leftFor11)-len31]
										}
										matched11 = true
										break
									}
								}
								if matched11 {
									break
								}
							}
						}

						if done || !matched11 {
							break
						}
					}
				}
			}

			if done || !matched8 || len(leftFor8) < 3*len42 {
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
