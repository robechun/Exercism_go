package diffsquares

import "math"

const testVersion = 1

func SumOfSquares(n int) int {

	sum := 0.0
	for i := 0; i <= n; i++ {
		sum += math.Pow(float64(i), 2.0)
	}

	return int(sum)

}

func SquareOfSums(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}

	return int(math.Pow(float64(sum), 2.0))

}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
