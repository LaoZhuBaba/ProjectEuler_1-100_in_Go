package challenge70

import (
	"euler/shared"
	"fmt"
	"sort"
)

const max = 10_000_000

func Challenge70() {
	m := make(map[int]string)
	var totientArray [max + 1]int
	var solution int
	var minDiv float64 = max + 1 // initialise to an impossibly high value
	for count := 2; count <= max; count++ {
		t := shared.Totient(count)
		totientArray[count] = t
		intSlice := shared.IntToIntSlice(count)
		sort.Ints(intSlice)
		var sortedTstr string
		for _, digit := range intSlice {
			sortedTstr += fmt.Sprintf("%d", digit)
		}
		m[count] = sortedTstr
	}
	for count := 1; count <= max; count++ {
		t := totientArray[count]
		if m[count] == m[t] {
			div := float64(count) / float64(t)
			if div < minDiv {
				minDiv = div
				solution = count
			}
		}
	}
	fmt.Printf("Challenge 70 solution is: %d\n", solution)
}
