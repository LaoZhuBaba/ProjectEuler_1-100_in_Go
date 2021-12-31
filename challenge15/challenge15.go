package challenge15

// Calculate (2*n)! / (n!)^2
// for n = 20

import (
	"fmt"
	"math/big"
)

var one = big.NewInt(1)

func factorial(n int64) *big.Int {
	if n == 1 {
		return one
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))

}
func Challenge15() {

	theNumber := int64(20)
	theNumberTimes2 := 2 * theNumber
	theNumberTimes2Factorial := factorial(theNumberTimes2)
	theNumberFactorial := factorial(theNumber)

	// We could also use theNumberFactorialSquared := new(big.Int), but that returns *Int rather than Int
	var theNumberFactorialSquared big.Int
	//theNumberFactorialSquared := new(big.Int)
	theNumberFactorialSquared.Mul(theNumberFactorial, theNumberFactorial)
	var finalAnswer big.Int
	finalAnswer.Div(theNumberTimes2Factorial, &theNumberFactorialSquared)
	// fmt.Printf("Final answer is: %d\n", &finalAnswer)
	fmt.Printf("Challenge 15 solution is: %s\n", finalAnswer.String())
}
