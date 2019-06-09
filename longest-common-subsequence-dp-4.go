package main

import "fmt"

func main() {
	A := "AGGTAB"
	B := "GXTXAYB"
	fmt.Println(lcs(A, B))
}

func lcs(A, B string) int {
	m := len(A)
	n := len(B)
	d := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		d[i] = make([]int, n+1)
	}

	fmt.Println("before: ", d)
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				d[i][j] = 0
			} else if A[i-1] == B[j-1] {
				fmt.Println("previsous: ", d[i-1][j-1])
				d[i][j] = d[i-1][j-1] + 1
			} else {
				d[i][j] = max(d[i-1][j], d[i][j-1])
			}
		}
	}

	fmt.Println("after: ", d)
	return d[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
