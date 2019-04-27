// https://leetcode.com/problems/k-closest-points-to-origin/submissions/
package main

import (
	"fmt"
	"sort"
)

func main() {

	test := [][]int{{1, 3}, {-2, 2}}
	fmt.Println(kClosest(test, 1))

	fmt.Println("Hello, playground")
}

func kClosest(points [][]int, K int) [][]int {
	var result = [][]int{}
	var distance = []int{}
	var distanceMap = map[int][][]int{}
	for _, coords := range points {
		
		count := coords[0]*coords[0] + coords[1]*coords[1]

		distance = append(distance, count)
		distanceMap[count] = append(distanceMap[count], coords)
	}
	sort.Ints(distance)
	for index, val := range distance {
		if index < K {
			result = append(result, distanceMap[val]...)
		}
	}

	return result
}

