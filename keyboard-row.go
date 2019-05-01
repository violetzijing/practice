package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{
		"asdfghjkl", "qwertyuiop", "zxcvbnm",
	}
	fmt.Println(findWords(input))
}

func findWords(words []string) []string {
	if len(words) == 0 {
		return words
	}

	firstRow := map[rune]int{
		'Q': 1,
		'W': 1,
		'E': 1,
		'R': 1,
		'T': 1,
		'Y': 1,
		'U': 1,
		'I': 1,
		'O': 1,
		'P': 1,
	}

	secondRow := map[rune]int{
		'A': 2,
		'S': 2,
		'D': 2,
		'F': 2,
		'G': 2,
		'H': 2,
		'J': 2,
		'K': 2,
		'L': 2,
	}

	thirdRow := map[rune]int{
		'Z': 3,
		'X': 3,
		'C': 3,
		'V': 3,
		'B': 3,
		'N': 3,
		'M': 3,
	}

	var result = []string{}

	for _, word := range words {
		if len(word) == 1 {
			result = append(result, word)
			continue
		}
		str := strings.ToUpper(word)
		equalFlag := true
		flag := 0
		if _, ok := firstRow[rune(str[0])]; ok {
			flag = 1
		}
		if _, ok := secondRow[rune(str[0])]; ok {
			flag = 2
		}
		if _, ok := thirdRow[rune(str[0])]; ok {
			flag = 3
		}

		for _, s := range str {
			switch flag {
			case 1:
				if _, ok := firstRow[s]; !ok {
					equalFlag = false
					break
				}
			case 2:
				if _, ok := secondRow[s]; !ok {
					equalFlag = false
					break
				}
			case 3:
				if _, ok := thirdRow[s]; !ok {
					equalFlag = false
					break
				}
			}
		}
		if equalFlag {
			result = append(result, word)
		}
	}

	return result
}
