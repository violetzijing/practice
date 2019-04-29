package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Let's take LeetCode contest"
	fmt.Println(reverseWords(input))
}

func reverseWords(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	strs := strings.Split(s, " ")
	result := []byte{}

	for _, str := range strs {
		tmp := []byte{}
		for i := len(str) - 1; i >= 0; i-- {
			tmp = append(tmp, str[i])
		}
		result = append(result, tmp...)
		result = append(result, ' ')
	}

	return string(result[:len(result)-1])
}
