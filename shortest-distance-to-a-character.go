package main

import (
	"fmt"
)

func main() {
	input := "loveleetcod"
	found := byte('e')
	fmt.Println("length of string: ", len(input))
	fmt.Println(shortestToChar(input, found))
}

func shortestToChar(S string, C byte) []int {
	if len(S) == 0 {
		return []int{}
	}

	first := 0
	second := 0
	result := make([]int, len(S))
	for first != len(S) {
		if S[first] != C {
			first++
			continue
		}
		if S[second] != C && S[first] == C {
			for i := second; i < first; i++ {
				result[i] = first - i
			}
			second = first
			first++
			continue
		}

		if S[second] == C && S[first] == C {
			for i := second; i < first; i++ {
				tmp1 := first - i
				tmp2 := i - second
				if tmp1 < tmp2 {
					result[i] = tmp1
				} else {
					result[i] = tmp2
				}
			}
			second = first
			first++

			continue
		}
	}

	if second != first {
		for i := second; i < first; i++ {
			result[i] = i - second
		}
	}

	return result
}
