package shared

func Totient(i int) int {
	if i == 1 {
		return 1
	}
	factors := GetPrimeFactors(i)
	if len(factors) == 1 {
		return i - 1
	}
	totient := float64(i)
	for k, factor := range factors {
		if k > 0 && factors[k-1] == factors[k] {
			continue
		}
		totient *= (1.0 - (1.0 / float64(factor)))
	}
	return int(totient)
}

// for n := 1; n <= max; n++ {
// 	factors := shared.GetPrimeFactors(n)
// 	factors = append(factors, n)
// 	totient := float64(n)
// 	var ratio float64
// 	if len(factors) == 2 {
// 		continue
// 	}
// 	for k, factor := range factors {
// 		if (k < len(factors)-2 && factor == factors[k+1]) || factor == n {
// 			continue
// 		} else {
// 			totient *= (1.0 - (1.0 / float64(factor)))
// 			ratio = float64(n) / totient
// 			if ratio > maxRatio {
// 				maxRatio = ratio
// 				solution = n
// 			}
// 		}
// 	}
// 	//fmt.Printf("%d %v %f %f\n", n, factors, totient, ratio)
// }
