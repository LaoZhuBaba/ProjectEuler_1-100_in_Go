// The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove
// digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work
// from right to left: 3797, 379, 37, and 3.

// Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

// NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

package challenges

import "fmt"

func Challenge37() {

	var solution int
	for n := 11; n <= 1_000_000; n++ {
		if !isPrime(n) {
			continue
		}
		nStr := intToIntSlice(n)
		length := len(nStr)
		notPrimeFlag := false
		for i := length - 1; i > 0; i-- {
			if !isPrime(intSliceToInt(nStr[0:i])) {
				// fmt.Printf("%d\n", intSliceToInt(nStr[0:i]))
				notPrimeFlag = true
				break
			} else {
				if !isPrime(intSliceToInt(nStr[length-i:])) {
					// fmt.Printf("%d\n", intSliceToInt(nStr[length-i:]))
					notPrimeFlag = true
					break
				}
			}
			// fmt.Printf("Truncate from right: %d\n", intSliceToInt(nStr[0:i]))
			// fmt.Printf("Truncate from left: %d\n", intSliceToInt(nStr[length-i:]))
		}
		if !notPrimeFlag {
			solution += n
		}
	}
	fmt.Printf("solution is %d\n", solution)

}
