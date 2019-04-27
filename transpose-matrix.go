package main

import "fmt"

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(transpose(input))
}

func transpose(A [][]int) [][]int {
	result := make([][]int, len(A[0]))

	for j := 0; j < len(A[0]); j++ {
		result[j] = make([]int, len(A))
		for i := 0; i < len(A); i++ {
			result[j][i] = A[i][j]
		}
	}

	return result
}
