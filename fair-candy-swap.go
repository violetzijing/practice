package main

import "fmt"

func main() {
	A := []int{1, 1}
	B := []int{2, 2}
	fmt.Println(fairCandySwap(A, B))
}

func fairCandySwap(A []int, B []int) []int {
	sumA := 0
	sumB := 0
	mapA := map[int]bool{}
	for _, item := range A {
		sumA += item
		mapA[item] = true
	}
	for _, item := range B {
		sumB += item
	}

	diff := (sumA - sumB) / 2
	for i := 0; i < len(B); i++ {
		targetA := B[i] + diff
		if mapA[targetA] {
			return []int{targetA, B[i]}
		}
	}

	return []int{}
}
