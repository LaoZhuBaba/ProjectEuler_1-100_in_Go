package challenge86

import (
	"euler/shared"
	"fmt"
)

type Triple [3]int
type Tuplet [2]int

var triples = make([]Triple, 0)

func PythagTriples(t *[]Triple, max int) {
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
				if (x > max && y > max) || delta > max {
					continue
				}

				for count := 1; x*count <= max || y*count <= max || z*count <= max; count++ {
					*t = append(*t, Triple{x * count, y * count, z * count})
				}
			}
		}
	}

}

func numberToAddends(num, max int) []Tuplet {
	ret := make([]Tuplet, 0)
	if max >= num {
		max = num - 1
	}
	for a, b := max, num-max; a-b > -1; a-- {
		//fmt.Printf("a, b: %d %d\n", a, b)
		ret = append(ret, Tuplet{a, b})
		b++
	}
	return ret
}
func Challenge86() {
	max, delta := 100, 400
	var lower bool
	// This loop implements a rough kind of binary search to avoid incrementing through every
	// possible value from 1 upwards.  We grow max at an increasing rate so long as the solutions
	// we find are under our target of 1 million.  Once we cross 1 million we decrease max at
	// a decreasing rate and gradually zero in on the minimum number which provide the solution.
	for {
		var m = make(map[Triple]bool)
		PythagTriples(&triples, max)
		var addends []Tuplet
		for _, triplet := range triples {
			addends = numberToAddends(triplet[0], triplet[1])
			for _, addend := range addends {
				if triplet[1] <= max {
					m[Triple{addend[1], addend[0], triplet[1]}] = true
				}
			}
			addends = numberToAddends(triplet[1], triplet[0])
			for _, addend := range addends {
				if triplet[0] <= max {
					m[Triple{addend[1], addend[0], triplet[0]}] = true
				}
			}
		}
		fmt.Printf("When max length is %d, there are %d compliant cuboids\n", max, len(m))
		if len(m) < 1_000_000 {
			if delta == 1 && !lower {
				fmt.Printf("Challenge 86 solution is: %d\n", max+1)
				return
			}
			lower = true
			max += delta
		} else {
			lower = false
			delta /= 2
			if delta == 0 {
				delta = 1
			}
			max -= delta
		}
		// fmt.Printf("map is: %v\n", m)
	}
}
