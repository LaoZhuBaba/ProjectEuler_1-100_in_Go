package challenge53

import (
	"euler/shared"
	"fmt"
	"math/big"
)

func combinatoricSelection(n, r int64) *big.Int {
	nFactorial := shared.Factorial(n)
	rFactorial := shared.Factorial(r)
	nMinusRFactorial := shared.Factorial(n - r)
	res1 := rFactorial.Mul(rFactorial, nMinusRFactorial)
	return nFactorial.Div(nFactorial, res1)
}
func Challenge53() {
	solution := 0
	million := big.NewInt(1_000_000)
	for n := int64(100); n >= 2; n-- {
		for r := n - 1; r > 1; r-- {
			result := combinatoricSelection(n, r)
			if result.Cmp(million) == 1 {
				fmt.Printf("%d (%d %d)\n", result, n, r)
				solution++
			}
		}
	}
	fmt.Printf("Package 53 solution is: %d\n", solution)
}
