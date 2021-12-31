package challenge60

import (
	"fmt"
	"shared"
)

func generatePrimes(total int, pSlice *[]int) {
	for count, n := 0, 2; count < total; n++ {
		if shared.IsPrime(n) {
			*pSlice = append(*pSlice, n)
			count++
		}
	}
}

func areConcatsPrime(i1, i2 int) bool {
	var n1, n2 int
	fmt.Sscanf(fmt.Sprintf("%d%d", i1, i2), "%d", &n1)
	if !shared.IsPrime(n1) {
		return false
	}
	fmt.Sscanf(fmt.Sprintf("%d%d", i2, i1), "%d", &n2)
	return shared.IsPrime(n2)
}
func testCombinations(s []int, i int, f func(int, int) bool) bool {
	for _, v := range s {
		if !f(i, v) {
			return false
		}
	}
	return true
}

const numPrimes = 1300

func Challenge60() {

	s := make([]int, 0, numPrimes)
	generatePrimes(numPrimes, &s)
	for k1, v1 := range s[1:] {
		s2 := s[k1+1:]
		for k2, v2 := range s2 {
			if testCombinations([]int{v1}, v2, areConcatsPrime) {
				s3 := s2[k2+1:]
				for k3, v3 := range s3 {
					if testCombinations([]int{v1, v2}, v3, areConcatsPrime) {
						s4 := s3[k3+1:]
						for k4, v4 := range s4 {
							if testCombinations([]int{v1, v2, v3}, v4, areConcatsPrime) {
								s5 := s4[k4+1:]
								for _, v5 := range s5 {
									if testCombinations([]int{v1, v2, v3, v4}, v5, areConcatsPrime) {
										fmt.Printf("%d %d %d %d %d\n", v1, v2, v3, v4, v5)
										fmt.Printf("Challenge 60 solution: %d\n", v1+v2+v3+v4+v5)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
