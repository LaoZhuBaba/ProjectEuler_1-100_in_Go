// By replacing the 1st digit of the 2-digit number *3, it turns out that six of the nine possible values: 13, 23, 43, 53, 73, and 83,
// are all prime.

// By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit number is the first example having seven primes among
// the ten generated numbers, yielding the family: 56003, 56113, 56333, 56443, 56663, 56773, and 56993. Consequently 56003, being the
// first member of this family, is the smallest prime with this property.

// Find the smallest prime which, by replacing part of the number (not necessarily adjacent digits) with the same digit, is part of an
// eight prime value family.

package challenges

import "fmt"

const c51NumerOfDigits = 8

func intToRuneSlice(n int) []rune {

	s := make([]rune, 0)
	for _, r := range fmt.Sprintf("%d", n) {
		s = append(s, r)
	}
	return s
}

func getDigitNumberRange(n int) (int, int) {
	ret := 1
	for count := 1; count < n; count++ {
		ret = ret * 10
	}
	return ret, (ret*10 - 1)
}

func maskSlice(s []rune, mask string) string {
	ret := ""
	var foundChar rune

	for count, eachChar := range mask {
		if eachChar == 'n' {
			if foundChar == 0 {
				foundChar = s[count]
				continue
			}
			if s[count] == foundChar {
				continue
			} else {
				return ""
			}
		}
	}
	for i := 0; i < len(s)-1; i++ {
		if mask[i] == 'y' {
			ret += string(s[i])
		} else {
			ret += "."
		}
	}
	return ret
}

func getPermutations(n int) []string {
	var ret = make([]string, 0)
	if n == 2 {
		return []string{"yy", "ny"}
	}
	for _, v := range getPermutations(n - 1) {
		ret = append(ret, "y"+v)
		ret = append(ret, "n"+v)
	}
	return ret
}
func Challenge51() {
	// fmt.Printf("%v\n", intToRuneSlice(1234))
	s, e := getDigitNumberRange(c51NumerOfDigits)
	var primeSlice = make([][]rune, 0)
	for count := s; count <= e; count++ {
		if isPrime(count) {
			primeSlice = append(primeSlice, intToRuneSlice(count))
		}
	}
	m := make(map[string][]rune)
	// We need to skip the first permutation because it is an empty mask (i.e., all y's)
	perms := getPermutations(c51NumerOfDigits)
	perms = perms[1:]
	for _, perm := range perms {
		fmt.Printf("perm is: %v\n", perm)
		for _, s := range primeSlice {
			key := maskSlice(s, perm)
			//fmt.Printf("key: %v\n", key)
			if key != "" {
				m[key] = append(m[key], s...)
			}
		}
		var maxLen int
		var maxKey string
		for k, v := range m {
			if len(v) > maxLen {
				maxLen = len(v)
				maxKey = k
			}
		}
		fmt.Printf("Longest: %c %v %d\n", m[maxKey], maxKey, maxLen)
	}

}
