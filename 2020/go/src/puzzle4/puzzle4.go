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

	valid1, valid2 := 0, 0
	tokN1, tokN2 := 0, 0

	for _, line := range lines {
		if len(line) > 0 {
			fields := strings.Split(line, " ")

			for _, field := range fields {
				tokens := strings.Split(field, ":")

				if tokens[0] == "byr" && len(tokens) > 1 {
					tokN1++

					// byr (Birth Year) - four digits; at least 1920 and at most 2002
					byr, _ := strconv.Atoi(tokens[1])
					if byr >= 1920 && byr <= 2002 {
						tokN2++
					}
				} else if tokens[0] == "iyr" && len(tokens) > 1 {
					tokN1++

					// iyr (Issue Year) - four digits; at least 2010 and at most 2020
					iyr, _ := strconv.Atoi(tokens[1])
					if iyr >= 2010 && iyr <= 2020 {
						tokN2++
					}
				} else if tokens[0] == "eyr" && len(tokens) > 1 {
					tokN1++

					// eyr (Expiration Year) - four digits; at least 2020 and at most 2030
					eyr, _ := strconv.Atoi(tokens[1])
					if eyr >= 2020 && eyr <= 2030 {
						tokN2++
					}
				} else if tokens[0] == "hgt" && len(tokens) > 1 {
					tokN1++

					// hgt (Height) - a number followed by either cm or in:
					// If cm, the number must be at least 150 and at most 193.
					// If in, the number must be at least 59 and at most 76.
					r, _ := regexp.Compile("(^[0-9]{2,3})([[:alpha:]]{2})$")
					matches := r.FindStringSubmatch(tokens[1])
					if len(matches) == 3 {
						hgtInt, _ := strconv.Atoi(matches[1])
						if (matches[2] == "cm" && (hgtInt >= 150 && hgtInt <= 193)) ||
							(matches[2] == "in" && (hgtInt >= 59 && hgtInt <= 76)) {
							tokN2++
						}
					}
				} else if tokens[0] == "hcl" && len(tokens) > 1 {
					tokN1++

					// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f
					matched, _ := regexp.MatchString("^#[[:xdigit:]]{6}$", tokens[1])
					if matched {
						tokN2++
					}
				} else if tokens[0] == "ecl" && len(tokens) > 1 {
					tokN1++

					// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth
					if tokens[1] == "amb" || tokens[1] == "blu" || tokens[1] == "brn" ||
						tokens[1] == "gry" || tokens[1] == "grn" || tokens[1] == "hzl" ||
						tokens[1] == "oth" {
						tokN2++
					}
				} else if tokens[0] == "pid" && len(tokens) > 1 {
					tokN1++

					// pid (Passport ID) - a nine-digit number, including leading zeroes
					matched, _ := regexp.MatchString("^[0-9]{9}$", tokens[1])
					if matched {
						tokN2++
					}
				} else if tokens[0] == "cid" && len(tokens) > 1 {
					// cid (Country ID) - ignored, missing or not
				}
			}
		} else {
			if tokN1 == 7 {
				valid1++
			}
			if tokN2 == 7 {
				valid2++
			}
			tokN1, tokN2 = 0, 0
		}
	}

	fmt.Printf("Part 1 = %d\n", valid1)
	fmt.Printf("Part 2 = %d\n", valid2)
}
