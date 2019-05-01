package main

import "fmt"

func main() {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 1},
	}
	fmt.Println(isToeplitzMatrix(input))
}

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix) == 1 {
		return true
	}
	n := len(matrix)
	m := len(matrix[0])
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}

	return true
}
