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

	res := make([]int, len(S))
	pos := 0 - len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			pos = i
		}
		tmp := 0
		if i > pos {
			tmp = i - pos
		} else {
			tmp = pos - i
		}
		if res[i] > tmp {
			res[i] = tmp
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if S[i] == C {
			pos = i
		}
		tmp := 0
		if i > pos {
			tmp = i - pos
		} else {
			tmp = pos - i
		}
		if res[i] > tmp {
			res[i] = tmp
		}
	}

	return res
}
