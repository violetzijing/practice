package main

import "sort"

func main() {

}

func longestStrChain(words []string) int {
	if len(words) == 0 {
		return 0
	}
	max := 0
	SortStrings(words)
	strMap := map[string]int{}

	for _, word := range words {
		if _, ok := strMap[word]; ok {
			continue
		}
		strMap[word] = 1
		for i := 0; i < len(word); i++ {
			tmp := []byte(word)
			tmp1 = tmp[:i]
			tmp2 = tmp[i+1:]
			tmp3 := []byte{}
			tmp3 = append(tmp3, tmp1...)
			tmp3 = append(tmp3, tmp2...)
			after := string(tmp3)
			if _, ok := strMap[after]; ok {
				val1 := strMap[after]
				val2 := strMap[word]
				if val1+1 > val2 {
					strMap[word] = val1 + 1
				}
			}
		}
		if val := strMap[word]; val > max {
			max = val
		}
	}

	return max
}

type sortArr []string

func (s sortArr) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s sortArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortArr) Len() int {
	return len(s)
}

func SortStrings(s []string) {
	sort.Sort(sortArr(s))
}
