package challenge27

import (
	"euler/shared"
	"fmt"
)

func quadratic(n, a, b int) int {
	return n*n + a*n + b
}
func intInList(i int, l []int) bool {
	for _, n := range l {
		if n == i {
			return true
		}
	}
	return false
}

// func absInt(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

func Challenge27() {
	c27Max := 0
	var solution int
	for a := 0; a <= 1000; a++ {
		//		if !isPrime(absInt(a)) {
		if !shared.IsPrime(a) {
			continue
		}
		for b := 0; b <= 1000; b++ {
			if !shared.IsPrime(b) {
				continue
			}
			list := make([]int, 0)
			for n := 0; n <= 1000; n++ {
				result := quadratic(n, -a, b)
				if result > 1 && shared.IsPrime(result) && !intInList(result, list) {
					list = append(list, result)
				} else {
					list = nil
					break
				}
				// fmt.Printf("For n=%d, a=%d, b=%d, x is: %d\n", n, a, b, result)
				if len(list) > c27Max {
					c27Max = len(list)
					solution = a * b
					fmt.Printf("max list is: %v when a is %d and b is %d with a product of %d and length %d\n", list, a, b, a*b, c27Max)
				}
			}
		}
	}
	fmt.Printf("Challenge 27 solution is: %d\n", -1*solution)
}
