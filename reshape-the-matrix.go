package main

import "fmt"

func main() {
	input := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(matrixReshape(input, 1, 4))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	numlist := []int{}
	for _, rows := range nums {
		for _, cols := range rows {
			numlist = append(numlist, cols)
		}
	}
	if len(numlist) != r*c {
		return nums
	}

	result := make([][]int, r)
	var index = 0
	for j := 0; j < r; j++ {
		for i := 0; i < c; i++ {
			result[j] = append(result[j], numlist[index])
			index++
		}
	}

	return result
}
