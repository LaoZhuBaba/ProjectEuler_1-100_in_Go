package challenge57

import (
	"fmt"
	"math/big"
)

func bigFractRootTwo(num, denom *big.Int) (nextNum *big.Int, nextDenom *big.Int) {
	temp := new(big.Int)
	dblDenom := new(big.Int).Set(temp.Mul(denom, big.NewInt(int64(2))))
	numMinusDenom := new(big.Int).Set(temp.Sub(num, denom))
	nextDenom = new(big.Int).Set(temp.Add(dblDenom, numMinusDenom))
	nextNum = new(big.Int).Set(temp.Add(nextDenom, denom))
	return nextNum, nextDenom
}

// func fractRootTwo(num, denom int) (nextNum int, nextDenom int) {
// 	nextDenom = (denom * 2) + (num - denom)
// 	nextNum = nextDenom + denom
// 	return nextNum, nextDenom
// }

func Challenge57() {
	var solution int
	num, denom := big.NewInt(3), big.NewInt(2)
	for count := 1; count < 1000; count++ {
		if len(num.String()) > len(denom.String()) {
			solution++
		}
		//		fmt.Printf("%d / %d\n", num, denom)
		num, denom = bigFractRootTwo(num, denom)
	}
	fmt.Printf("Challenge 57 solution: %d\n", solution)
}
