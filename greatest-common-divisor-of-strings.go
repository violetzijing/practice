package main

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
}

func gcdOfStrings(str1 string, str2 string) string {
	n := len(str1)
	m := len(str2)

	d := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		d[i] = make([]int, m+1)
	}

	maxlen := 0
	resultMap := map[string][]byte{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			key := fmt.Sprintf("%d-%d", i, j)
			resultMap[key] = []byte{}
			if str1[i-1] == str2[j-1] {
				d[i][j] = d[i-1][j-1] + 1
				resultMap[key] = append(resultMap[key], byte(str1[i-1]))
			} else {
				d[i][j] = 0
				resultMap[key] = []byte{}
			}
		}
	}

	maxI := 0
	maxJ := 0
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			if maxlen < d[i][j] {
				maxlen = d[i][j]
				maxI = i
				maxJ = j
			}
		}
	}
	maxKey := fmt.Sprintf("%d-%d", maxI, maxJ)

	return string(resultMap[maxKey])
}
