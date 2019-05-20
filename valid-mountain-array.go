package main

import "fmt"

func main() {
	input := []int{}
	fmt.Println(validMountainArray(input))
}

func validMountainArray(A []int) bool {
	if len(A) == 0 || len(A) == 1 {
		return false
	}

	max := A[0]
	maxIndex := 0
	for index, item := range A {
		if max < item {
			max = item
			maxIndex = index
		}
	}
	if max == A[0] || max == A[len(A)-1] {
		return false
	}

	for i := 0; i < maxIndex; i++ {
		if A[i] >= A[i+1] {
			return false
		}
	}
	for i := maxIndex; i < len(A)-1; i++ {
		if A[i] <= A[i+1] {
			return false
		}
	}

	return true
}

func validMountainArray2(A []int) bool {
	n := len(A)
	up := true
	i := 0
	if n < 3 {
		return false
	}
	if A[0] >= A[1] {
		return false
	}
	for i = 1; i < n-1; i++ {
		if A[i] == A[i+1] {
			return false
		}
		if up {
			if A[i+1] < A[i] {
				up = false
			}
		} else {
			if A[i+1] > A[i] {
				return false
			}
		}
	}

	return !up
}
