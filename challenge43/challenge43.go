// The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the
// digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.

// Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

// d2d3d4=406 is divisible by 2
// d3d4d5=063 is divisible by 3
// d4d5d6=635 is divisible by 5
// d5d6d7=357 is divisible by 7
// d6d7d8=572 is divisible by 11
// d7d8d9=728 is divisible by 13
// d8d9d10=289 is divisible by 17
// Find the sum of all 0 to 9 pandigital numbers with this property.
package challenge43

import (
	"euler/shared"
	"fmt"
)

func Challenge43() {
	count := 0
	for n := 1_023_456_789; n <= 9_876_543_210; n++ {
		first2 := n / 100_000_000
		first3 := n / 10_000_000
		first4 := n / 1_000_000
		first5 := n / 100_000
		first6 := n / 10_000
		first7 := n / 1_000
		first8 := n / 100
		first9 := n / 10
		// If first four is an odd number then increment first four by one
		if first4%2 != 0 {
			n = (first4 + 1) * 1_000_000
			continue
		}
		three45 := first5 % 1000
		if three45%3 != 0 {
			three45 = three45 - (three45 % 3) + 3 // Increment to the next multiple of 3
			n = (first2*1000 + three45) * 100_000
			continue
		}
		four56 := first6 % 1000
		if four56%5 != 0 {
			four56 = four56 - (four56 % 5) + 5 // Increment to the next multiple of 5
			n = (first3*1000 + four56) * 10_000
			continue
		}
		five67 := first7 % 1000
		if five67%7 != 0 {
			five67 = five67 - (five67 % 7) + 7
			n = (first4*1000 + five67) * 1_000
			continue
		}
		six78 := first8 % 1000
		if six78%11 != 0 {
			six78 = six78 - (six78 % 11) + 11
			n = (first5*1000 + six78) * 100
			continue
		}
		seven89 := first9 % 1000
		if seven89%13 != 0 {
			seven89 = seven89 - (seven89 % 13) + 13
			n = (first6*1000 + seven89) * 10
			continue
		}
		eight910 := n % 1000
		if eight910%17 != 0 {
			continue
		}

		if !shared.IntContainsDuplicateDigits(n) {
			fmt.Printf("%d\n", n)
			count += n
		}
	}
	fmt.Printf("Challenge 43 solution is: %d\n", count)
}
