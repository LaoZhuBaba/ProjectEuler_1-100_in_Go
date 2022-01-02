package challenge64

import (
	"euler/shared"
	"fmt"
)

func Challenge64() {
	var solution int
	var si []int
	var cycle int
	var longestCycle int
	for count := 1; count <= 10_000; count++ {
		si, cycle = shared.Cfsr(count)
		// fmt.Printf("%d: %v -- %d\n", count, si, cycle)
		if cycle%2 != 0 {
			solution++
		}
		if cycle > longestCycle {
			longestCycle = cycle
			fmt.Printf("longestCycle so far is: %d\n", longestCycle)
			fmt.Printf("%d: %v\n", count, si)
		}
	}
	fmt.Printf("solution is: %d\n", solution)
}
