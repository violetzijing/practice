package main

func main() {
	input := []int{1, 3, -1, -3, 5, 3, 6, 7}

}

func find(input []int, n int, k int) {
	i := 0
	var mins []int
	var maxs []int
	for i < n {
		j := 0
		var min int
		var max int
		for j < k {
			if min > input[i] {
				min = input[i]
			}
			if max < input[i] {
				max = input[i]
			}
			j += 1
		}
		mins = append(mins, min)
		maxs = append(maxs, max)

		i += 1
	}
}
