package matrix

import "fmt"

// DisplayRune prints the matrix with runes.
func DisplayRune(matrix [][]rune) {
	M := len(matrix)
	N := len(matrix[0])
	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			fmt.Printf("%c", matrix[r][c])
		}
		fmt.Printf("\n")
	}
}

// RotateCwRune rotates the matrix with runes clockwise.
func RotateCwRune(matrix [][]rune) {
	N := len(matrix)
	x := N - 1
	y := N / 2

	for r := 0; r < y; r++ {
		for c := r; c < x-r; c++ {
			tmp := matrix[r][c]
			matrix[r][c] = matrix[x-c][r]
			matrix[x-c][r] = matrix[x-r][x-c]
			matrix[x-r][x-c] = matrix[c][x-r]
			matrix[c][x-r] = tmp
		}
	}
}

// FlipVRune flips the matrix with runes vertically.
func FlipVRune(matrix [][]rune) {
	N := len(matrix)

	for c := 0; c < N; c++ {
		for r := 0; r < N/2; r++ {
			tmp := matrix[r][c]
			matrix[r][c] = matrix[N-1-r][c]
			matrix[N-1-r][c] = tmp
		}
	}
}

// FlipHRune flips the matrix with runes horizontally.
func FlipHRune(matrix [][]rune) {
	N := len(matrix)

	for r := 0; r < N; r++ {
		for c := 0; c < N/2; c++ {
			tmp := matrix[r][c]
			matrix[r][c] = matrix[r][N-1-c]
			matrix[r][N-1-c] = tmp
		}
	}
}
