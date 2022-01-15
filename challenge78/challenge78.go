package challenge78

import (
	"euler/shared"
	"fmt"
)

const (
	// It's a bit dumb to use a fix sized array here but I reused code and
	// didn't bother to change to use a slice.  100_000 seems to be big enough.
	arraySize = 100_000
	target    = 1_000_000
)

func Challenge78() {
	// The first two values are just "given"
	array := [arraySize + 1]int{1, 1}

	for index := 2; index <= arraySize; index++ {
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
		array[index] %= target
		// The challenge is to find the lowest integer for which the total number of possible partitions
		// is divisible by 1 million.  This means that we only need the first 6 digits and don't need to
		// worry about of exceeding the possible size of a 64 bit int.  In this case the partition with
		// size of 1 is valid.
		if array[index]%target == 0 {
			fmt.Printf("Challenge 78 solution is: %d\n", index)
			return
		}
	}
}
