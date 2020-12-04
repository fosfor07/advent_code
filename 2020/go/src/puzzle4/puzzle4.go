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
	// Append an empty line to validate the last passport.
	lines = append(lines, "")

	valid1 := 0
	valid2 := 0
	tokN := 0
	tok1, tok2, tok3, tok4 := false, false, false, false
	tok5, tok6, tok7 := false, false, false

	for _, line := range lines {
		if len(line) > 0 {
			fields := strings.Split(line, " ")

			for _, field := range fields {
				tok := strings.Split(field, ":")

				if tok[0] == "byr" && len(tok) > 1 {
					tokN++

					// byr (Birth Year) - four digits; at least 1920 and at most 2002
					byr, _ := strconv.Atoi(tok[1])
					if byr >= 1920 && byr <= 2002 {
						tok1 = true
					}
				}
				if tok[0] == "iyr" && len(tok) > 1 {
					tokN++

					// iyr (Issue Year) - four digits; at least 2010 and at most 2020
					iyr, _ := strconv.Atoi(tok[1])
					if iyr >= 2010 && iyr <= 2020 {
						tok2 = true
					}
				}
				if tok[0] == "eyr" && len(tok) > 1 {
					tokN++

					// eyr (Expiration Year) - four digits; at least 2020 and at most 2030
					eyr, _ := strconv.Atoi(tok[1])
					if eyr >= 2020 && eyr <= 2030 {
						tok3 = true
					}
				}
				if tok[0] == "hgt" && len(tok) > 1 {
					tokN++

					// hgt (Height) - a number followed by either cm or in:
					// If cm, the number must be at least 150 and at most 193.
					// If in, the number must be at least 59 and at most 76.
					r, _ := regexp.Compile("(^[0-9]{2,3})([[:alpha:]]{2})$")
					matches := r.FindStringSubmatch(tok[1])
					if len(matches) == 3 {
						hgtInt, _ := strconv.Atoi(matches[1])
						if matches[2] == "cm" && (hgtInt >= 150 && hgtInt <= 193) {
							tok4 = true
						}
						if matches[2] == "in" && (hgtInt >= 59 && hgtInt <= 76) {
							tok4 = true
						}
					}
				}
				if tok[0] == "hcl" && len(tok) > 1 {
					tokN++

					// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f
					matched, _ := regexp.MatchString("^#[[:xdigit:]]{6}$", tok[1])
					if matched {
						tok5 = true
					}
				}
				if tok[0] == "ecl" && len(tok) > 1 {
					tokN++

					// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth
					if tok[1] == "amb" || tok[1] == "blu" || tok[1] == "brn" ||
						tok[1] == "gry" || tok[1] == "grn" || tok[1] == "hzl" ||
						tok[1] == "oth" {
						tok6 = true
					}
				}
				if tok[0] == "pid" && len(tok) > 1 {
					tokN++

					// pid (Passport ID) - a nine-digit number, including leading zeroes
					matched, _ := regexp.MatchString("^[0-9]{9}$", tok[1])
					if matched {
						tok7 = true
					}
				}
				if tok[0] == "cid" && len(tok) > 1 {
					// cid (Country ID) - ignored, missing or not
				}
			}
		} else {
			if tokN == 7 {
				valid1++
			}

			if tok1 && tok2 && tok3 && tok4 && tok5 && tok6 && tok7 {
				valid2++
			}

			tokN = 0
			tok1, tok2, tok3, tok4 = false, false, false, false
			tok5, tok6, tok7 = false, false, false
		}
	}

	fmt.Printf("Part 1 = %d\n", valid1)
	fmt.Printf("Part 2 = %d\n", valid2)
}
