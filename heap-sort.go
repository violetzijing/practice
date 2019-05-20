package main

import "fmt"

func main() {
	input := []int{12, 11, 13, 5, 6, 7}
	fmt.Println(heapSort(input, len(input)))
}

func heapify(arr []int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		tmp := arr[i]
		arr[i] = arr[largest]
		arr[largest] = tmp
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int, n int) []int {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		tmp := arr[0]
		arr[0] = arr[i]
		arr[i] = tmp

		heapify(arr, i, 0)
	}

	return arr
}
