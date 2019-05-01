package main

import "fmt"

func main() {
	input := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(islandPerimeter(input))
}

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	n := len(grid)
	m := len(grid[0])
	result := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				fmt.Println("before if", "i: ", i, "j: ", j)
				fmt.Println("===========")
				if i == 0 {
					fmt.Println("if i == 0")
					result++
				}
				if j == 0 {
					fmt.Println("if j == 0")
					result++
				}
				if i == n-1 {
					fmt.Println("if i == n-1")
					result++
				}
				if j == m-1 {
					fmt.Println("if j == m-1 ")
					result++
				}
				if i != 0 {
					if grid[i-1][j] == 0 {
						fmt.Println("if grid[i-1][j] == 0")
						result++
					}
				}
				if j != 0 {
					if grid[i][j-1] == 0 {
						fmt.Println("if grid[i][j-1] == 0")
						result++
					}
				}
				if i != n-1 {
					if grid[i+1][j] == 0 {
						fmt.Println("if grid[i+1][j] == 0")
						result++
					}
				}
				if j != m-1 {
					if grid[i][j+1] == 0 {
						fmt.Println("if grid[i][j+1] == 0")
						result++
					}
				}
			}
		}
	}

	return result
}
