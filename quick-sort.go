package main

import "fmt"

func main() {
	input := []int{10, 80, 30, 90, 40, 50, 70}
	fmt.Println(quickSort(input, 0, len(input)-1))
}

func quickSort(arr []int, low int, high int) []int {
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}

	return arr
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			i++
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
		}
	}
	tmp := arr[i+1]
	arr[i+1] = arr[high]
	arr[high] = tmp

	return i + 1
}
