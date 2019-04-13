// https://leetcode.com/problems/unique-email-addresses

package main

import "fmt"

func main() {
	mails := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}

	fmt.Println(numUniqueEmails(mails))
}

func numUniqueEmails(emails []string) int {
	emailMap := map[string]bool{}
	for _, email := range emails {
		emailStr := []byte{}
		plusFlag := false
		index := 0
		for index = 0; index < len(email); index++ {
			if email[index] == '@' {
				break
			}
			if plusFlag {
				continue
			} else {
				if email[index] == '.' {
					continue
				}
				if email[index] == '+' {
					plusFlag = true
					continue
				}
				emailStr = append(emailStr, email[index])
			}
		}

		for i := index; i < len(email); i++ {
			emailStr = append(emailStr, email[i])
		}

		emailMap[string(emailStr)] = true
	}

	return len(emailMap)
}
