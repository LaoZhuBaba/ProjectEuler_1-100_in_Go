package challenge69

import (
	"euler/shared"
	"fmt"
)

const max = 1_000_000

// const max = 10

func Challenge69() {
	var solution int
	var maxRatio float64

	for n := 1; n <= max; n++ {
		factors := shared.GetPrimeFactors(n)
		factors = append(factors, n)
		totient := float64(n)
		var ratio float64
		if len(factors) == 2 {
			continue
		}
		for k, factor := range factors {
			if (k < len(factors)-2 && factor == factors[k+1]) || factor == n {
				continue
			} else {
				totient *= (1.0 - (1.0 / float64(factor)))
				ratio = float64(n) / totient
				if ratio > maxRatio {
					maxRatio = ratio
					solution = n
				}
			}
		}
		//fmt.Printf("%d %v %f %f\n", n, factors, totient, ratio)
	}
	fmt.Printf("Challenge69 solution: %d\n", solution)
}

// func sizeOfIntersection(s1, s2 []int) int {
// 	var count int
// 	for _, v1 := range s1 {
// 		for _, v2 := range s2 {
// 			if v1 == v2 {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
// func Challenge69() {
// 	m := make(map[int][]int)
// 	for count := 1; count <= max; count++ {
// 		var s []int
// 		shared.Factorise(count, &s)
// 		s = append(s, count)
// 		m[count] = append(m[count], s...)
// 		// fmt.Printf("%d: %v\n", count, s)
// 	}
// 	var solution int
// 	var maxRatio float64
// 	var totalIntersections int
// 	for n1 := 2; n1 <= max; n1++ {
// 		totalIntersections = 1
// 		for n2 := 1; n2 < n1; n2++ {
// 			if sizeOfIntersection(m[n1], m[n2]) == 1 {
// 				// fmt.Printf("Nothing for %d:%v & %d:%v\n", n1, m[n1], n2, m[n2])
// 				totalIntersections++
// 			}
// 		}
// 		// if totalIntersections > solution {
// 		// 	solution
// 		// }
// 		ratio := float64(n1) / float64(totalIntersections)
// 		// fmt.Printf("%d : %d (%f)\n", n1, totalIntersections, ratio)
// 		if ratio > maxRatio {
// 			maxRatio = ratio
// 			solution = n1
// 		}
// 	}
// 	fmt.Printf("Challenge 69 solution: %d\n", solution)
// }
