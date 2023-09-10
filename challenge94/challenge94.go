package challenge94

import (
	"euler/shared"
	"fmt"
	"math"
)

type Triple [3]int

func getTriples2(max int) (ret []Triple) {
	var hypotenuse, side1, side2 int
	var side1Square, side2Square, hypotenuseSquare int
	var triple Triple
	for side1 = 3; side1 <= max; side1++ {
		if side1%2 == 0 {
			hypotenuse = (2 * side1) + 1
		} else {
			hypotenuse = (2 * side1) - 1
		}
		side1Square = side1 * side1
		hypotenuseSquare = hypotenuse * hypotenuse
		side2Square = hypotenuseSquare - side1Square
		side2 = int(math.Sqrt(float64(side2Square)))
		if side2*side2 == int(side2Square) {
			// Found a triple
			triple[0] = side1
			triple[1] = side2
			triple[2] = hypotenuse
			ret = append(ret, triple)
		}
	}
	return ret
}

// getTriples is another possible solution based on findind Pythagorean triples.
// This is just a normal Pythagorean triple calculate except I've added a filter
// so that only triples that can create nearly equilateral trianges are included.
func getTriples(t *[]Triple, max int) {
	var delta int
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
				if x*2 != z+1 && y*2 != z-1 {
					continue
				}
				// Add the following condition to ignore triples which are not "primative"
				//
				// Euclid's theorem only finds "primitive" (simplest form) triples if x & y and not both odd.
				if x%2 != 1 && y%2 != 1 {
					continue
				}

				// Because triples are not produced in order of the length of their hypotenuse
				// we may see some triples with a hypotenuse > max followed by some with a
				// hypotenuse < max.  So ignore any with hypotenuse > max by following the
				// continue statement.  Once we get to a point where x + y > max*2 then we
				// can return because by that point we can be sure that we have got them all.
				if x+y > max*4 {
					// fmt.Printf("Exiting on: %d %d %d\n", x, y, z)
					return
				}
				// fmt.Printf("x y z: %d %d %d\n", x, y, z)
				if x > y {
					delta = x - y
				} else {
					delta = y - x
				}
				// filter for nearly equilateral triangels
				if (x > max && y > max) || delta > max {
					continue
				}
				*t = append(*t, Triple{x, y, z})
			}
		}
	}
}

func Challenge94() {
	var solution, perimeter int
	var perimeterStr, combinedStr string
	var triples = make([]Triple, 0)
	const maxPerim = 1_000_000_000

	triples = getTriples2(maxPerim/3 + 1)

	for _, triple := range triples {
		if triple[0] < triple[1] {
			perimeter = (triple[0] + triple[0] + triple[2] + triple[2])
			perimeterStr = fmt.Sprintf("{%d, %d, %d} with perimeter: %d\n", triple[0], triple[1], triple[2], perimeter)
		} else {
			perimeter = (triple[1] + triple[1] + triple[2] + triple[2])
			perimeterStr = fmt.Sprintf("{%d, %d, %d} with perimeter: %d\n", triple[1], triple[0], triple[2], perimeter)
		}
		if perimeter > maxPerim {
			break
		}
		solution += perimeter
		combinedStr += perimeterStr
	}
	fmt.Printf("%s", combinedStr)
	fmt.Printf("22 challenge 94 solution is: %d\n", solution)
}
