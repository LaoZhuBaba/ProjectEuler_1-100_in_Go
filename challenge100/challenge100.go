package challenge100

import (
	"fmt"

	"lukechampine.com/uint128"
)

func intDiviedBySqRootOf2(n uint64) uint64 {
	// 1607521/1136689 is a fractional approximation of the square root of 2
	n128 := uint128.From64(n)
	// Use uint128 to avoid losing significant bits during this calculation
	return n128.Mul64(1136689).Div64(1607521).Lo
}
func Challenge100() {

	const max uint64 = 10_000_000_000_000
	var denom1, denom2, numer1, numer2 uint64
	// Experimentation shows that each successive solution increases by a factor of approx 5.82
	for denom1 = 2; denom1 < max; denom1 = denom1 * 582 / 100 {
		for ; ; denom1++ {
			// fmt.Printf("%d %d %d %d\n", numer1, denom1, numer2, denom2)
			numer1 = intDiviedBySqRootOf2(denom1)
			// If numer1 is not even then no chance of a result.
			if numer1%2 != 0 {
				continue
			}
			denom2 = denom1 + 1
			numer2 = numer1 + 1
			dd := denom1 * denom2
			nn := numer1 * numer2 * 2
			if dd == nn {
				// Detect if the solution is > 10^12
				fmt.Printf("%d / %d * %d / %d == 2\n", numer1, denom1, numer2, denom2)
				if denom2 > 999_999_999_999 {
					fmt.Printf("challenge 100 solution is: %d\n", numer2)
					return
				}
				break
			}
		}
	}
}
