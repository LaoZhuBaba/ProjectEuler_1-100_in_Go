// The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways:
//  (i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.

// There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there
// is one other 4-digit increasing sequence.

// What 12-digit number do you form by concatenating the three terms in this sequence?

package challenges

import (
	"fmt"
	"sort"
)

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

// Take a four digit integer and it as a string with the runes in order.  E.g., 3214 -> "1234"
func fourDigitsIntToOrderedString(n int) string {
	a := make([]int, 4)
	s := fmt.Sprintf("%d", n)
	for i, r := range s {
		a[i] = int(r)
	}
	sort.Ints(a)
	s = ""
	for _, i := range a {
		s += fmt.Sprintf("%c", i)
	}
	return s
}

func Challenge49() {
	m := make(map[string][]int)
	for n := 1000; n <= 9999; n++ {
		if isPrime(n) {
			// Use the string returned byt fourDigitsIntToOrderedString as a map key
			// containing all the prime numbers which have that string in common.
			// E.g., 1487 & 4817 would be appended to the same map value.
			nString := fourDigitsIntToOrderedString(n)
			m[nString] = append(m[nString], n)
		}
	}
	for _, v := range m {
		if len(v) >= 3 {
			var diff int
			// Step through every map key find the difference between each integer in the
			// map value slice.  If we find an integer that differs from two other integers
			// in the list by the same value then we have found a possible solution.
			for _, n1 := range v {
				diffMap := make(map[int]int)
				for _, n2 := range v {
					if n1 == n2 {
						continue
					}
					if n1 > n2 {
						diff = n1 - n2
					} else {
						if n1 < n2 {
							diff = n2 - n1
						}
					}
					// Each time we calculate diff we increment the value diffMap[diff]. If after incrementing the value is >=2
					// then n1 must lie equi-distant between two other values in the slice.  This is a possible solution.
					diffMap[diff]++
					if diffMap[diff] >= 2 {
						fmt.Printf("%d%d%d\n", n1-diff, n1, n1+diff)
					}
				}
			}
		}
	}
}
