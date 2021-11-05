// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package challenges

import (
	"fmt"
	"math"
)

func getFactors(n int) *[]int {
	factors := make([]int, 0)
	var i int
	for i = 2; i < (int(math.Sqrt(float64(n))) + 1); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
			i--
		}
	}
	factors = append(factors, n)
	return &factors
}

func Challenge5() {
	const X int = 20
	total := 1
	for n := 2; n <= X; n++ {
		factors := *getFactors(n)
		// We only care about primes (i.e., length == 0)
		if len(factors) == 1 {
			prime := factors[0]
			primeMultiple := prime
			// Keep multiplying each prime by itself so long as the total is <= X
			// e.g., if X is 20 then 2 will become 16 and 3 will become 9.  5 and above remain unchanged.
			for primeMultiple*prime <= X {
				primeMultiple *= prime

			}
			total *= primeMultiple
		}
	}

	for count := 2; count <= X; count += 1 {
		fmt.Printf("%d / %d = %d with a remainder of %d\n", total, count, total/count, total%count)
	}
}
