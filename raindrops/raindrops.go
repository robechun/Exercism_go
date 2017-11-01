// Pacakge raindrops outputs messages based on factors
package raindrops

import (
	"fmt"
)

const testVersion = 3

// Convert converts integer to message
func Convert(i int) string {

	res := ""
	converted := false

	if i%3 == 0 {
		res += "Pling"
		converted = true
	}
	if i%5 == 0 {
		res += "Plang"
		converted = true
	}
	if i%7 == 0 {
		res += "Plong"
		converted = true
	}

	if !converted {
		res = fmt.Sprintf("%v", i)
	}

	return res
}
