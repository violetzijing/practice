// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
package main

import (
	"fmt"
)

func main() {
	test := []int{1,2,3,3}
	fmt.Println(repeatedNTimes(test))

	fmt.Println("Hello, playground")
}

func repeatedNTimes(A []int) int {
	aMap := map[int]int{}
	for _, val := range A {
		_, ok := aMap[val]
		if !ok {
			aMap[val] = 1
		} else {
			aMap[val] += 1
		}
		
		if aMap[val] >= len(A)/2 {
			return val
		}
	}

	return 0
}

