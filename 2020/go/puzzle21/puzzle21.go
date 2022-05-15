package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
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

	alergenList := make(map[string][]string)
	ingList := make(map[string]int)

	for _, line := range lines {
		fields := strings.Split(line, "(contains")
		fields[0] = strings.TrimSpace(fields[0])
		fields[1] = strings.TrimSpace(fields[1])

		ingredients := strings.Split(fields[0], " ")
		for _, ing := range ingredients {
			ing = strings.TrimSpace(ing)
			if _, ok := ingList[ing]; !ok {
				ingList[ing] = 1
			} else {
				ingList[ing]++
			}
		}

		alergens := strings.Split(fields[1], ",")
		for _, alergen := range alergens {
			alergen = strings.TrimSpace(alergen)
			alergen = strings.Trim(alergen, ")")
			alergenList[alergen] = append(alergenList[alergen], fields[0])
		}
	}

	alg2ing := make(map[string]string)

	for len(alergenList) > 0 {
		for alg, ingStrings := range alergenList {
			ingCnt := make(map[string]int)
			maxCnt, maxOcc, ingName := 0, 0, ""

			for _, ingString := range ingStrings {
				ings := strings.Split(ingString, " ")
				for _, ing := range ings {
					if _, ok := ingCnt[ing]; !ok {
						ingCnt[ing] = 1
					} else {
						ingCnt[ing]++
					}
				}
			}

			for ing, cnt := range ingCnt {
				if cnt > maxCnt {
					maxCnt = cnt
					maxOcc = 1
					ingName = ing
				} else if cnt == maxCnt {
					maxOcc++
				}
			}

			if maxCnt == len(ingStrings) && maxOcc == 1 {
				alg2ing[alg] = ingName
				delete(alergenList, alg)
				delete(ingList, ingName)

				for _, ingStrings := range alergenList {
					for idx := range ingStrings {
						ingStrings[idx] = strings.ReplaceAll(ingStrings[idx], ingName, "")
						ingStrings[idx] = strings.TrimSpace(ingStrings[idx])
						ingStrings[idx] = strings.Join(strings.Fields(ingStrings[idx]), " ")
					}
				}
				break
			}
		}
	}

	part1 := 0
	for _, cnt := range ingList {
		part1 += cnt
	}

	keys := make([]string, 0, len(alg2ing))
	for k := range alg2ing {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	part2 := ""
	for _, k := range keys {
		part2 += alg2ing[k] + ","
	}
	part2 = strings.Trim(part2, ",")

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %s\n", part2)
}
