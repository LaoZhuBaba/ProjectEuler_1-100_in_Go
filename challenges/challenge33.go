// The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may
// incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

// We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

// There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing
// two digits in the numerator and denominator.

// If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
package challenges

import "fmt"

func Challenge33() {

	rawNumerator := 1
	rawDenominator := 1
	for a := 10; a < 100; a++ {
		aDigits := intToIntSlice(a)
		if aDigits[0] == aDigits[1] {
			continue
		}
		for b := 10; b < 100; b++ {
			bDigits := intToIntSlice(b)
			if bDigits[0] == bDigits[1] {
				continue
			}
			if aDigits[1] == bDigits[0] {
				if float64(a)/float64(b) == float64(aDigits[0])/float64(bDigits[1]) {
					fmt.Printf("%d / %d == %d / %d\n", a, b, aDigits[0], bDigits[1])
					rawNumerator *= aDigits[0]
					rawDenominator *= bDigits[1]

				}
			}
		}
	}
	fmt.Printf("%d %d\n", rawNumerator, rawDenominator)
	for count := rawNumerator; count >= 2; count-- {
		if rawDenominator%count == 0 && rawNumerator%count == 0 {
			fmt.Printf("solution is: %d\n", rawDenominator/count)
			break
		}
	}
}
