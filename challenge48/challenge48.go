// The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

// Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.

package challenge48

import (
	"fmt"
	"math"
)

const c48Length = 10
const c48Max = 1000

// Add two integers but truncate to the n least significant digits
func addTrunc(x, y, n uint64) uint64 {
	return (x + y) % uint64((math.Pow(10, float64(n))))
}

// Multiply two integers but truncate to the n least significant digits
func mulTrunc(x, y, n uint64) uint64 {
	return (x * y) % uint64((math.Pow(10, float64(n))))
}

// Raise x to the y power, but truncate to the n least significant digits
func powTrunc(x, y, n uint64) uint64 {
	tot := x
	for y > 1 {
		tot = mulTrunc(tot, x, n)
		y--
	}
	return tot
}

func Challenge48() {

	n := uint64(0)
	for count := uint64(1); count <= c48Max; count++ {
		// fmt.Printf("count is: %d : n is %d\n", count, n)
		// fmt.Printf("%d to the poweer of %d is: %d\n", count, count, powTrunc(count, count, 10))
		// n = addTrunc(n, powTrunc(count, count, c48Length), c48Length)
		n = addTrunc(n, powTrunc(count, count, c48Length), c48Length)
		// fmt.Printf("count is: %d : n is %d\n", count, n)

	}
	fmt.Printf("Challenge 48 solution: %d\n", n)
}
