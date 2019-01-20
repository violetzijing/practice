// https://leetcode.com/problems/squares-of-a-sorted-array/submissions/

package main

import (
	"fmt"
	"sort"
)

func main() {

	test := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(test))

	fmt.Println("Hello, playground")
}

func sortedSquares(A []int) []int {
	result := []int{}
	for _, a := range A {
		result = append(result, a*a)
	}
	sort.Ints(result)

	return result
}

