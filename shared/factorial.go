package shared

import "math/big"

func Factorial(n int64) *big.Int {
	var one = big.NewInt(1)

	if n == 1 {
		return one
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, Factorial(n-1))

}
