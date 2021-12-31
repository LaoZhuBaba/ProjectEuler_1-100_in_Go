// An irrational decimal fraction is created by concatenating the positive integers:

// 0.123456789101112131415161718192021...

// It can be seen that the 12th digit of the fractional part is 1.

// If dn represents the nth digit of the fractional part, find the value of the following expression.

// d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000

package challenge40

import (
	"fmt"
	"math"
)

// var n int
// var nStr string

func createIterator() func() string {
	var n int
	var nStr string

	return func() string {

		if nStr == "" {
			n++
			nStr = fmt.Sprintf("%d", n)
		}
		retStr := nStr[0:1]
		if len(nStr) > 1 {
			nStr = nStr[1:]
		} else {
			nStr = ""
		}
		return retStr

	}
}

func Challenge40() {
	solution := 1
	iter := createIterator()
	for n := 1; ; n++ {
		// fmt.Printf("%s\n", c40Increment())
		log10 := math.Log10(float64(n))
		isExp10 := int(log10*1_000_000) == int(log10)*1_000_000
		i := iter()
		var iNum int
		fmt.Printf("%d %s %v\n", n, i, isExp10)
		if isExp10 {
			fmt.Sscanf(i, "%d", &iNum)
			solution *= iNum
		}
		if n == 1_000_000 {
			break
		}
	}
	fmt.Printf("Challenge 40 solution: %d\n", solution)

}
