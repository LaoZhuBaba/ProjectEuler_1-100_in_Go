package challenges

// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
// What is the sum of the digits of the number 2^1000?

import (
	"fmt"
	"math/big"
)

func Challenge16() {

	one := big.NewInt(int64(1))
	two := big.NewInt(int64(2))
	oneThousand := big.NewInt(int64(1000))
	bigN := big.NewInt(int64(2))

	for count := big.NewInt(int64(2)); count.Cmp(oneThousand) != 1; count.Add(count, one) {
		bigN = bigN.Mul(bigN, two)
		// fmt.Printf("count is %d, power of 2 is %s\n", count, bigN)
	}
	bigNStr := bigN.String()
	fmt.Printf("%s\n", bigNStr)
	sum := 0
	for _, r := range bigNStr {
		sum += int(r) - '0'
	}
	fmt.Printf("Sum is: %d\n", sum)
}
