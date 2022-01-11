package challenge73

import (
	"euler/shared"
	"fmt"
)

const max = 12_000

// Find the number reduced proper fractions exist which are > 1/3 and < 1/2, using numerator and denominators < max
func Challenge73() {
	var solution int
	for denom := 2; denom <= max; denom++ {
		for numer := 1; numer < denom; numer++ {
			if denom != 3*numer && denom != 2*numer {
				// If denom/numer is not exactly 2 or 3 then if integer division of denom/numer is equal to three
				// then numer/denom must be between 1/3 and 1/2.
				if denom/numer == 2 {
					// Now we just need to exclude fractions which are not "reduced" (e.g., 2/4 or 3/15), which
					// we can confirm to see if they have any prime factors in common.
					if shared.SizeOfIntersection(shared.GetPrimeFactors(numer), shared.GetPrimeFactors(denom)) != 0 {
						continue
					}
					solution++
				}
			}
		}
	}
	fmt.Printf("Challenge 73 solution is: %d\n", solution)
}
