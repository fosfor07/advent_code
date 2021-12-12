package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
	"unicode"
)

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func findPaths1(caves map[string][]string, cave string, path string, pathsNum int) int {
	path += "-" + cave + "-"

	if cave == "end" {
		pathsNum++
	} else {
		for _, c := range caves[cave] {
			if c != "start" {
				if IsLower(c) {
					cnt := strings.Count(path, "-"+c+"-")
					if cnt == 0 {
						pathsNum = findPaths1(caves, c, path, pathsNum)
					}
				} else {
					pathsNum = findPaths1(caves, c, path, pathsNum)
				}
			}
		}
	}

	return pathsNum
}

func findPaths2(caves map[string][]string, cave string, path string, pathsNum int, lower map[string]bool) int {
	path += "-" + cave + "-"

	if cave == "end" {
		pathsNum++
	} else {
		for _, c := range caves[cave] {
			if c != "start" {
				if IsLower(c) {
					limit := false
					for k := range lower {
						cnt := strings.Count(path, "-"+k+"-")
						if cnt > 1 {
							limit = true
							break
						}
					}

					cnt := strings.Count(path, "-"+c+"-")
					if cnt == 0 || (cnt == 1 && !limit) {
						pathsNum = findPaths2(caves, c, path, pathsNum, lower)
					}
				} else {
					pathsNum = findPaths2(caves, c, path, pathsNum, lower)
				}
			}
		}
	}

	return pathsNum
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

	caves := make(map[string][]string)
	lower := make(map[string]bool)

	for _, line := range lines {
		tokens := strings.Split(line, "-")
		caves[tokens[0]] = append(caves[tokens[0]], tokens[1])
		caves[tokens[1]] = append(caves[tokens[1]], tokens[0])
	}

	for k := range caves {
		if IsLower(k) && k != "start" && k != "end" {
			lower[k] = true
		}
	}

	part1 := findPaths1(caves, "start", "", 0)
	part2 := findPaths2(caves, "start", "", 0, lower)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
