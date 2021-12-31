// If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly
// three solutions for p = 120.

// {20,48,52}, {24,45,51}, {30,40,50}

// For which value of p ≤ 1000, is the number of solutions maximised?

package challenge39

import (
	"fmt"
	"math"
)

const p39Max = 1000

var squares [p39Max / 2]int
var tripletSums = make(map[int]int)

func Challenge39() {

	// Initialise array containing squares
	for n := 0; n < p39Max/2-2; n++ {
		squares[n] = n * n
	}
	for a := 1; a < p39Max/2-2; a++ {
		for b := a + 1; b < p39Max/2-2; b++ {
			aSquare := squares[a]
			bSquare := squares[b]
			cSquare := aSquare + bSquare
			c := int(math.Sqrt(float64(cSquare)))
			if c*c == cSquare {
				tripletSum := a + b + c
				if tripletSum <= p39Max {
					fmt.Printf("Found triple: %d, %d, %d\n", a, b, c)
					tripletSums[tripletSum]++
				}
			}
		}
	}
	var commonest, solution int
	for k, v := range tripletSums {
		if v > commonest {
			commonest = v
			solution = k
		}
	}
	fmt.Printf("Challenge 39 solution is %d with %d matches\n", solution, commonest)
}
