package main

import (
	"fmt"
	"strings"
)

func main() {
	A := "apple apple apple"
	B := "banana banana"
	fmt.Println(uncommonFromSentences(A, B))
}

func uncommonFromSentences(A string, B string) []string {
	if len(B) == 0 {
		return []string{}
	}
	bSlice := strings.Split(B, " ")
	if len(A) == 0 {
		return bSlice
	}

	result := []string{}
	mapA := map[string]int{}
	for _, str := range strings.Split(A, " ") {
		if _, ok := mapA[str]; !ok {
			mapA[str] = 1
		} else {
			mapA[str]++
		}
	}

	for _, str := range bSlice {
		if _, ok := mapA[str]; !ok {
			mapA[str] = 1
		} else {
			mapA[str]++
		}
	}
	for key, val := range mapA {
		if val == 1 {
			result = append(result, key)
		}
	}

	return result
}
