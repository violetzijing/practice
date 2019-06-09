package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
	c := make([][]int, m)
	for i := 0; i < m; i++ {
		c[i] = make([]int, n)
	}
	c[0][0] = 1
	for i := 0; i < m; i++ {
		c[i][0] = 1
	}
	for i := 0; i < n; i++ {
		c[0][i] = 1
	}
	fmt.Println("c: ", c)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			c[i][j] = c[i-1][j] + c[i][j-1]
		}
	}
	fmt.Println("c: ", c)
	return c[m-1][n-1]
}
