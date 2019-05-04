package main

import "fmt"

func main() {
	input := 15
	fmt.Println(fizzBuzz(input))
}

func fizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			result = append(result, "Fizz")
			continue
		}
		if i%5 == 0 {
			result = append(result, "Buzz")
			continue
		}

		result = append(result, fmt.Sprintf("%d", i))
	}

	return result
}
