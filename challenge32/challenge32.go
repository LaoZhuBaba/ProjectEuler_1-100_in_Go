// We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once;
// for example, the 5-digit number, 15234, is 1 through 5 pandigital.
// The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1
// through 9 pandigital.
// Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
// HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
package challenge32

import (
	"euler/shared"
	"fmt"
)

// func intToIntSlice(n int) []int is defined in challenge30.go

func intLen(n int) int {
	return len(shared.IntToIntSlice(n))
}
func Challenge32() {
	m := make(map[int]bool)
	for a := 1; a < 9999; a++ {
		if shared.IntContainsDigit0(a) {
			continue
		}
		if shared.IntContainsDuplicateDigits(a) {
			continue
		}
		aLen := intLen(a)
		for b := 1; b < 9999; b++ {
			bLen := intLen(b)
			product := a * b
			if intLen(product) != 9-aLen-bLen {
				continue
			}
			// From here we know that the combined length of the product and factors is 9
			if shared.IntContainsDigit0(b) {
				continue
			}
			if shared.IntContainsDuplicateDigits(b) {
				continue
			}
			ab := append(shared.IntToIntSlice(a), shared.IntToIntSlice(b)...)
			if shared.ContainsDuplicates(ab) {
				continue
			}
			if shared.IntContainsDigit0(product) {
				continue
			}
			if shared.IntContainsDuplicateDigits(product) {
				continue
			}
			if shared.ContainsDuplicates(append(shared.IntToIntSlice(product), ab...)) {
				continue
			}
			// m[a] = true
			// m[b] = true
			m[product] = true
			fmt.Printf("%d x %d = %d\n", a, b, product)
		}
	}
	fmt.Printf("%v\n", m)
	var solution int
	for k := range m {
		solution += k
	}
	fmt.Printf("Challenge 32 solution is: %d\n", solution)
}
