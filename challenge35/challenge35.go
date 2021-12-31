// The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
// There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
// How many circular primes are there below one million?

package challenge35

import (
	"euler/shared"
	"fmt"
)

func Challenge35() {
	var solution int
	for n := 2; n <= 1_000_000; n++ {
		nSlice := shared.IntToIntSlice(n)
		notPrimeFlag := false
		for count := len(nSlice); count > 0; count-- {
			n2 := shared.IntSliceToInt(nSlice)
			if !shared.IsPrime(n2) {
				notPrimeFlag = true
				break
			}
			shared.RotateRight(nSlice)
		}
		if !notPrimeFlag {
			solution++
			fmt.Printf("%d\n", n)
		}
	}
	fmt.Printf("Challenge 35 solution is: %d\n", solution)
}
