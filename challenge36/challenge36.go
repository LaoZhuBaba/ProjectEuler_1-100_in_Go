// The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.
// Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
// (Please note that the palindromic number, in either base, may not include leading zeros.)

package challenge36

import "fmt"

func isStrPalindrome(s string) bool {
	length := len(s)
	half := length / 2
	for c := 0; c < half; c++ {
		if s[c] != s[length-c-1] {
			return false
		}
	}
	return true
}
func Challenge36() {
	var solution int
	for c := 1; c <= 1_000_000; c++ {
		decStr := fmt.Sprintf("%d", c)
		binStr := fmt.Sprintf("%b", c)
		if isStrPalindrome(decStr) && isStrPalindrome(binStr) {
			fmt.Printf("%s %s\n", decStr, binStr)
			solution += c
		}
	}
	fmt.Printf("Challenge 36 solution is: %d\n", solution)
}
