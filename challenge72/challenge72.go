package challenge72

import (
	"euler/shared"
	"fmt"
)

// const max = 1_000_000
const max = 3

//const max = 10

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

func Challenge72() {

	var solution int

	for n := 2; n <= max; n++ {
		solution += shared.Totient(n)
	}
	// // primes := make(map[int]bool)
	// primes := make([]bool, max+1)
	// for n := 2; n <= max; n++ {
	// 	primes[n] = shared.IsPrime(n)
	// }
	// for denom := 2; denom <= max; denom++ {
	// 	// If denom is prime count one reduced proper fraction for every integer from 1 to denom - 1
	// 	if primes[denom] {
	// 		// fmt.Printf("denom of %d is prime, so adding %d\n", denom, denom-1)
	// 		solution += denom - 1
	// 		// fmt.Printf("solution is now: %d\n", solution)
	// 		continue
	// 	}
	// 	// 1/denom will always be a reduced proper faction, so add for a start
	// 	solution++
	// 	// fmt.Printf("For assumed number of 1 incremented solution to %d\n", solution)
	// 	for numer := 2; numer < denom; numer++ {
	// 		// denomFactors := shared.GetPrimeFactors(denom)
	// 		// numFactors := shared.GetPrimeFactors(numer)
	// 		// fmt.Printf("Hit here with solution: %d, numer: %d & denom: %d\n", solution, numer, denom)
	// 		if denom%numer != 0 && primes[numer] {
	// 			// solution++
	// 			// fmt.Printf("numer is prime (%d) so incremented solution to %d\n", numer, solution)
	// 			for numerMultiplier := numer; numerMultiplier < denom; numerMultiplier *= numer {
	// 				if denom%numer != 0 && primes[numer] {
	// 					solution++
	// 				}
	// 			}
	// 			// fmt.Printf("%v %v\n", numFactors, denomFactors)
	// 			// if sizeOfIntersection(denomFactors, numFactors) == 0 && isInIntMap(numer, primes) {
	// 			// if isInIntMap(numer, primes) {
	// 			// 	solution++
	// 			// 	// fmt.Printf("factors don't intersect so incremented solution to %d\n", solution)
	// 			// }
	// 		}
	// 	}
	// }
	fmt.Printf("Challenge 72 solution is: %d\n", solution)
}
