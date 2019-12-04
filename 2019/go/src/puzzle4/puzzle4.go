package main

import (
	"fmt"
	"strconv"
)

func main() {
	const RangeStart int = 197487
	const RangeEnd int = 673251

	var prev byte     // previous character in string
	passCnt1 := 0     // number of valid passwords - part 1
	passCnt2 := 0     // number of valid passwords - part 2
	samePair := false // 2 adjacent digits are the same
	noDecr := true    // no decrease
	pairConf := false // pair has been confirmed
	cntSame := 0      // count the same digits

	// part 1
	for i := RangeStart; i <= RangeEnd; i++ {
		// reset variables
		samePair = false
		noDecr = true

		strCode := strconv.Itoa(i)

		// verify if password meets criteria
		for j := 1; j < len(strCode); j++ {
			prev = strCode[j-1]

			if strCode[j] == prev {
				samePair = true
			}

			if strCode[j] < prev {
				noDecr = false
				break
			}
		}

		if samePair && noDecr {
			passCnt1++
		}
	}

	// part 2
	for i := RangeStart; i <= RangeEnd; i++ {
		// reset variables
		samePair = false
		noDecr = true
		pairConf = false
		cntSame = 0

		strCode := strconv.Itoa(i)

		// verify if password meets criteria
		for j := 1; j < len(strCode); j++ {
			prev = strCode[j-1]

			if strCode[j] == prev && !pairConf {
				cntSame++
				if cntSame == 1 {
					samePair = true
				} else {
					samePair = false
				}
			} else {
				if cntSame == 1 {
					pairConf = true
				}
				cntSame = 0
			}

			if strCode[j] < prev {
				noDecr = false
				break
			}
		}

		if samePair && noDecr {
			passCnt2++
		}
	}

	fmt.Printf("PART 1 - Passwords meeting criteria: %d\n", passCnt1)
	fmt.Printf("PART 2 - Passwords meeting criteria: %d\n", passCnt2)
}
