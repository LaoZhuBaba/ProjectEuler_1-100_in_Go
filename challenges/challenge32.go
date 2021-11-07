// We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once;
// for example, the 5-digit number, 15234, is 1 through 5 pandigital.
// The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1
// through 9 pandigital.
// Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
// HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
package challenges

import "fmt"

func containsDuplicates(s []int) bool {
	for index := 0; index < len(s)-1; index++ {
		n1 := s[index]
		//for index, n1 in range s {
		for _, n2 := range s[index+1:] {
			if n2 == n1 {
				return true
			}
		}
	}
	return false
}

// func intToIntSlice(n int) []int is defined in challenge30.go

func intContainsDuplicateDigits(n int) bool {
	return containsDuplicates(intToIntSlice(n))
}
func intContainsDigit0(n int) bool {
	nSlice := intToIntSlice(n)
	for _, digit := range nSlice {
		if digit == 0 {
			return true
		}
	}
	return false
}
func intLen(n int) int {
	return len(intToIntSlice(n))
}
func Challenge32() {
	m := make(map[int]bool)
	for a := 1; a < 9999; a++ {
		if intContainsDigit0(a) {
			continue
		}
		if intContainsDuplicateDigits(a) {
			continue
		}
		aLen := intLen(a)
		for b := 1; b < 9999; b++ {
			bLen := intLen(b)
			product := a * b
			if intLen(product) != 9-aLen-bLen {
				continue
			}
			// From here we know that the combined length of the product and factors is 9
			if intContainsDigit0(b) {
				continue
			}
			if intContainsDuplicateDigits(b) {
				continue
			}
			ab := append(intToIntSlice(a), intToIntSlice(b)...)
			if containsDuplicates(ab) {
				continue
			}
			if intContainsDigit0(product) {
				continue
			}
			if intContainsDuplicateDigits(product) {
				continue
			}
			if containsDuplicates(append(intToIntSlice(product), ab...)) {
				continue
			}
			// m[a] = true
			// m[b] = true
			m[product] = true
			fmt.Printf("%d x %d = %d\n", a, b, product)
		}
	}
	fmt.Printf("%v\n", m)
	var solution int
	for k := range m {
		solution += k
	}
	fmt.Printf("Solution is: %d\n", solution)
}
