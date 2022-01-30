package challenge87

import (
	"euler/shared"
	"fmt"
)

const (
	target   = 50_000_000
	max2root = 7071 // the largest number that squares to < 50,000,000
	max3root = 368  // the largest number that cubes to < 50,000,000
	max4root = 85   // the largest number with 4th power < 50,000,000
)

func Challenge87() {
	primes := make([]int, 0)
	primeSquares := make([]int, 0)
	primeCubes := make([]int, 0)
	prime4thPowers := make([]int, 0)
	m := make(map[int]bool)

	for count := 2; count < max2root; count++ {
		if shared.IsPrime(count) {
			primes = append(primes, count)
		}
	}
	for _, prime := range primes {
		primeSquares = append(primeSquares, prime*prime)
		if prime <= max3root {
			primeCubes = append(primeCubes, prime*prime*prime)
		}
		if prime <= max4root {
			prime4thPowers = append(prime4thPowers, prime*prime*prime*prime)
		}
	}
	for _, primeSquare := range primeSquares {
		for _, primeCube := range primeCubes {
			if primeSquare+primeCube >= target {
				break
			}
			for _, prime4thPower := range prime4thPowers {
				sum := primeSquare + primeCube + prime4thPower
				if sum < target {
					m[sum] = true
				} else {
					break
				}
			}
		}
	}
	// fmt.Printf("%v\n", primeSquares)
	// fmt.Printf("%v\n", primeCubes)
	// fmt.Printf("%v\n", prime4thPowers)
	fmt.Printf("Challenge 87 solution is: %d\n", len(m))
}
