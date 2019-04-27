package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := []string{
		"900 google.mail.com",
		"50 yahoo.com",
		"1 intel.mail.com",
		"5 wiki.org",
	}
	fmt.Println(subdomainVisits(input))
}

func subdomainVisits(cpdomains []string) []string {
	result := map[string]int{}

	for _, domain := range cpdomains {
		result1 := strings.Split(domain, " ")
		count, _ := strconv.Atoi(string(result1[0]))
		result[result1[1]] += count

		subdomain := result1[1]
		for i, str := range subdomain {
			if str == '.' {
				result[subdomain[i+1:]] = count + result[subdomain[i+1:]]
			}
		}
	}

	strResult := []string{}
	for key, value := range result {
		strResult = append(strResult, fmt.Sprintf("%d %s", value, key))
	}

	return strResult
}
