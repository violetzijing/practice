package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4}
	query := [][]int{
		{1, 0},
		{-3, 1},
		{-4, 0},
		{2, 3},
	}
	fmt.Println(sumEvenAfterQueries(input, query))
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	result := []int{}
	sum := 0
	for _, a := range A {
		if a%2 == 0 {
			sum += a
		}
	}
	for i := 0; i < len(queries); i++ {
		index := queries[i][1]
		if A[index]%2 == 0 {
			sum -= A[index]
		}
		A[index] += queries[i][0]
		if A[index]%2 == 0 {
			sum += A[index]
		}
		result = append(result, sum)
	}
	return result
}
