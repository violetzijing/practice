package main

import "fmt"

func main() {
	input := [][]int{
		{1},
	}
	fmt.Println(uniquePathsWithObstacles(input))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)
	c := make([][]int, m)
	for i := 0; i < m; i++ {
		c[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 1 {
		c[0][0] = 0
	} else {

		c[0][0] = 1
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			c[i][0] = 0
		} else {
			c[i][0] = c[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			c[0][j] = 0
		} else {
			c[0][j] = c[0][j-1]
		}
	}

	fmt.Println("before: ", c)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				c[i][j] = c[i-1][j] + c[i][j-1]
			} else {
				fmt.Println("i: ", i)
				fmt.Println("j: ", j)
				c[i][j] = 0
			}
		}
	}
	fmt.Println("after: ", c)

	return c[m-1][n-1]
}
