// Two=fer either prints One for you, one for me or replaces you with a person's name
package twofer

// Prints One for you, one for me if s is empty or One for s, one for me if s isnt empty
func ShareWith(s string) string {
	if s == "" {
		s = "you"
	}
	return "One for " + s + ", one for me."
}
