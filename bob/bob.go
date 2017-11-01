package bob

import (
	"strings"
)

const testVersion = 3

func Hey(s string) string {

	// precedence -->
	// if all uppercase letters than you reply with
	// Whoa Chill out

	// if ends with ? reply with
	// Sure

	// if only whitespace
	// Fine, be that way!

	// Everything else
	// Whatever.

	uppercase := false
	lowercase := false
	nonwhite := false
	m := strings.Trim(s, " \t\n\r")
	mlen := len(m)

	if mlen == 0 {
		return "Fine. Be that way!"
	}

	for i := 0; i < len(m); i++ {
		if m[i] >= 'a' && m[i] <= 'z' {
			lowercase = true
		}
		if m[i] >= 'A' && m[i] <= 'Z' {
			uppercase = true
		}
		if m[i] != ' ' || m[i] != '\t' || m[i] != '\n' || m[i] != '\r' {
			nonwhite = true
		}

	}

	if uppercase && !lowercase {
		return "Whoa, chill out!"
	} else if m[mlen-1] == '?' {
		return "Sure."
	} else if !nonwhite {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
