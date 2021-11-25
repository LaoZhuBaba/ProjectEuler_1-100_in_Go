// The prime 41, can be written as the sum of six consecutive primes:

// 41 = 2 + 3 + 5 + 7 + 11 + 13
// This is the longest sum of consecutive primes that adds to a prime below one-hundred.

// The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

// Which prime, below one-million, can be written as the sum of the most consecutive primes?

package challenges

import "fmt"

const c50Max = 1_000_000

// Take a (pointer to) a slice of ints and return the sum of the range
// from start to end or zero if max is exceeded
func sumOfRange(r *[]int, start, end, max int) int {
	var sum int
	//	fmt.Printf("HIT ME\n")
	for n := start; n <= end; n++ {
		sum += (*r)[n]
		if sum > max {
			return 0
		}
	}
	return sum
}

func Challenge50() {
	var primes = make([]int, 0)

	// Build a slice of ints containing all primes below  or equal to c50Max
	for n := 2; n <= c50Max; n++ {
		if isPrime(n) {
			primes = append(primes, n)
		}
	}
	var sum int
	//  Start with the maxium possible length and on each iteration reduce by one.
	//  Within this the second loop runs sumOfRange() on every possible range of that length.
	//  E.g., for the whole of len(primes) there is only one range, but for len(primes)-1
	//  there are two ranges, etc.  Return as soon as we find a range that sums to a prime.
	for l := len(primes); l > 6; l-- {
		start := 0
		for end := start + l; end < len(primes); start, end = start+1, end+1 {
			sum = sumOfRange(&primes, start, end, c50Max)
			if isPrime(sum) {
				fmt.Printf("%d primes from %d to %d sum to prime %d\n", end-start, primes[start], primes[end], sum)
				return
			}
		}

	}

}
