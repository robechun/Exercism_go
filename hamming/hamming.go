// Package hamming calculates the hamming distance
// 	between two strings
package hamming

import "errors"

const testVersion = 6

// Distance calulates the distance between the two
// 	given strings
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("Different length strings")
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
