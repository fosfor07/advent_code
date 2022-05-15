package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

const input string = "yzbqklnj"

func main() {
	part1, part2, cnt := 0, 0, 1

	for {
		text := fmt.Sprintf("%s%d", input, cnt)
		hash := md5.Sum([]byte(text))
		hashHexStr := fmt.Sprintf("%x", hash)

		if part1 == 0 {
			matched1, _ := regexp.MatchString("^0{5}", hashHexStr)
			if matched1 {
				part1 = cnt
			}
		}

		if part2 == 0 {
			matched2, _ := regexp.MatchString("^0{6}", hashHexStr)
			if matched2 {
				part2 = cnt
			}
		}

		if part1 != 0 && part2 != 0 {
			break
		}
		cnt++
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
