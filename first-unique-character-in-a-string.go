package main

import "fmt"

func main() {
	input := "loveleetcode"
	fmt.Println(firstUniqChar(input))
}

func firstUniqChar(s string) int {
	countMap := map[rune]int{}

	for _, item := range s {
		if _, ok := countMap[item]; !ok {
			countMap[item] = 1
			continue
		}
		countMap[item]++
	}
	for index, item := range s {
		if countMap[item] == 1 {
			return index
		}
	}

	return -1
}
