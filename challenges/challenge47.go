// The first two consecutive numbers to have two distinct prime factors are:

// 14 = 2 × 7
// 15 = 3 × 5

// The first three consecutive numbers to have three distinct prime factors are:

// 644 = 2² × 7 × 23
// 645 = 3 × 5 × 43
// 646 = 2 × 17 × 19.

// Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?

package challenges

import (
	"fmt"
	"math"
)

const c47Max = 1000000
const consecTarget = 4

func returnFactors(n int) *[]int {
	factors := new([]int)
	var i int
	for i = 2; i < (int(math.Sqrt(float64(n))) + 1); i++ {
		if n%i == 0 {
			*factors = append(*factors, i)
			n /= i
			i--
		}
	}
	*factors = append(*factors, n)
	return factors
}
func returnFactorExponentMap(n int) *map[int]int {
	pFactors := returnFactors(n)
	m := make(map[int]int)
	for _, f := range *pFactors {
		m[f]++
	}
	return &m
}

func Challenge47() {
	consecutive := 0
	for c := 2; c < c47Max; c++ {
		pFactorExponentMap := returnFactorExponentMap(c)
		// fmt.Printf("%d %v\n", c, *pFactorExponentMap)
		if len(*pFactorExponentMap) == consecTarget {
			consecutive++
		} else {
			consecutive = 0
		}
		if consecutive == consecTarget {
			fmt.Printf("Solution is %d\n", c-3)
			return
		}
	}
}
