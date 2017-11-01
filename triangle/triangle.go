package triangle

const testVersion = 3

const (
	NaT = iota // not a triangle
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {

	if (a <= 0 || b <= 0 || c <= 0) || !(a+b >= c && b+c >= a && a+c >= b) {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else if a != b && b != c && a != c {
		return Sca
	}

	return NaT
}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

// Organize your code for readability.
