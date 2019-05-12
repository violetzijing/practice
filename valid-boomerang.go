package main

import "fmt"

func main() {
	input := [][]int{
		{1, 1},
		{2, 2},
		{3, 2},
	}
	fmt.Println(isBoomerang(input))
}

func isBoomerang(p [][]int) bool {
	return (p[0][0]-p[1][0])*(p[0][1]-p[2][1]) != (p[0][0]-p[2][0])*(p[0][1]-p[1][1])

}
