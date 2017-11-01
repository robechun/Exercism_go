// Package clock should handle times without dates
package clock

import "fmt"

const testVersion = 4

// Clock type has two fields
// hour and minutes
type Clock struct {
	hour   int
	minute int
}

// New creates new clock with given hour/minute
func New(hour, minute int) Clock {

	m := minute
	h := hour
	for m >= 60 {
		m -= 60
		h++
	}

	for m < 0 {
		m += 60
		h--
	}

	for h >= 24 {
		h -= 24
	}
	for h < 0 {
		h += 24
	}

	return Clock{h, m}
}

// Outputs the time in string format
func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds time (or subtracts) to the clock
func (c Clock) Add(minutes int) Clock {

	c.minute += minutes

	for c.minute >= 60 {
		c.minute -= 60
		c.hour++
	}

	for c.minute < 0 {
		c.minute += 60
		c.hour--
	}

	for c.hour >= 24 {
		c.hour -= 24
	}

	for c.hour < 0 {
		c.hour += 24
	}

	return c
}
