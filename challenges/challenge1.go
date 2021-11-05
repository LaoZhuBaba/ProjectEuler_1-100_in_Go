// Find the sum of all the multiples of 3 or 5 below 1000.”
package challenges

import "fmt"

func Challenge1() {
	sum := 0
	for i := 1; i < 1000; i++ {
		// fmt.Printf("i is %d\n", i)
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("sum: %d\n", sum)
}
