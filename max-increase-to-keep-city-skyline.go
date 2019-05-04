package main

import "fmt"

func main() {
	input := [][]int{
		{59, 88, 44},
		{3, 18, 38},
		{21, 26, 51},
	}
	fmt.Println(maxIncreaseKeepingSkyline(input))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rowMax := make([]int, len(grid))
	columnMax := make([]int, len(grid[0]))
	sum1 := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if rowMax[i] < grid[i][j] {
				rowMax[i] = grid[i][j]
			}
			if columnMax[i] < grid[j][i] {
				columnMax[i] = grid[j][i]
			}
			sum1 += grid[i][j]
		}
	}

	sum2 := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if rowMax[j] < columnMax[i] {
				grid[i][j] = rowMax[j]
			} else {
				grid[i][j] = columnMax[i]
			}
			sum2 += grid[i][j]
		}
	}

	return sum2 - sum1
}
