package challenge14

import "fmt"

func collatz(n int) int {
	count := 1
	for {
		if n == 1 {
			return count
		}
		if (n % 2) == 0 {
			n /= 2
			count++
			continue
		} else {
			n = n*3 + 1
			count++
		}
	}
}

func Challenge14() {

	var maxNum, maxSeq, seq int

	for n := 1; n <= 1_000_000; n++ {
		seq = collatz(n)
		if seq > maxSeq {
			maxSeq = seq
			maxNum = n
		}
	}
	fmt.Printf("maxNum is: %d; maxSeq is %d\n", maxNum, maxSeq)
	fmt.Printf("Challenge 14 solution is: %d\n", maxNum)
}
