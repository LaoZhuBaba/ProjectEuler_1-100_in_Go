package challenge58

import (
	"euler/shared"
	"fmt"
)

func Challenge58() {

	list := []int{1}
	var tmp, count, countPrimes, pc int
	for count = 3; pc != 9; count += 2 {
		cxc := count * count
		tmp = cxc - 3*(count-1)
		list = append(list, tmp)
		if shared.IsPrime(tmp) {
			countPrimes++
		}
		tmp = cxc - 2*(count-1)
		list = append(list, tmp)
		if shared.IsPrime(tmp) {
			countPrimes++
		}
		tmp = cxc - 1*(count-1)
		list = append(list, tmp)
		if shared.IsPrime(tmp) {
			countPrimes++
		}
		tmp = cxc
		list = append(list, tmp)
		if shared.IsPrime(tmp) {
			countPrimes++
		}
		pc = 100 * countPrimes / len(list)
	}
	fmt.Printf("Challenge 58 solution: %d\n", count)
}
