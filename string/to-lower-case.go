// https://leetcode.com/problems/to-lower-case

package main

import (
	"fmt"
)

func main() {
	teststr := "Hello"
	fmt.Println(toLowerCase(teststr))

	fmt.Println("Hello, playground")
}

func toLowerCase(str string) string {
	generated := []rune{}
	for _, r := range str {
		if r >= 65 && r <= 90 {
			r += 32
		}
		generated = append(generated, r)

	}

	return string(generated)
}

