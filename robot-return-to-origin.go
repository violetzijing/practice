package main

import "fmt"

func main() {
	input := "ULR"
	fmt.Println(judgeCircle(input))
}

func judgeCircle(moves string) bool {
	rowCount := 0
	columnCount := 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'U':
			columnCount++
		case 'D':
			columnCount--
		case 'L':
			rowCount++
		case 'R':
			rowCount--
		}
	}

	return (rowCount == 0) && (columnCount == 0)
}
