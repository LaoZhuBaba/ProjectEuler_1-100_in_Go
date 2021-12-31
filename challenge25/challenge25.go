// The Fibonacci sequence is defined by the recurrence relation:

// Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
// Hence the first 12 terms will be:

// F1 = 1
// F2 = 1
// F3 = 2
// F4 = 3
// F5 = 5
// F6 = 8
// F7 = 13
// F8 = 21
// F9 = 34
// F10 = 55
// F11 = 89
// F12 = 144
// The 12th term, F12, is the first term to contain three digits.

// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

package challenge25

import (
	"fmt"
	"math/big"
)

func Challenge25() {
	n1 := big.NewInt(0)
	n2 := big.NewInt(1)
	temp := big.NewInt(0)

	for count := 1; ; count++ {
		if len(n2.Text(10)) >= 1000 {
			fmt.Printf("%d\n", n2)
			fmt.Printf("Challenge 25 solution: %d\n", count)
			break
		}
		// fmt.Printf("index: %d is %d\n", count, n2)
		temp.Add(n1, n2)
		n1.Set(n2)
		n2.Set(temp)
	}
}
