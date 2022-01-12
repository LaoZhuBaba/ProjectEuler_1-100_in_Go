package challenge75

import (
	"euler/shared"
	"fmt"
)

const max = 1_500_000

// const max = 49

func Challenge75() {
	var solution int

	var triples [max + 1]int
outer:
	for n1 := 2; ; n1++ {
		for n2 := 1; n2 < n1; n2++ {
			// First step is to get all possible pairs of numbers which are "co-prime".  We establish this
			// by checking that they have no prime factors in common.
			if shared.SizeOfIntersection(shared.GetPrimeFactors(n1), shared.GetPrimeFactors(n2)) == 0 {
				// Now that we have a co-prime pair of n1 & n2 we can use Euclid's formula to generate
				// Pythagorean triples...
				x := n1*n1 - n2*n2
				y := 2 * n2 * n1
				z := n2*n2 + n1*n1
				totalLength := x + y + z
				// Euclid's theorem only finds "primitive" (simplest form) triples if x & y and not both odd.
				if x%2 != 1 && y%2 != 1 {
					continue
				}
				// See explanation below for why we break on max*2
				if totalLength > max*2 {
					break outer
				}
				if totalLength > max {
					// At this point we can't break because the triples are not necessarily generated in order of their total sum.
					// E.g., (742995 12068 743093) is discovered before (744765 3452 744773).  This means that we can't stop
					// just because we have found one triple which is too large.  However, if we continue until we have found
					// a triple which exceeds double the max, we can be certain that no more triples will be found which are
					// less than max.  Hence the "break outer" above
					continue
				}
				// For the purposes of this challenge we need to account not just for the "primitive" triple but also its
				// multiples.  E.g., after finding (3 4 5) we need to mark off not just 12 (3+4+5) but also 24, 36, 48, etc.
				for multipleLengths := totalLength; multipleLengths <= max; multipleLengths += totalLength {
					triples[multipleLengths]++
				}
				fmt.Printf("Triple is: %d %d %d (%d)\n", x, y, z, x+y+z)
			}
		}
	}
	// The challenge is to find all instances where a total only has one triple associated with it.  E.g., 12 has exactly (3+4+5)
	// and no other examples.
	for _, v := range triples {
		if v == 1 {
			solution++
		}
	}
	fmt.Printf("Challenge75 solution is: %d\n", solution)
}
