// If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.

// Not all numbers produce palindromes so quickly. For example,

// 349 + 943 = 1292,
// 1292 + 2921 = 4213
// 4213 + 3124 = 7337

// That is, 349 took three iterations to arrive at a palindrome.

// Although no one has proved it yet, it is thought that some numbers, like 196, never produce a palindrome.
// A number that never forms a palindrome through the reverse and add process is called a Lychrel number. Due
// to the theoretical nature of these numbers, and for the purpose of this problem, we shall assume that a number
// is Lychrel until proven otherwise. In addition you are given that for every number below ten-thousand, it will
// either (i) become a palindrome in less than fifty iterations, or, (ii) no one, with all the computing power
// that exists, has managed so far to map it to a palindrome. In fact, 10677 is the first number to be shown to
// require over fifty iterations before producing a palindrome: 4668731596684224866951378664 (53 iterations,
// 28-digits).

// Surprisingly, there are palindromic numbers that are themselves Lychrel numbers; the first example is 4994.

// How many Lychrel numbers are there below ten-thousand?
package challenges

import (
	"fmt"
	"math/big"
)

func reverseDecimalBig(n *big.Int) *big.Int {
	// fmt.Printf("reverseDecimal called with %d\n", n)
	// fmt.Printf("in reverseDecimalBig: n is %d\n", n)
	temp1 := new(big.Int)
	temp2 := new(big.Int)
	r := big.NewInt(int64(0))
	for n.Cmp(big.NewInt(0)) != 0 {
		r.Add(temp1.Mul(r, big.NewInt(10)), temp2.Mod(n, big.NewInt(10)))
		n.Div(n, big.NewInt(10))
	}
	// fmt.Printf("reverseDecimal is returning: %d\n", r)
	return r
}

func isPalindromeBig(n *big.Int) bool {
	nCopy := new(big.Int).Set(n)
	n2 := reverseDecimalBig(nCopy)
	return (n2.Cmp(n)) == 0
}

func isLychrel(n int64) bool {
	//fmt.Printf("%d\n", n)
	nBig := big.NewInt(n)
	for count := 1; count < 50; count++ {
		//fmt.Printf("  %d\n", nBig)
		nBig.Set(reverseAndAddBig(nBig))
		if isPalindromeBig(nBig) {
			return false
		}
	}
	return true
}

func reverseAndAddBig(n *big.Int) *big.Int {
	nCopy := new(big.Int).Set(n)
	return n.Add(n, reverseDecimalBig(nCopy))
}
func Challenge55() {
	var solution int
	for count := int64(1); count < 10_000; count++ {
		// bigCount := big.NewInt(count)
		// fmt.Printf("%d \n", bigCount) //reverseDecimalBig(bigCount))
		// reversed := reverseDecimalBig(bigCount)
		// // fmt.Printf("reversed: %d\n", reverseDecimalBig(bigCount))
		// fmt.Printf("reversed: %d\n", reversed)
		// if isPalindromeBig(bigCount) {
		// 	fmt.Printf("%d is a palindrome!\n", bigCount)
		// }
		isLy := isLychrel(count)
		if isLy {
			fmt.Printf("Is %d a Lychrel number? %v\n", count, isLy)
			solution++
		}
	}
	fmt.Printf("solution is: %d\n", solution)

}
