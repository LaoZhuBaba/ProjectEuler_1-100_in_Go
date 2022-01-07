package challenge71

import (
	"fmt"
	"math/big"
)

const max = 1_000_000

func Challenge71() {
	threeOver7 := big.NewRat(int64(3), int64(7))
	bestSoFar := new(big.Rat)
	for denom := int64(1); denom <= max; denom++ {
		num := denom * 3 / 7
		nOverD := big.NewRat(num, denom)
		if nOverD.Cmp(threeOver7) == -1 && bestSoFar.Cmp(nOverD) == -1 {
			bestSoFar.Set(nOverD)
		}
		// fmt.Printf("numerator: %d denom: %d %v %v\n", num, denom, nOverD, bestSoFar)
	}
	fmt.Printf("Challenge 71 solution is the numerator of: %v\n", bestSoFar)
}
