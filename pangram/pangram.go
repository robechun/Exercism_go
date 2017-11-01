package pangram

import "strings"

const testVersion = 2

func IsPangram(s string) bool {

	alphaMap := map[string]int{}

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			alphaMap[strings.ToLower(string(s[i]))]++
		}
	}

	if len(alphaMap) < 26 {
		return false
	}

	return true

}
