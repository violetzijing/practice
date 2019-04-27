package main

import "fmt"

func main() {
	input := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(input))
}

func sortArrayByParityII(A []int) []int {
	if len(A) == 0 || len(A) == 1 {
		return A
	}

	even := 0
	odd := 1

	for odd < len(A) && even < len(A) {
		for odd < len(A) && A[odd]%2 == 1 {
			odd += 2
		}
		for even < len(A) && A[even]%2 == 0 {
			even += 2
		}
		if odd < len(A) && even < len(A) {
			tmp := A[odd]
			A[odd] = A[even]
			A[even] = tmp
		}
	}

	return A
}
