package challenge100

import (
	"fmt"

	"lukechampine.com/uint128"
)

func intDiviedBySqRootOf2(n uint128.Uint128) uint128.Uint128 {
	// 1607521/1136689 is a fractional approximation of the square root of 2
	return n.Mul64(1136689).Div64(1607521)
}
func Challenge100() {

	const max uint64 = 10_000_000_000_000
	var denom1, denom2, numer1, numer2 uint128.Uint128
	// Experimentation shows that each successive solution increases by a factor of approx 5.82
	for denom1 = uint128.From64(2); denom1.Cmp64(max) == -1; denom1 = denom1.Mul64(582).Div64(100) {
		for {
			denom2 = denom1.Add64(1)
			numer1 = intDiviedBySqRootOf2(denom1)
			numer2 = numer1.Add64(1)
			dd := denom1.Mul(denom2)
			nn := numer1.Mul(numer2).Mul64(2)
			if dd.Cmp(nn) == 0 {
				// Detect if the solution is > 10^12
				if denom2.Cmp64(999_999_999_999) == 1 {
					fmt.Printf("challenge 100 solution is: %s\n", numer2.String())
					return
				}
				break
			}
			denom1 = denom1.Add64(1)
		}
	}
}
