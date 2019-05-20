package main

import (
	"fmt"
	"strings"
)

func main() {
	A := "abcde"
	B := "abced"
	fmt.Println(rotateString(A, B))
}

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	testStr := fmt.Sprintf("%s%s", A, A)

	return strings.Contains(testStr, B)
}
