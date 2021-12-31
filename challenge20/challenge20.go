// n! means n × (n − 1) × ... × 3 × 2 × 1

// For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

// Find the sum of the digits in the number 100!

package challenge20

import (
	"euler/shared"
	"fmt"
)

func Challenge20() {
	// factoria() is defined in challenge15.go
	answer := fmt.Sprintf("%s", shared.Factorial(100))
	var total int
	fmt.Printf("%s\n", answer)
	for _, r := range answer {
		total += int(r - '0')
	}
	fmt.Printf("Challenge 20 solution is: %d\n", total)
}
