package main

import "fmt"

func main() {
	name := "alex"
	typed := "aaleex"
	fmt.Println(isLongPressedName(name, typed))
}

func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}

	i := 0
	for j := 0; j < len(typed); j++ {
		if i < len(name) && name[i] == typed[j] {
			i++
		} else if j == 0 || typed[j] != typed[j-1] {
			return false
		}
	}

	return i == len(name)
}
