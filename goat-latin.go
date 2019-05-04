package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "I speak Goat Latin"
	fmt.Println(toGoatLatin(input))
}

func toGoatLatin(S string) string {
	if len(S) == 0 {
		return ""
	}
	firstMap := map[byte]string{
		'a': "ma",
		'A': "ma",
		'e': "ma",
		'E': "ma",
		'i': "ma",
		'I': "ma",
		'o': "ma",
		'O': "ma",
		'u': "ma",
		'U': "ma",
	}

	strs := strings.Split(S, " ")
	result := []string{}
	for index, str := range strs {
		fmt.Println("str: ", str)
		if val, ok := firstMap[str[0]]; ok {
			str = fmt.Sprintf("%s%s", str, val)
			fmt.Println("first if: ", str)
		} else {
			firstLetter := str[0]
			fmt.Println("firstLetter: ", string(firstLetter))
			str = str[1:]
			str = fmt.Sprintf("%s%s%s", str, string(firstLetter), "ma")
		}
		for i := 0; i <= index; i++ {
			str = fmt.Sprintf("%s%s", str, "a")
		}
		fmt.Println("str after operation: ", str)
		result = append(result, str)
	}

	return strings.Join(result, " ")
}
