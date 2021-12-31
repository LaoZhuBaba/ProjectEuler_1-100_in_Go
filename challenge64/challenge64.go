package challenge64

import (
	"fmt"
	"math/big"
)

// Found by trial and error that I needed a precision of at least 777 in order to
// compute the solution which Project Euler accepted as correct.  The longest
// sequence for an integer <= 10,000 is 9949 which repeats after 217 numbers
// so we need very good precision.
const precision = 777

// Calculate the continued fraction square root pattern for integer i.  Returning
// a slice of ints which represents the integer square root followed by a single
// iteration of the continued fraction sequence.  Second return value is the length
// of the sequence.
func cfsr(i int) ([]int, int) {
	var si []int    // slice of ints to store digit part which will be returned
	var ss []string // slice of strings used to detect when a repeat occurs
	var count int
	// f is just i converted to a big float so we square root it
	f := big.NewFloat(float64(i)).SetPrec(precision)
	residualNumber := f.Sqrt(f)
	//  If i is a square number then it's not interesting
	if residualNumber.IsInt() {
		sqrtInt, _ := residualNumber.Int64()
		return []int{int(sqrtInt)}, 0
	}

	for {
		// This is just to catch any infinite loops
		if count > 400 {
			fmt.Printf("Too many iterations without finding a repeat pattern!\n")
			break
		}
		// Split the square root into an integer part and a decimal fraction.
		intPart, _ := residualNumber.Int64()
		intPartAsFloat := big.NewFloat(float64(intPart)).SetPrec(precision)
		fractPart := big.NewFloat(0).SetPrec(precision).Sub(residualNumber, intPartAsFloat)
		si = append(si, int(intPart))
		ss = append(ss, fmt.Sprintf("%.20f", fractPart))
		// On each iteration we take the decimal fraction part of residualNumber and recipricate it
		// and then assign back to residual number.  The integer part of residualNumber is what we
		// need to collect in si, but we also need to record the decimal fraction part in ss because
		// when we encounter a repeated value in ss then this means we have reached the end of the
		// repeat pattern and can exit.
		residualNumber.Set(big.NewFloat(0).SetPrec(precision).Quo(big.NewFloat(1), fractPart))
		// As soon as any value in ss is a repeat of ss[0] then we have reached the start of a second
		// sequence iteration.  The length of the repeat pattern is one less than the length of the
		// si/ss slices.
		if len(ss) != 1 && ss[len(ss)-1] == ss[0] {
			return si, len(ss) - 1
		}
		count++
	}
	return si, 0
}

func Challenge64() {
	var solution int
	var si []int
	var cycle int
	var longestCycle int
	for count := 1; count <= 10_000; count++ {
		si, cycle = cfsr(count)
		// fmt.Printf("%d: %v -- %d\n", count, si, cycle)
		if cycle%2 != 0 {
			solution++
		}
		if cycle > longestCycle {
			longestCycle = cycle
			fmt.Printf("longestCycle so far is: %d\n", longestCycle)
			fmt.Printf("%d: %v\n", count, si)
		}
	}
	fmt.Printf("solution is: %d\n", solution)
}
