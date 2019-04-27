package main

import "fmt"

func main() {
	input := [][]int{
		{1, 1, 0, 0},
		{1, 0, 0, 1},
		{0, 1, 1, 1},
		{1, 0, 1, 0},
	}

	fmt.Println(flipAndInvertImage(input))
}

func flipAndInvertImage(A [][]int) [][]int {
	numMap := map[int]int{
		0: 1,
		1: 0,
	}
	for _, item := range A {
		var i = 0
		var j = len(item) - 1
		for i < j {
			tmp := item[i]
			item[i] = item[j]
			item[j] = tmp
			i++
			j--
		}

		for n := 0; n < len(item); n++ {
			item[n] = numMap[item[n]]
		}
	}

	return A
}
