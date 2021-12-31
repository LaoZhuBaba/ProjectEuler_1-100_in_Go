package challenge65

import (
	"fmt"
	"math/big"
)

const (
	limit = 100
)

// Recursively evaluate a "continuous" fraction
func contFractToSimplFract(s []int64) (num, denom *big.Int) {
	if len(s) == 1 {
		return big.NewInt(1), big.NewInt(s[0])
	}
	num, denom = contFractToSimplFract(s[1:])

	// The next line is equivalent to: return denom, s[0]*denom + num
	return denom, new(big.Int).Add(new(big.Int).Mul(big.NewInt(s[0]), denom), num)
}

// This is the original version which uses uint64.  However to reach the limit of
// 100 requires numbers which exceed the limit of uint64.
//
// func contFractToSimplFract(s []uint64) (num, denom uint64) {
// 	if len(s) == 1 {
// 		fmt.Printf("terminal returning %d, %d with s[0] equal to %d\n", 1, s[0], s[0])
// 		return 1, s[0]
// 	}
// 	num, denom = contFractToSimplFract(s[1:])
// 	fmt.Printf("non-terminal returning %d, %d with s[0] equal to: %d\n", s[0]*denom+num, denom, s[0])
// 	return denom, s[0]*denom + num
// }

// You can calculate the value of the constant "e" by calculating the result of a "continued fraction"
// following the pattern [2; 1, 2, 1, 1, 4, 1, 1, 6, 1, 1, 8, 1 ... 1, n, 1] to an arbitrary level of
// precision.  This challenge asks for a precision of 100 values.  The following function populates
// a slice with the correct 100 numbers.
func populateSlice() []int64 {
	e := make([]int64, limit)
	e[0] = 2
	var evenCounter int64 = 2
	for count := 1; count < limit; count++ {
		if (count+1)%3 == 0 {
			e[count] = evenCounter
			evenCounter += 2
		} else {
			e[count] = 1
		}
	}
	return e
}

func Challenge65() {
	s := populateSlice()
	// fmt.Printf("%v\n", s)
	denom, num := contFractToSimplFract(s)
	fmt.Printf("num: %d, denom: %d\n", num, denom)
	floatDenom := new(big.Float).SetInt(denom)
	floatNum := new(big.Float).SetInt(num)
	var solution rune
	for _, digit := range num.String() {
		solution += digit - '0'
	}
	e := new(big.Float).Quo(floatNum, floatDenom)
	fmt.Printf("Approximate value of e is: %.100f\n", e)
	fmt.Printf("Challenge 65 solution is: %d\n", solution)
}
