package challenge92

import (
	"euler/shared"
	"fmt"
)

func numberChain(n int) (rn int) {
	sl := shared.IntToIntSlice(n)
	for index := range sl {
		sl[index] = sl[index] * sl[index]
	}
	sum := shared.SumOfList(&sl)
	// fmt.Printf("sum is: %d\n", sum)
	if sum == 89 || sum == 1 {
		return sum
	} else {
		return numberChain(sum)
	}

}
func Challenge92() {
	var solution int
	for n := 1; n < 10_000_000; n++ {
		if numberChain(n) == 89 {
			solution++
		}
	}
	fmt.Printf("challenge92 solution is: %d\n", solution)
}
