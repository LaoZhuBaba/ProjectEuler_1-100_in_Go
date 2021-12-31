// We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n
// exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

// What is the largest n-digit pandigital prime that exists?

package challenge41

import (
	"euler/shared"
	"fmt"
)

func intContainsDigit(n, d int) bool {
	nSlice := shared.IntToIntSlice(n)
	for _, digit := range nSlice {
		if digit == d {
			return true
		}
	}
	return false
}

func Challenge41() {
	var maxN int
	for n := 1234567; n <= 12345678; n++ {
		if n%2 == 0 {
			continue
		}
		if n%3 == 0 {
			continue
		}
		if intContainsDigit(n, 0) {
			continue
		}
		if intContainsDigit(n, 9) {
			continue
		}
		if intContainsDigit(n, 8) {
			continue
		}
		if shared.IntContainsDuplicateDigits(n) {
			continue
		}
		if !shared.IsPrime(n) {
			continue
		}
		fmt.Printf("%d\n", n)
		if n > maxN {
			maxN = n
		}
	}
	fmt.Printf("Challenge 41 solution: %d\n", maxN)

}
