package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
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

	readRules := true
	rules := make(map[string][]string)

	for _, line := range lines {
		if line == "" {
			readRules = false
		} else if readRules {
			pages := strings.Split(line, `|`)
			rules[pages[0]] = append(rules[pages[0]], pages[1])
		} else {
			correctUpdate := true
			pages := strings.Split(line, `,`)

			correctUpdate = isUpdateCorrect(pages, rules)
			if !correctUpdate {
				seenRules := make(map[string]struct{})
				pages = applyRule(pages, rules, seenRules)
			}

			middlePage, _ := strconv.Atoi(pages[len(pages)/2])
			if correctUpdate {
				part1 += middlePage
			} else {
				part2 += middlePage
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func isUpdateCorrect(updatePages []string, rules map[string][]string) bool {
	for updatePageIdx, updatePage := range updatePages {
		rulePages, ok := rules[updatePage]
		if ok {
			for _, rulePage := range rulePages {
				orderIsCorrect := followsRule(updatePages, rulePage, updatePageIdx)
				if !orderIsCorrect {
					return false
				}
			}
		}
	}
	return true
}

func followsRule(updatePages []string, rulePage string, updatePageIdx int) bool {
	for idx, page := range updatePages {
		if page == rulePage && idx < updatePageIdx {
			return false
		}
	}
	return true
}

func applyRule(updatePages []string, rules map[string][]string, seenRules map[string]struct{}) []string {
	var elementsBefore []string
	var elementsAfter []string

	for referenceIdx, referencePage := range updatePages {
		rulePages, ruleExists := rules[referencePage]
		_, wasSeen := seenRules[referencePage]
		if ruleExists && !wasSeen {
			seenRules[referencePage] = struct{}{}
			elementsAfter = append(elementsAfter, referencePage)

			for adjustedIdx, adjustedPage := range updatePages {
				// skip reference page
				if adjustedPage != referencePage {
					ruleApplied := false

					// pages affected by the rule
					for _, rulePage := range rulePages {
						if adjustedPage == rulePage {
							elementsAfter = append(elementsAfter, adjustedPage)
							ruleApplied = true
							break
						}
					}

					// pages not affected by the rule
					if !ruleApplied {
						if adjustedIdx < referenceIdx {
							// Page is before reference page. Leave it there.
							elementsBefore = append(elementsBefore, adjustedPage)
						} else {
							// Page is after reference page. Leave it there.
							elementsAfter = append(elementsAfter, adjustedPage)
						}
					}
				}
			}
			break
		}
	}

	if len(elementsBefore) == 0 && len(elementsAfter) == 0 {
		return updatePages
	} else {
		return applyRule(slices.Concat(elementsBefore, elementsAfter), rules, seenRules)
	}
}
