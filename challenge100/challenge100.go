package challenge100

import (
	"fmt"
)

func intDiviedBySqRootOf2(n uint64) uint64 {
	return n * 470832 / 665857
}
func Challenge100() {
	//gpf := shared.GetPrimeFactorsUint64
	for denom1 := uint64(1); denom1 < 20_000_000_000; denom1++ {
		denom2 := denom1 + 1
		numer1 := intDiviedBySqRootOf2(denom1)
		numer2 := numer1 + 1
		// denom1Factors := gpf(denom1)
		// denom2Factors := gpf(denom2)
		// numer1Factors := gpf(numer1)
		// numer2Factors := gpf(numer2)

		if denom1*denom2 == numer1*numer2*2 {
			fmt.Printf("%d/%d * %d/%d = 1/2\n", numer1, denom1, numer2, denom2)
			fmt.Printf("%d - %d = %d\n", denom2, numer1, denom2-numer1)
			// fmt.Printf("%v\n", denom1Factors)
			// fmt.Printf("%v\n", denom2Factors)
			// fmt.Printf("%v\n", numer1Factors)
			// fmt.Printf("%v\n", numer2Factors)
		}
	}
}

// 20/14 * 21/15 = 420
// 84/60 * 85/61

// 84/119 * 85/120
// 84/7 * 5/120
// 7/7 * 5/10
