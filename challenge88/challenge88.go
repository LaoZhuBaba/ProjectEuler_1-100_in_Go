package challenge88

import (
	"euler/shared"
	"fmt"
	"math"
)

const TARGET = 12_000

// Return a slice containing all the factors of n, without duplicates
func GetFactorsNoDups(n int) []int {
	if n < 2 {
		return []int{1}
	}
	factors := make([]int, 0)
	var i int
	for i = 2; i < (int(math.Sqrt(float64(n))) + 1); i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// Takes an integer "n" and returns a slice of slices.  Each inner slice is a set of factors which multiply to
// equal "n".  Duplicate sets are eliminated.  For example when called with n=8, the slice of slices returned
// is [[2 4] [2 2 2]].  When called with n=12 the slice returned is [[2 6] [2 2 3] [3 4]].  If passed a prime
// number the function returns any empty slice.
func intToFactorSets(n int) [][]int {
	factors := GetFactorsNoDups(n)

	ret := make([][]int, 0)
	// m is used to eliminate duplicates.  Before saving a factor set we add all the contained values
	// together and multiply by the square of the length of the set.  This value can be used to check
	// if an equivalent factor set has already been saved.  I'm fairly sure that it is impossible for
	// two factor sets which are not equal to have generate the same hash value according to this logic.
	// E.g., [4 4] has the same sum and product as [4 2 2] but different length will distinguish them.
	// But [3 2] and [2 3] have the same sum, product and length so are considered indentical.  I
	// found that I needed to use the square of the length rather than just the length.  Otherwise
	// some non-identical sets ended up with the same hash check number.
	m := make(map[int]bool)
	if len(factors) == n {
		return ret
	}
	// iterate through the unique factors of n
	for _, f := range factors {
		factor2 := n / f
		factorPair := []int{f, factor2}
		hash := (f + factor2) * 2 * 2 * 2
		// Check m to see if this set has been seen before and append to ret if it hasn't.
		if !m[hash] {
			ret = append(ret, factorPair)
			m[hash] = true
		}
		// If factor2 is prime then there is nothing more to do on this iteration.
		if shared.IsPrime(factor2) {
			continue
		}
		// If factor2 is not prime then make a recursive call to the current function
		// to get a unique list of factors for factor2
		for _, additional := range intToFactorSets(factor2) {
			// Create candidate sets by iteratively appending the returned sets of factors to f.
			// This set will obviously have a length greater than 2 so to generate the hash check
			// value we need to iterate through create a sum which is multiplied by the length
			// of candidate as desribed above.  Finally, check the hash check against m to see
			// if the resulting set has been seen before.
			candidate := append([]int{f}, additional...)
			var sum int
			for _, cf := range candidate {
				sum += cf
			}
			sum = sum * len(candidate) * len(candidate) * len(candidate)
			// If this set hasn't been seen before then append it to ret
			if !m[sum] {
				ret = append(ret, candidate)
				m[sum] = true
			}
		}
	}
	return ret
}

func Challenge88() {
	// Map keyed on total length, containing the smallest Product/Sum value for a set of that length
	tlMap := make(map[int]int)
	// Map keyed on the Product/Sum value.  The only purpose of this map is to eliminate duplicates.
	productMap := make(map[int]bool)

	var totalLength int

	// The maximum we are targeting is the size of the Product/Sum set but we are iterating through product.
	// because the product of the maximum Product Sum set may be significantly greater then its length and
	// it is hard to predict how much higher it will be, we will iterate a total length which is double
	// the target.  By that point we can be pretty confident that we have covered everything.  In fact
	// experimentation tells me that we only need to a bit higher than TARGET, but that would look confusing.
	for product := 2; totalLength <= TARGET*2; product++ {
		factorSetList := intToFactorSets(product)
		if len(factorSetList) == 0 {
			continue
		}
		for _, factorSet := range factorSetList {
			sum := shared.SumOfList(&factorSet)
			numberOfOnes := product - sum
			totalLength = len(factorSet) + numberOfOnes
			if tlMap[totalLength] == 0 {
				fmt.Printf("%d: %d x 1 + %v (%d)\n", totalLength, numberOfOnes, factorSet, product)
				tlMap[totalLength] = product
				if totalLength <= TARGET {
					productMap[product] = true
				}
			}
		}
	}
	fmt.Printf("%v\n", productMap)
	var solution int
	for k := range productMap {
		solution += k
	}
	fmt.Printf("Challenge 88 solution is: %d\n", solution)
}
