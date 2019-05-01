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
	grid2 := [][]int{}
}
