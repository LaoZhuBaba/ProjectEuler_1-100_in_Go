//What is the 10,001 prime number?
package challenges

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
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
	if len(factors) == 1 {
		return true
	} else {
		return false
	}
}

func Challenge7() {
	count := 0
	for n := 2; true; n += 1 {
		if isPrime(n) {
			count += 1
			if count == 10001 {
				fmt.Printf("count is: %d\n", n)
				return
			}
		}
	}
}
