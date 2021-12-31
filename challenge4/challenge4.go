// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit number
// is 9009 = 91 × 99. Find the largest palindrome made from the product of two 3-digit numbers.
package challenge4

import "fmt"

func Challenge4() {
	var prod int
	palindromeList := make([]int, 0)
	for n1 := 999; n1 > 500; n1 -= 1 {
		for n2 := 999; n2 > 500; n2 -= 1 {
			// fmt.Printf("n1 is: %d and n2 is %d\n", n1, n2)
			prod = n1 * n2
			// fmt.Printf("prod is: %d\n", prod)
			if isPalindrome(prod) {
				palindromeList = append(palindromeList, prod)
			}
		}
	}
	max := 0
	for _, i := range palindromeList {
		if i > max {
			max = i
		}
	}
	fmt.Printf("Challenge 4 solution is: %d\n", max)
}

// func Challenge4() {
// 	// var prod int
// 	// palindromeList := make([]int, 0)
// 	for n := 999999; n > 0; n -= 1 {
// 		if isPalindrome(n) {
// 			sr := int(math.Sqrt(float64(n)))
// 			//fmt.Printf("square root of palindrome is: %d\n", sr)
// 			for div1 := sr; div1 > 0; div1 -= 1 {
// 				if n/div1 < 1000 {
// 					if n%div1 == 0 {

// 						fmt.Printf("palindrome %d has factors: %d & %d\n", n, div1, n/div1)
// 						return
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// for n1 := 999; n1 > 500; n1 -= 1 {
// 	for n2 := 999; n2 > 500; n2 -= 1 {
// 		// fmt.Printf("n1 is: %d and n2 is %d\n", n1, n2)
// 		prod = n1 * n2
// 		// fmt.Printf("prod is: %d\n", prod)
// 		if isPalindrome(prod) {
// 			palindromeList = append(palindromeList, prod)
// 		}
// 	}
// }
// max := 0
// for _, i := range palindromeList {
// 	if i > max {
// 		max = i
// 	}
// }
// fmt.Printf("Highest palindrome is: %d\n", max)

func reverseDecimal(n int) int {
	// fmt.Printf("reverseDecimal called with %d\n", n)
	r := 0
	for n != 0 {
		r = r*10 + n%10
		n /= 10
	}
	// fmt.Printf("reverseDecimal is returning: %d\n", r)
	return r
}

func isPalindrome(n int) bool {
	var tf bool
	// fmt.Printf("isPalindrome called with: %d\n", n)
	n2 := reverseDecimal(n)
	tf = (n == n2)
	// fmt.Printf("tf is: %v\n", tf)
	return tf
}
