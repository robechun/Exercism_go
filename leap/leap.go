// package checks if a year is a leap year
package leap

const testVersion = 3

// IsLeapYear calculates if given year is a leap year
func IsLeapYear(y int) bool {
	if y % 4 == 0 {
		if y % 100 != 0 {
			return true
		} else if y >= 400 && y % 400 != 0 { 
			return false
		}
		return true
	}

	return false
}
