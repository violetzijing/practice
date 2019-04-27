package main

import (
	"fmt"
)

func main() {
	input := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(input))
}

func sortArrayByParity(A []int) []int {
	if len(A) == 0 || len(A) == 1 {
		return A
	}
	var i = 0
	var j = len(A) - 1

	for i < j {
		iMod := A[i] & 1
		jMod := A[j] & 1
		if iMod != 0 && jMod == 0 {
			tmp := A[i]
			A[i] = A[j]
			A[j] = tmp
			i++
			j--
		}
		if iMod == 0 {
			i++
		}
		if jMod != 0 {
			j--
		}
	}

	return A
}
