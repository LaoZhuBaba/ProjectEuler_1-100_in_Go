// Find the sum of all the primes below two million

package challenge10

import (
	"euler/shared"
	"fmt"
)

func Challenge10() {

	grandTotal := 0
	for n := 2; n < 2_000_000; n++ {

		if shared.IsPrime(n) {
			grandTotal += n
			// fmt.Printf("%d is prime\n", n)
		}
	}
	fmt.Printf("Challenge 10 solution: %d\n", grandTotal)
}
