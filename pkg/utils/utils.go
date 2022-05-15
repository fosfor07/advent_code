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

// UniqueIntsSlice returns a unique subset of the provided int slice.
func UniqueIntsSlice(input []int) []int {
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

// UniqueStringsSlice returns a unique subset of the provided string slice.
func UniqueStringsSlice(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

// ReverseString returns reversed string.
func ReverseString(str string) string {
	result := ""
	for _, ch := range str {
		result = string(ch) + result
	}
	return result
}

// ReverseSlice returns reversed slice of integers.
func ReverseSlice(a []int) []int {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}

// Min returns the smallest number in slice.
func Min(values []int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min
}

// Max returns the largest number in slice.
func Max(values []int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}

	return max
}

// EqualIntSlices checks if 2 slices with int values are the same.
func EqualIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
