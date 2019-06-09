package main

import (
	"fmt"
)

func main() {
	input := []int{2, 1, 1}
	fmt.Println(rearrangeBarcodes(input))
}

func rearrangeBarcodes(barcodes []int) []int {
	if len(barcodes) == 0 || len(barcodes) == 1 {
		return barcodes
	}

	n := len(barcodes)
	for i := 0; i < n/2; i += 2 {
		if barcodes[i] == barcodes[i+1] {
			j := i
			for j < n {
				if barcodes[j] != barcodes[i] {
					barcodes[j], barcodes[i+1] = barcodes[i+1], barcodes[j]
					break
				}
				j++
			}
		}
	}

	for i := 1; i < n/2; i += 2 {
		if barcodes[i] == barcodes[i+1] {
			j := i
			for j < n {
				if barcodes[j] != barcodes[i] {
					barcodes[j], barcodes[i+1] = barcodes[i+1], barcodes[j]
					break
				}
				j++
			}
		}
	}

	return barcodes
}
