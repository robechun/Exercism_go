// Package acronym gets acronym from given string
package acronym

import "strings"

const testVersion = 3

// Abbreviate gets given string and makes it into its acryonym.
func Abbreviate(s string) string {
	res := strings.ToUpper(string(s[0]))

	for i := 1; i < len(s); i++ {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') {
			for (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') {
				i++
			}

			res += strings.ToUpper(string(s[i]))
		}

	}

	return res

}
