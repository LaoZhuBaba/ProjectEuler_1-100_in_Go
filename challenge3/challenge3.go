// What is the largest prime factor of the number 600851475143?”
package challenge3

import (
	"fmt"
	"math"
)

func largest(factors []int) int {
	fmt.Printf("factors is: %v\n", factors)
	max := 0
	for _, factor := range factors {
		if factor > max {
			max = factor
		}
	}
	return max
}

func printFactors(n int) int {
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

	return largest(factors)

}
func Challenge3() {
	fmt.Printf("Challenge 3 solution is: %d\n", printFactors(600851475143))

}
