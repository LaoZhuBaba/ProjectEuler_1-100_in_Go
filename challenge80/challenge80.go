package challenge80

import (
	"fmt"
	"math/big"
)

func Challenge80() {
	var solution int
	for count := float64(1); count <= 100; count++ {
		bigCount := big.NewFloat(count).SetPrec(400)
		sr := new(big.Float).Sqrt(bigCount)
		if sr.IsInt() {
			continue
		}
		// I had to use 102 here.  One extra digit is needed because we skip the decimal dot.  I think
		// the other is to do with rounding: if we limit the exact number of digits then the last digit
		// may be rounded.  The challenge gave the total for the square root of two so I just made sure
		// that this was correct and calculated in exactly the same way for them all.
		srStr := sr.Text('g', 102)

		var n int
		var decDigitSum int
		// 101 because the decimal dot accounts for one rune.
		for countPlaces := 0; countPlaces < 101; countPlaces++ {
			k := srStr[countPlaces]
			if k == '.' {
				continue
			} else {
				fmt.Sscanf(string(k), "%d", &n)
				decDigitSum += n
			}
		}
		solution += decDigitSum
	}
	fmt.Printf("Challenge 80 solution is: %d", solution)
}
