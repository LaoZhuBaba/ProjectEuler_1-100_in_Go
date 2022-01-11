package challenge72

import (
	"euler/shared"
	"fmt"
)

const max = 1_000_000

func Challenge72() {

	var solution int
	// The totient function tells us for any number as a denominator how many numerators below it in value
	// form reduced proper fractions.  For example, the totient of 4 is 2.  One way to confirm this is to
	// check that only 1/4, 3/4 are reduced proper fractions (I., 2/4 isn't.)  So all we actually need to do
	// is calculate the totent of every number from 2 to 1,000,000.

	for n := 2; n <= max; n++ {
		solution += shared.Totient(n)
	}
	fmt.Printf("Challenge 72 solution is: %d\n", solution)
}
