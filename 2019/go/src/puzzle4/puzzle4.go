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
	pairFnd1 := false // group of 2 or more the same digits
	pairFnd2 := false // group of only 2 the same digits
	noDecr := true    // no decrease
	pairConf := false // pair has been confirmed
	cntSame := 0      // count the same digits

	for i := RangeStart; i <= RangeEnd; i++ {
		pairFnd1 = false // part 1
		pairFnd2 = false // part 2
		noDecr = true
		pairConf = false
		cntSame = 0

		strCode := strconv.Itoa(i)

		// verify if password meets criteria
		for j := 1; j < len(strCode); j++ {
			prev = strCode[j-1]

			if strCode[j] == prev && !pairConf {
				pairFnd1 = true
				cntSame++
				if cntSame == 1 {
					pairFnd2 = true
				} else {
					pairFnd2 = false
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

		if pairFnd1 && noDecr {
			passCnt1++
		}

		if pairFnd2 && noDecr {
			passCnt2++
		}
	}

	fmt.Printf("Part 1: %d\n", passCnt1)
	fmt.Printf("Part 2: %d\n", passCnt2)
}
