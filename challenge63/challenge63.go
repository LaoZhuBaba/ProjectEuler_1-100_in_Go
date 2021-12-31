package challenge63

import (
	"fmt"
	"math"
)

func Challenge63() {
	var solution int
	for n1 := float64(1); n1 <= 9; n1++ {
		for n2 := float64(1); n2 <= 30; n2++ {
			s := fmt.Sprintf("%0.f", math.Pow(n1, n2))
			l := len(s)
			if float64(l) == n2 {
				solution++
				fmt.Printf("%s %d\n", s, l)
			}

		}
	}
	fmt.Printf("Challenge 63 solution is: %d\n", solution)
}
