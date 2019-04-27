package main

// leetcode.com/problems/jewels-and-stones

func numJewelsInStones(J string, S string) int {
	sMap := map[rune]int{}
	jMap := map[rune]int{}
	count := 0

	for _, s := range S {
		sMap[s] = sMap[s] + 1

	}

	for _, j := range J {
		jMap[j] = jMap[j] + 1
	}

	for key, _ := range jMap {
		if sCount, ok := sMap[key]; ok {
			count += sCount
		}
	}

	return count
}
