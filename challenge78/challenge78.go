package challenge78

import (
	"euler/shared"
	"fmt"
)

const max = 100_000

func Challenge78() {
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
		array[index] %= 1_000_000
		if array[index]%1_000_000 == 0 {
			fmt.Printf("Challenge 78 solution is: %d\n", index)
			return
		}
	}
}
