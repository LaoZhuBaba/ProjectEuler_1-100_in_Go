package challenge77

import (
	"euler/shared"
	"fmt"
)

// Return a generator function that when run repeated will return each prime number below n in order
// from high to low.  The generator will return 0 when there are no remaining primes.
func genPrimeList(n int) func() int {
	lastChecked := n
	return func() int {
		for lastChecked > 1 && !shared.IsPrime(lastChecked) {
			lastChecked--
		}
		if lastChecked <= 1 {
			return 0
		} else {
			defer func() { lastChecked-- }()
			return lastChecked
		}
	}
}

var counter int

// n is the number being analysed.  We need ceiling because in some cases we have already found
// solutions for numbers >= ceiling.  We may need to continue to find combinations for smaller numbers
// but we need to limit this.
func sumOfPrimes(n, ceiling, level int) {

	// indent := strings.Repeat("\t", level)
	// fmt.Printf("%ssumOfPrimes() called with n=%d (level %d)\n", indent, n, level)
	if n < 2 {
		// fmt.Printf("%ssumOfPrimes() called with n < 2\n", indent)
		return
	}
	if n < 4 {
		if n <= ceiling {
			counter++
			// fmt.Printf("%ssumOfPrimes() called with 2 or 3.  Increment counter and return\n", indent)
		}
		return
	}
	getPrime := genPrimeList(ceiling)
	// fmt.Printf("%sAbout to start for loop\n", indent)
	for {
		// If the prime which is returned is equal to n then n is a solution without further
		// subdivision.  So increment count but continue to recurse because n may itself
		// have valid prime divisions.
		prime := getPrime()
		// fmt.Printf("%sgetPrime() has returned %d\n", indent, prime)
		if prime == 0 {
			// fmt.Printf("%sbreak!\n", indent)
			break
		}
		if prime == n {
			// fmt.Printf("%sFound a solution with prime: %d.  Continuing to look for subdivisions of %d\n", indent, n, n)
			counter++
			continue
		}
		// Reaching here means we have found a prime which is at least two less than n and and can
		// therefore be fitted into n with the remainder passing to another recursion of sumOfPrimes.
		// We need to set the ceiling value to prime because solutions for values higher than prime
		// have already been found.  This covers situations where, for example we are passing 8 as the
		// first value and need to find 2, 2, 2, 2, but we don't want to find 2, 3, 3 because that will
		// already have been discovered as 3, 2, 2 in a previous iteration.
		sumOfPrimes(n-prime, prime, level+1)
	}
}

const max = 5000

func Challenge77() {
	var solution int
	for count := 4; ; count++ {
		counter = 0
		sumOfPrimes(count, count-1, 0)
		fmt.Printf("The number of possibilies for %d is %d\n", count, counter)
		if counter >= max {
			solution = count
			break
		}
	}
	fmt.Printf("The solution for challenge 77 is: %d\n", solution)
}
