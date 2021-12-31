package shared

import "math"

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
