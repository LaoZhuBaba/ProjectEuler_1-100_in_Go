// It was proposed by Christian Goldbach that every odd composite number can be written as the sum of
// a prime and twice a square.

// 9 = 7 + 2×12
// 15 = 7 + 2×22
// 21 = 3 + 2×32
// 25 = 7 + 2×32
// 27 = 19 + 2×22
// 33 = 31 + 2×12

// It turns out that the conjecture was false.

// What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
package challenge46

import (
	"euler/shared"
	"fmt"
)

const p46Max = 1_000_000

var c50Primes [p46Max]int
var dblSquares [p46Max]int
var oddComposites [p46Max]int

func Challenge46() {
	for n, i := 3, 0; i < p46Max; n += 2 {
		if !shared.IsPrime(n) {
			oddComposites[i] = n
			i++
		}
	}
	for n, i := 2, 0; i < p46Max; n++ {
		if shared.IsPrime(n) {
			c50Primes[i] = n
			i++
		}
	}

	for n := 0; n < p46Max; n++ {
		dblSquares[n] = 2 * n * n
	}

	for n1 := 0; n1 < p46Max; n1++ {
		successFlag := true
		//		fmt.Printf("Checking odd composite: %d\n", n1)
		for n2 := 0; n2 < p46Max; n2++ {
			diff := oddComposites[n1] - dblSquares[n2]
			if diff <= 2 {
				break
			}
			if shared.IsPrime(diff) {
				//				fmt.Printf("    It looks like %d is the sum of %d & %d\n", oddComposites[n1], dblSquares[n2], diff)
				successFlag = false
				break
			}
		}
		if successFlag {
			fmt.Printf("Challenge 46 solution: %d\n", oddComposites[n1])
			return
		}
		//		fmt.Printf("%d %d %d %d\n", n, primes[n], dblSquares[n], oddComposites[n])
	}
}
