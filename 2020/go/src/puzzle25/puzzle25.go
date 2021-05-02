package main

import "fmt"

func transform(value, subjectNumer int) int {
	return (value * subjectNumer) % 20201227
}

func main() {
	cardKey := 9789649
	doorKey := 3647239
	subjectNumer := 7
	cardLS, doorLS := 0, 0

	key := 1
	for key != cardKey {
		cardLS++
		key = transform(key, subjectNumer)
	}

	key = 1
	for key != doorKey {
		doorLS++
		key = transform(key, subjectNumer)
	}

	encKey1 := 1
	for i := 0; i < doorLS; i++ {
		encKey1 = transform(encKey1, cardKey)
	}

	encKey2 := 1
	for i := 0; i < cardLS; i++ {
		encKey2 = transform(encKey2, doorKey)
	}

	if encKey1 == encKey2 {
		fmt.Printf("Part 1: %d\n", encKey1)
	}
}
