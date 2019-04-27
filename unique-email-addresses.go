package main

func numUniqueEmails(emails []string) int {
	emailMap := map[string]bool{}
	for _, email := range emails {
		emailStr := []rune{}
		atFlat := false
		plusFlag := false
		for index, str := range emails {
			if str == '@' {
				atFlat = true
			}
			if !atFlat && !plusFlag {
				if str == '.' {
					continue
				}
				if str == '+' {
					plusFlag = true
					continue
				}
				emailStr = append(emailStr, str)
			} else {
				emailStr = append(emailStr, str)
			}
		}
		if index == len(emails)-1 {
			emailMap[emailStr] = true
		}
	}

	sum := 0
	for _ := range emailMap {
		sum++
	}

	return sum
}
