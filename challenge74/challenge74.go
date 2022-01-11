package challenge74

import (
	"euler/shared"
	"fmt"
)

func factorialDigit(n int) int {
	a := [10]int{
		1,
		1,
		2,
		3 * 2,
		4 * 3 * 2,
		5 * 4 * 3 * 2,
		6 * 5 * 4 * 3 * 2,
		7 * 6 * 5 * 4 * 3 * 2,
		8 * 7 * 6 * 5 * 4 * 3 * 2,
		9 * 8 * 7 * 6 * 5 * 4 * 3 * 2,
	}
	return a[n]
}

func measureDFChain(n int) int {
	m := make(map[int]bool)
	m[n] = true
	count := 1
	current := digitFactorialChain(n)
	for {
		if _, ok := m[current]; ok {
			break
		}
		m[current] = true
		count++
		current = digitFactorialChain(current)
	}
	return count
}
func digitFactorialChain(n int) (ret int) {
	for _, i := range shared.IntToIntSlice(n) {
		ret += factorialDigit(i)
	}
	return
}

func Challenge74() {
	var length, solution int
	for count := 1; count < 1_000_000; count++ {
		length = measureDFChain(count)
		if length == 60 {
			solution++
			fmt.Printf("%d has a digit factorial chain length of 60\n", count)
		}
	}
	fmt.Printf("Challenge 74 solution is: %d\n", solution)
}
