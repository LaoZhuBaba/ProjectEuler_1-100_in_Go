// We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n
// exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

// What is the largest n-digit pandigital prime that exists?

package challenges

import "fmt"

func intContainsDigit(n, d int) bool {
	nSlice := intToIntSlice(n)
	for _, digit := range nSlice {
		if digit == d {
			return true
		}
	}
	return false
}

func Challenge41() {

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
		if intContainsDuplicateDigits(n) {
			continue
		}
		if !isPrime(n) {
			continue
		}
		fmt.Printf("%d\n", n)
	}

}
