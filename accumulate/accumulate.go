// Package accumulate operates on a list and returns the
// 	operated list
package accumulate

const testVersion = 1

// Accumulate operates on a list and returns the
// 	operated list
func Accumulate(c []string, op func(string) string) []string {

	res := c[:0]
	for _, elem := range c {
		res = append(res, op(elem))
	}
	
	return res
}
