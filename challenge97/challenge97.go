package challenge97

import "fmt"

// Calculate last 10 digits of the prime number 28433*(2^7830457)+1

func Challenge97() {
	var n uint64 = 28433
	for count := 1; count <= 7830457; count++ {
		// eliminate the most significant digit to prevent overrunning 64 bits
		// With this amount of precision or more the last 10 digits don't change.
		if n > 100_000_000_000_000_000 {
			n -= 100_000_000_000_000_000
		}
		n *= 2
	}
	// Divide by 10,000,000,000 strips off the 10 rightmost digits.  Then multiply
	// to add 10 zeros to the right.  If we subtract this from n we get the 10
	// last digits of n
	fmt.Printf("challenge 98 solution: %d\n", n-(n/10_000_000_000*10_000_000_000)+1)
}
