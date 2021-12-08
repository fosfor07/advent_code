package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"sort"
	"strconv"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
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

	part1, part2 := 0, 0
	nums := make(map[int]string)
	var g235 []string
	var g039 []string
	var g069 []string

	for _, line := range lines {
		parts := strings.Split(line, "|")
		uniqueStr := strings.TrimSpace(parts[0])
		digitsStr := strings.TrimSpace(parts[1])

		unique := strings.Fields(uniqueStr)
		digits := strings.Fields(digitsStr)

		for _, d := range digits {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				part1++
			}
		}

		for _, u := range unique {
			if len(u) == 2 {
				nums[1] = u
			} else if len(u) == 3 {
				nums[7] = u
			} else if len(u) == 4 {
				nums[4] = u
			} else if len(u) == 7 {
				nums[8] = u
			} else if len(u) == 5 {
				g235 = append(g235, u)
			} else if len(u) == 6 {
				g069 = append(g069, u)
			}
		}

		// Create group with 0, 3 and 9 to find segments of 3 and 6.
		for _, u := range unique {
			if len(u) == 5 || len(u) == 6 {
				cnt := 0
				for _, ch1 := range u {
					for _, ch2 := range nums[1] {
						if ch1 == ch2 {
							cnt++
						}
					}
					if cnt == 2 {
						g039 = append(g039, u)
						break
					}
				}
			}

			// Find segments of 9 using segments of 4.
			if len(u) == 6 {
				cnt := 0
				for _, ch1 := range u {
					for _, ch2 := range nums[4] {
						if ch1 == ch2 {
							cnt++
						}
					}
				}
				if cnt == 4 {
					nums[9] = u
				}
			}
		}

		// Find segments of 3 by comparing groups with 0,3,9 and 0,6,9.
		for _, n1 := range g039 {
			fnd := false
			for _, n2 := range g069 {
				if n1 == n2 {
					fnd = true
				}
			}
			if !fnd {
				nums[3] = n1
				break
			}
		}

		// Find segments of 6 by comparing groups with 0,6,9 and 0,3,9.
		for _, n1 := range g069 {
			fnd := false
			for _, n2 := range g039 {
				if n1 == n2 {
					fnd = true
				}
			}
			if !fnd {
				nums[6] = n1
				break
			}
		}

		// Find segments of 0. We can do it because we know segments of 3 and 9.
		for _, n1 := range g039 {
			if n1 != nums[3] && n1 != nums[9] {
				nums[0] = n1
				break
			}
		}

		// All segments of 5 are also in 6. We will use it to distinguish 2 and 5.
		for _, u := range unique {
			if len(u) == 5 {
				if u != nums[3] {
					cnt := 0
					for _, ch1 := range u {
						for _, ch2 := range nums[6] {
							if ch1 == ch2 {
								cnt++
							}
						}
					}
					if cnt == 5 {
						nums[5] = u
						break
					}
				}
			}
		}

		// Only segments of 2 are still unknown. Find them by inspecting group with 2,3 and 5.
		for _, n1 := range g235 {
			if n1 != nums[3] && n1 != nums[5] {
				nums[2] = n1
			}
		}

		for k := range nums {
			nums[k] = SortString(nums[k])
		}

		var outStr string
		for _, d := range digits {
			d = SortString(d)

			for k := range nums {
				if d == nums[k] {
					outStr += strconv.Itoa(k)
				}
			}
		}
		outInt, _ := strconv.Atoi(outStr)
		part2 += outInt

		g235 = nil
		g039 = nil
		g069 = nil

		for k := range nums {
			delete(nums, k)
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
