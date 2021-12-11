// A googol (10^100) is a massive number: one followed by one-hundred zeros; 100^100 is almost unimaginably
// large: one followed by two-hundred zeros. Despite their size, the sum of the digits in each number is
// only 1.

// Considering natural numbers of the form, a^b, where a, b < 100, what is the maximum digital sum?

package challenges

import (
	"fmt"
	"math/big"
)

func bigPow(n, o *big.Int) *big.Int {
	return n.Exp(n, o, nil)
}
func bigStrSum(bigN *big.Int) int {

	var sum int
	bigStr := bigN.String()
	for _, digit := range bigStr {
		sum += int(digit) - '0'
		// fmt.Printf("%c,", digit)
	}
	return sum
}

func Challenge56() {
	var solution int
	for n1 := 1; n1 < 100; n1++ {
		for n2 := 1; n2 < 100; n2++ {
			bigN1 := big.NewInt(int64(n1))
			bigN2 := big.NewInt(int64(n2))
			sum := bigStrSum(bigPow(bigN1, bigN2))
			if sum > solution {
				solution = sum
			}
			fmt.Printf("%d^%d: %d\n", n1, n2, sum)

		}
	}
	fmt.Printf("solution: %d\n", solution)
}
