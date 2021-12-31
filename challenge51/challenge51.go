// By replacing the 1st digit of the 2-digit number *3, it turns out that six of the nine possible values: 13, 23, 43, 53, 73, and 83,
// are all prime.

// By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit number is the first example having seven primes among
// the ten generated numbers, yielding the family: 56003, 56113, 56333, 56443, 56663, 56773, and 56993. Consequently 56003, being the
// first member of this family, is the smallest prime with this property.

// Find the smallest prime which, by replacing part of the number (not necessarily adjacent digits) with the same digit, is part of an
// eight prime value family.

package challenge51

import (
	"euler/shared"
	"fmt"
)

const c51NumerOfDigits = 6

func replaceDigit(s []int, n, replace int) (bool, []int) {
	ret := make([]int, len(s))
	tf := false
	for k, v := range s {
		if v == n {
			ret[k] = replace
			tf = true
		} else {
			ret[k] = v
		}
	}
	return tf, ret
}

func Challenge51() {
	s, e := shared.GetDigitNumberRange(c51NumerOfDigits)
	for count := s; count <= e; count++ {
		if !shared.IsPrime(count) {
			continue
		}
		primeList := shared.IntToIntSlice(count)
		var replacementList []int
		var changed bool
		for _, digit := range [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
			primeCount := 1 // Start at one because the original number is a prime
			notPrimeCount := 0
			for _, replacementDigit := range [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
				if digit != replacementDigit {
					changed, replacementList = replaceDigit(primeList, digit, replacementDigit)
					if changed {
						i := shared.IntSliceToInt(replacementList)
						if shared.IsPrime(i) && i > s {
							primeCount++
						} else {
							notPrimeCount++
						}
					}
				}
				if notPrimeCount > 2 {
					break
				}
				if primeCount >= 8 {
					fmt.Printf("Challenge 51 solution is %d\n", shared.IntSliceToInt(primeList))
					return
				}

			}
		}
	}
}
