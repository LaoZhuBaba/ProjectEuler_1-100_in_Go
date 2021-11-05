// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
// The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

// Evaluate the sum of all the amicable numbers under 10000.

package challenges

import (
	"fmt"
	"math"
)

const c21Max = 10_000

func factorise(n int, list *[]int) {
	sr := int(math.Sqrt(float64(n)))
	*list = append(*list, 1)
	for c := 2; c <= sr; c++ {
		if n%c == 0 {
			if c != n/c {
				*list = append(*list, c, n/c)
			} else {
				*list = append(*list, c)
			}
		}
	}
}

func sumOfList(l *[]int) int {
	sum := 0
	for _, n := range *l {
		sum += n
	}
	return sum
}

var factorSumList [c21Max + 1]int

func Challenge21() {
	list := new([]int)
	var factorSum int
	var total int
	// Calculate factors for each number from 1 to c21Max then sum them and store in factorSumList
	for c := 2; c <= c21Max; c++ {
		factorise(c, list)
		// fmt.Printf("factors of %d are: %v\n", c, *list)
		factorSum = sumOfList(list)
		factorSumList[c] = factorSum
		// fmt.Printf("sum of factors is: %d\n", factorSum)
		if factorSum < c {
			if c == factorSumList[factorSum] {
				//fmt.Printf("factorSumList[factorSum] is: %d\n", factorSumList[factorSum])
				fmt.Printf("Found amicable pair: %d & %d\n", c, factorSum)
				total = total + c + factorSum
			}

		}
		//fmt.Printf()

		*list = nil
	}
	fmt.Printf("Solution is: %d\n", total)
}
