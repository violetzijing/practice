package main

import "fmt"

func main() {
	input := "a-bC-dEf-ghIj"
	fmt.Println(reverseOnlyLetters(input))
}

func reverseOnlyLetters(S string) string {
	if len(S) == 0 {
		return S
	}

	str := []byte(S)
	i := 0
	j := len(str) - 1
	for i < j {
		for i < j && !((str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z')) {
			i += 1
		}
		for i < j && !((str[j] >= 'a' && str[j] <= 'z') || (str[j] >= 'A' && str[j] <= 'Z')) {
			j -= 1
		}
		tmp := str[i]
		str[i] = str[j]
		str[j] = tmp

		i += 1
		j -= 1
	}

	return string(str)
}
