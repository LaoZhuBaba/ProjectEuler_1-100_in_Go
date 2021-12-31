//What is the 10,001 prime number?
package challenge7

import (
	"euler/shared"
	"fmt"
)

func Challenge7() {
	count := 0
	for n := 2; true; n += 1 {
		if shared.IsPrime(n) {
			count += 1
			if count == 10001 {
				fmt.Printf("Challenge 7 solution is: %d\n", n)
				return
			}
		}
	}
}
