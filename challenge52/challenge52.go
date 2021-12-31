// It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.

// Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.

package challenge52

import (
	"euler/shared"
	"fmt"
	"sort"
)

const c52Limit = 6

// func rotateIntRight(i int, numDigits int) int {
// 	rightmost := i % 10
// 	// fmt.Printf("rightmost is: %d\n", rightmost)
// 	for c := 1; c < numDigits; c++ {
// 		rightmost *= 10
// 	}
// 	// fmt.Printf("now rightmost is: %d\n", rightmost)
// 	shiftedRight := i / 10
// 	return rightmost + shiftedRight
// }
func intToKey(i int) string {
	sl := shared.IntToIntSlice(i)
	sort.Ints(sl)
	return fmt.Sprintf("%v", sl)
}
func Challenge52() {
	s, e := shared.GetDigitNumberRange(c52Limit)
	e /= 6

	for count := s; count <= e; count++ {
		key := intToKey(count)
		bingo := true
		for _, v := range []int{2, 3, 4, 5, 6} {
			if intToKey(count*v) != key {
				bingo = false
			}
		}
		if bingo {
			fmt.Printf("Challenge 52 solution: %d (%d %d %d %d %d)\n", count, count*2, count*3, count*4, count*4, count*6)
		}
	}

}
