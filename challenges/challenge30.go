package challenges

import (
	"fmt"
)

// func intToIntSlice(n int) []int {
// 	nStr := fmt.Sprintf("%d", n)
// 	slice := make([]int, 0)
// 	var i int
// 	for _, r := range nStr {
// 		fmt.Sscanf(string(r), "%d", &i)
// 		slice = append(slice, i)
// 	}
// 	return slice
// }
func intToIntSlice(n int) []int {
	slice := make([]int, 0)

	for n != 0 {
		slice = append([]int{n % 10}, slice...)
		n = n / 10
	}
	return slice
}

// func intSliceToInt(s []int) int {
// 	var ret int
// 	for _, v := range s {
// 		ret = ret*10 + v
// 	}
// 	return ret
// }

func Challenge30() {
	var grandTotal int
	for c := 1; c < 1_000_000; c++ {
		tot := 0
		for _, digit := range intToIntSlice(c) {
			tot = tot + (digit * digit * digit * digit * digit)
		}
		if tot == c && tot != 1 {
			fmt.Printf("%d\n", c)
			grandTotal += tot
		}
	}
	fmt.Printf("Solution is: %d\n", grandTotal)
}
