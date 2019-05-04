package main

import "fmt"

func main() {
	input := []int{5, 3, 2, 4, 1}
	fmt.Println(isMonotonic(input))
}

func isMonotonic(A []int) bool {
	if len(A) == 0 || len(A) == 1 {
		return true
	}

	max := A[0]
	min := A[0]
	maxFlag := true
	minFlag := true
	for _, item := range A {
		if max > item {
			maxFlag = false
		} else {
			max = item
		}
		if min < item {
			minFlag = false
		} else {
			min = item
		}
		if !maxFlag && !minFlag {
			return false
		}
	}

	return true
}
