package challenge76

import (
	"euler/shared"
	"fmt"
)

// The following code works and generates all possible combinations for up to about 30 but
// after that the calculation time became ridiculously long.  So unsuitable to use, but
// but still quite cool.
//
// func allAddends(i int) [][]int {
// 	s := make([][]int, 0)
// 	if i == 2 {
// 		return [][]int{{1, 1}}
// 	}
// 	for x := 1; x <= i/2; x++ {
// 		y := i - x
// 		s = append(s, []int{x, y})
// 		for _, v := range allAddends(y) {
// 			suffix := append([]int{x}, v...)
// 			if !sort.IntsAreSorted(suffix) {
// 				continue
// 			}
// 			s = append(s, suffix)
// 		}
// 	}
// 	return s
// }

const max = 100

func Challenge76() {
	// The first two values are just "given"
	array := [max + 1]int{1, 1}

	for index := 2; index <= max; index++ {
		generator := shared.PentGen()
		var subtractor, toggle int
		// index is the index into array[] which stores the results for each value up to max
		// subtractor is used to look up values from lower index values.  I.e., later values are
		// computed based on earlier values.  The size of the subtraction is based on the series
		// of pentangular numbers which we get from generator().  We need toggle because values
		// we look up via subtractor are sometimes added and sometimes subtracted.  The pattern is
		// add, add, subtract, subtract, add, add, ...  This all comes from Euler's Pentagonal
		// Theorem.
		for {
			subtractor = generator()
			if index-subtractor < 0 {
				break
			}
			if toggle%4 < 2 {
				array[index] += array[index-subtractor]
			} else {
				array[index] -= array[index-subtractor]
			}
			toggle++
		}
	}
	fmt.Printf("%d\n", array[max]-1)
}
