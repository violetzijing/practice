package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(input))
}

func sortedSquares(A []int) []int {
	result := make([]int, len(A))
	size := len(A)
	i := 0
	j := size - 1
	for p := size - 1; p >= 0; p-- {
		if math.Abs(float64(A[i])) > math.Abs(float64(A[j])) {
			result[p] = A[i] * A[i]
			i++
		} else {
			result[p] = A[j] * A[j]
			j--
		}
	}

	return result
}
