package challenge28

import "fmt"

func c28Solve(n uint64) uint64 {
	if n < 2 {
		return 1
	}
	return 4*((n*n)-(((n-1)/2)*3)) + c28Solve(n-2)
}

func Challenge28() {
	fmt.Printf("Challenge 28 solution is %d\n", c28Solve(1001))
}
