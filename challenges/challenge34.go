// 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
// Find the sum of all numbers which are equal to the sum of the factorial of their digits.
// Note: As 1! = 1 and 2! = 2 are not sums they are not included.
package challenges

import "fmt"

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

func Challenge34() {
	var solution int
	// I chose 2,540,160 as maximum because it is equal to 9! * 9  This number is greater than than 999,999,999
	// so clearly any number higher than this must always be greater then the sum of digits factorialised.

	for n := 3; n <= 2_540_160; n++ {
		var sum int
		for _, digit := range intToIntSlice(n) {

			sum += factorialDigit(digit)
		}
		if sum == n {
			fmt.Printf("%d\n", n)
			solution += sum
		}
	}
	fmt.Printf("solution is %d\n", solution)

}
