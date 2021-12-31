// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package challenge6

import "fmt"

func Challenge6() {
	sum, sumSquare := 0, 0
	for n := 1; n <= 100; n++ {
		sum += n
		sumSquare += n * n
	}
	fmt.Printf("Square of sums: %d\n", sum*sum)
	fmt.Printf("Sum of squares: %d\n", sumSquare)
	fmt.Printf("Challenge 6 solution: difference is %d\n", sum*sum-sumSquare)

}
