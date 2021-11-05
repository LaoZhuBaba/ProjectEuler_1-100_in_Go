// Find the sum of all the primes below two million

package challenges

import "fmt"

func Challenge10() {

	grandTotal := 0
	for n := 2; n < 2_000_000; n++ {

		if isPrime(n) {
			grandTotal += n
			// fmt.Printf("%d is prime\n", n)
		}
	}
	fmt.Printf("Grand total: %d\n", grandTotal)
}
