// The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
// There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
// How many circular primes are there below one million?

package challenges

import "fmt"

func intSliceToInt(s []int) int {
	var ret int
	for _, v := range s {
		ret = ret*10 + v
	}
	return ret
}

func Challenge35() {
	var solution int
	for n := 2; n <= 1_000_000; n++ {
		nSlice := intToIntSlice(n)
		notPrimeFlag := false
		for count := len(nSlice); count > 0; count-- {
			n2 := intSliceToInt(nSlice)
			if !isPrime(n2) {
				notPrimeFlag = true
				break
			}
			rotateRight(nSlice)
		}
		if !notPrimeFlag {
			solution++
			fmt.Printf("%d\n", n)
		}
	}
	fmt.Printf("solution is: %d\n", solution)
}
