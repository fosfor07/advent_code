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

// operators have the same precedence, and are evaluated left-to-right
func evaluate1(tkns []string, sIdx int, eIdx int) int {
	res, _ := strconv.Atoi(tkns[sIdx])
	for i := sIdx + 1; i < eIdx; i++ {
		if tkns[i] == "*" {
			n2, _ := strconv.Atoi(tkns[i+1])
			res *= n2
		} else if tkns[i] == "+" {
			n2, _ := strconv.Atoi(tkns[i+1])
			res += n2
		}
	}

	return res
}

// addition is evaluated before multiplication
func evaluate2(tkns []string, sIdx int, eIdx int) int {
	var nTkns []string

	for i := sIdx; i < eIdx; i++ {
		if tkns[i] == "+" {
			n1, _ := strconv.Atoi(nTkns[len(nTkns)-1])
			n2, _ := strconv.Atoi(tkns[i+1])
			nTkns[len(nTkns)-1] = strconv.Itoa(n1 + n2)
			i++
		} else {
			nTkns = append(nTkns, tkns[i])
		}
	}

	res, _ := strconv.Atoi(nTkns[0])
	for i := 1; i < len(nTkns); i++ {
		if nTkns[i] == "*" {
			n2, _ := strconv.Atoi(nTkns[i+1])
			res *= n2
		}
	}

	return res
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

	for _, line := range lines {
		var sTkns1, sTkns2 []string

		tkns := strings.Split(line, " ")
		for i := 0; i < len(tkns); i++ {
			tkns[i] = strings.TrimSpace(tkns[i])

			oIdx := strings.Index(tkns[i], "(")
			for oIdx >= 0 {
				sTkns1 = append(sTkns1, "(")
				tkns[i] = strings.Replace(tkns[i], "(", "", 1)
				oIdx = strings.Index(tkns[i], "(")
			}

			cIdx := strings.Index(tkns[i], ")")
			for cIdx >= 0 {
				if cIdx > 0 {
					r := regexp.MustCompile(`(^[0-9]+)(.*)$`)
					matches := r.FindStringSubmatch(tkns[i])
					sTkns1 = append(sTkns1, matches[1])
					tkns[i] = matches[2]
				}
				sTkns1 = append(sTkns1, ")")
				tkns[i] = strings.Replace(tkns[i], ")", "", 1)
				cIdx = strings.Index(tkns[i], ")")
			}

			if len(tkns[i]) > 0 {
				sTkns1 = append(sTkns1, tkns[i])
			}
		}

		// initialize tokens of part 2 by copying tokens of part 1
		sTkns2 = make([]string, len(sTkns1))
		copy(sTkns2, sTkns1)

		res1, res2 := 0, 0
		for {
			oIdx, cIdx := 0, 0
			brackets := false

			for i := 0; i < len(sTkns1); i++ {
				if sTkns1[i] == "(" {
					oIdx = i
				} else if sTkns1[i] == ")" {
					cIdx = i
					res1 = evaluate1(sTkns1, oIdx+1, cIdx)
					res2 = evaluate2(sTkns2, oIdx+1, cIdx)
					brackets = true
					break
				}
			}

			if brackets {
				uTkns1 := append(sTkns1[:oIdx], strconv.Itoa(res1))
				uTkns1 = append(uTkns1, sTkns1[cIdx+1:]...)
				sTkns1 = uTkns1

				uTkns2 := append(sTkns2[:oIdx], strconv.Itoa(res2))
				uTkns2 = append(uTkns2, sTkns2[cIdx+1:]...)
				sTkns2 = uTkns2
			} else {
				res1 = evaluate1(sTkns1, 0, len(sTkns1))
				res2 = evaluate2(sTkns2, 0, len(sTkns2))
				break
			}
		}

		part1 += res1
		part2 += res2
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
