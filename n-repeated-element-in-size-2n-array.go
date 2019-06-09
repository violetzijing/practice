package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 3}
	fmt.Println(repeatedNTimes(input))
}

func repeatedNTimes(A []int) int {
	for i := 3; i < len(A); i++ {
		if A[i-1] == A[i] || A[i-2] == A[i] {
			return A[i]
		}
	}

	return 0
}
