package utils

// Abs returns absolute value of number.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// SumDigits returns sum of digits in a number.
func SumDigits(num int) int {
	if num == 0 {
		return 0
	}
	num = Abs(num)
	return num%10 + SumDigits(num/10)
}

// UniqueInts returns a unique subset of the provided int slice.
func UniqueInts(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

// ReverseSlice returns reversed slice of integers.
func ReverseSlice(a []int) []int {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
