package challenge66

import (
	"euler/shared"
	"fmt"
	"math/big"
)

const max = 1000

// Recursively evaluate a "continuous" fraction
func contFractToSimplFract(s []int) (num, denom *big.Int) {
	if len(s) == 1 {
		return big.NewInt(1), big.NewInt(int64(s[0]))
	}
	num, denom = contFractToSimplFract(s[1:])

	// The next line is equivalent to: return denom, s[0]*denom + num
	return denom, new(big.Int).Add(new(big.Int).Mul(big.NewInt(int64(s[0])), denom), num)
}

// Test whether the numbers provided are a valid solution
// Two criteria:
//   (x * x * n + 1) should be a square number and must be equal to y
func isDiophantineSolution(x, y *big.Int, n int) bool {
	// xSquared := new(big.Int).Mul(x, x)
	// xSquaredN := new(big.Int).Mul(xSquared, big.NewInt(int64(n)))
	// xSquaredNPlus1 := new(big.Int).Add(xSquaredN, big.NewInt(int64(1)))
	xSquaredNPlus1 := new(big.Int).Add(new(big.Int).Mul(new(big.Int).Mul(x, x), big.NewInt(int64(n))), big.NewInt(int64(1)))
	// sr will be an INTEGER square root of xSquaredNPlus1
	sr := new(big.Int).Sqrt(xSquaredNPlus1)
	// if xSquaredNPlus1 is a square number then the square of its integer square root will equal itself
	if (new(big.Int).Mul(sr, sr)).Cmp(xSquaredNPlus1) != 0 {
		return false
	}
	return sr.Cmp(y) == 0
}
func Challenge66() {
	var solution int
	maxDenom := new(big.Int)
	for count := 2; count <= max; count++ {
		sliceInt, _ := shared.Cfsr(count)
		// Cfsr returns a slice in the form [n, a, b,... x] where "n" is the base integer component and the remaining
		// digits are an infinitely repeating sequence of digits.  In our case we need to follow the sequence for an
		// unknown distance until a fraction is produced which is a valid Diophantine solution, so I don't know how
		// many iterations of the recurring sequence we need.  So just append a few iterations  sliceInt[0] is NOT
		// part of the recurring sequence.
		sliceInt = append(sliceInt, sliceInt[1:]...)
		sliceInt = append(sliceInt, sliceInt[1:]...)
		sliceInt = append(sliceInt, sliceInt[1:]...)
		// This is a bit inefficient but we need to start with testing a sub-slice with a length of 1
		// and then keep testing with increasing lengths until the returned values form a valid
		// Diophantine solution.
		for k := 1; k < len(sliceInt); k++ {
			num, denom := contFractToSimplFract(sliceInt[0:k])
			if isDiophantineSolution(num, denom, count) {
				if denom.Cmp(maxDenom) > 0 {
					maxDenom.Set(denom)
					solution = count
				}
				break
			}
		}
	}
	fmt.Printf("Challenge 66 solution is: %d (with an x value of %d)\n", solution, maxDenom)
}
