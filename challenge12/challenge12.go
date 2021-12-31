// What is the value of the first triangle number to have over five hundred divisors?

package challenge12

import (
	"fmt"
	"math"
)

func countFactors(n int) int {
	//factors := make([]int, 0)
	// factors := []int{1}
	factorCount := 1
	var i int
	if n == 1 {
		return 1
	}
	sr := int(math.Sqrt(float64(n)))
	// Every factor below the square root of n has a matching factor above the square root of n
	// so we only need to count up to square root of n and then double the number of factors.
	// The one proviso is that square numbers have a single extra factor which is its square root.
	for i = 2; i <= sr; i++ {
		if n%i == 0 {
			// increment once...
			factorCount++
			if i != sr {
				// and increment again so long as i is not the square root of n
				factorCount++
			}
		}
	}
	factorCount++
	// factors = append(factors, n)
	//fmt.Printf("%d has factors: %v\n", n, factors)

	// return len(factors)
	return factorCount

}

func Challenge12() {

	for tri, count := 1, 1; ; tri += count {
		factorCount := countFactors(tri)
		// fmt.Printf("%d has %d factors\n", tri, factorCount)
		count++
		if factorCount >= 500 {
			fmt.Printf("Challenge 12 solution: triangular number is %d (with %d factors)\n", tri, factorCount)
			break
		}
	}
}
