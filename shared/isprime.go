package shared

import (
	"math"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	factors := make([]int, 0)
	var i int
	for i = 2; i < (int(math.Sqrt(float64(n))) + 1); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
			i--
		}
	}
	factors = append(factors, n)
	if len(factors) == 1 {
		return true
	} else {
		return false
	}
}

func GetPrimeFactors(n int) []int {
	if n < 2 {
		return []int{1}
	}
	factors := make([]int, 0)
	var i int
	for i = 2; i < (int(math.Sqrt(float64(n))) + 1); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
			i--
		}
	}
	return append(factors, n)
}
func GetPrimeFactorsNoDups(n int) []int {
	if n < 2 {
		return []int{1}
	}
	factors := make([]int, 0)
	factorMap := make(map[int]bool)
	var i int
	for i = 2; i < (int(math.Sqrt(float64(n))) + 1); i++ {
		if n%i == 0 {
			factorMap[i] = true
			n /= i
			i--
		}
	}
	factorMap[n] = true
	for k := range factorMap {
		factors = append(factors, k)
	}
	return factors
}
