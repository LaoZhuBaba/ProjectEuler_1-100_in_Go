package challenge93

import (
	"euler/shared"
	"fmt"
	"math"
	"sort"
)

const (
	ADD = iota + 1000
	SUBTRACT
	DIVIDE
	MULTIPLY
)

// maxConseqFloat64 takes a slice of floats, orders them and then returns the
// maximum value which can be counted to starting at 1.  For example, [1 2 5 6 7]
// would return 2
func maxConseqFloat64(floats []float64) (max float64) {
	sort.Float64s(floats)
	for k, v := range floats {
		if float64(k) == v-1 {
			max = v
		} else {
			return max
		}
	}
	return max
}

// rp takes a list of integers, which are either digits in the range 1 - 9 or
// operators in the range 1000 - 1003 (ADD, SUBTRACT, DIVIDE, MULTIPLY, above).
// These digits and operators are then interpreted as a Reverse Polish expression.
// and a float64 result returned.  If the calculated value is not a whole number
// then -1 is returned.
func rp(r []int) float64 {
	var stack []float64
	var tmp float64
	for _, v := range r {
		vf := float64(v)
		// fmt.Printf("v is: %d\n", v)
		switch {
		case v <= 9 && v >= 0:
			stack = append([]float64{vf}, stack...)
		case v == ADD:
			// fmt.Printf("adding %f & %f\n", stack[0], stack[1])
			tmp = stack[0] + stack[1]
			stack = stack[1:]
			stack[0] = tmp
		case v == MULTIPLY:
			// fmt.Printf("multiplying %f & %f\n", stack[0], stack[1])
			tmp = stack[0] * stack[1]
			stack = stack[1:]
			stack[0] = tmp
		case v == DIVIDE:
			if stack[0] == 0 {
				return -1
			}
			// fmt.Printf("DIVIDING %f by %f\n", stack[1], stack[0])
			tmp = stack[1] / stack[0]
			stack = stack[1:]
			stack[0] = tmp
		case v == SUBTRACT:
			// fmt.Printf("subtracting %f from %f\n", stack[0], stack[1])
			tmp = stack[1] - stack[0]
			stack = stack[1:]
			stack[0] = tmp
		}
		// fmt.Printf("tmp: %f\n", tmp)
	}
	if math.Round(stack[0]) != stack[0] {
		return -1
	}
	return stack[0]
}

func ParamToString(param []int) (ret string) {
	for _, p := range param {
		switch {
		case p <= 9:
			ret += fmt.Sprintf("%d", p)
		case p == ADD:
			ret += "+"
		case p == SUBTRACT:
			ret += "-"
		case p == MULTIPLY:
			ret += "*"
		case p == DIVIDE:
			ret += "/"
		}
	}
	return ret
}

// rpGroupings takes a slice of four numbers and three operators evaluate these as a
// Reverse Polish expression with five possible groupings of numbers and operators
// and return a slice of the five possible results.
func rpGroupings(numbers []int, operators []int) (ret []float64) {
	var params []int
	var floatResult float64

	// num, num, num, num, oper, oper, oper
	params = append(numbers, operators...)
	floatResult = rp(params)
	if floatResult > 0 && math.Round(floatResult) == floatResult {
		ret = append(ret, rp(params))
	}
	// num, num, num, oper, oper, num,oper
	params = []int{numbers[0], numbers[1], numbers[2], operators[0], operators[1], numbers[3], operators[2]}
	floatResult = rp(params)
	if floatResult > 0 && math.Round(floatResult) == floatResult {
		ret = append(ret, rp(params))
	}
	// num, num, num, oper, num, oper, oper
	params = []int{numbers[0], numbers[1], numbers[2], operators[0], numbers[3], operators[1], operators[2]}
	floatResult = rp(params)
	if floatResult > 0 && math.Round(floatResult) == floatResult {
		ret = append(ret, rp(params))
	}
	// num, num, oper, num, num, oper, oper
	params = []int{numbers[0], numbers[1], operators[0], numbers[2], numbers[3], operators[1], operators[2]}
	floatResult = rp(params)
	if floatResult > 0 && math.Round(floatResult) == floatResult {
		ret = append(ret, rp(params))
	}
	// num, num, oper, num, oper, num, oper
	params = []int{numbers[0], numbers[1], operators[0], numbers[2], operators[1], numbers[3], operators[2]}
	floatResult = rp(params)
	if floatResult > 0 && math.Round(floatResult) == floatResult {
		ret = append(ret, rp(params))
	}
	return ret
}

func Challenge93() {
	// maxConseq will be used to judge the solution
	var maxConseq float64
	var maxConseqKey string

	// calculate all possible permutations with repetition of a three
	// member subset of the integers 0 to 3.
	operatorIndexes := shared.PermutationsWithRepetition(3, 4)
	// operatorIndexes be based on 0 - 3, but we want the same pattern
	// based on 1000 - 1003.  So just add 1000 to every value.  This now
	// represents every possible 3 element permutation with repetition of
	// the set {ADD, SUBTRACT, MULTIPLY, DIVIDE}
	operatorCombinations := make([][]int, len(operatorIndexes))
	for index, combination := range operatorIndexes {
		for _, v := range combination {
			operatorCombinations[index] = append(operatorCombinations[index], v+1000)
		}
	}

	// Similar the above but this time for digits which cannot repeat.
	// So four element subsets of the set {1,2,3,4,5,6,7,8,9} without repetition.
	digitIndexes := shared.NMemberSubsets(4, 9)
	digitCombinations := make([][]int, len(digitIndexes))
	// Transform from 0-8 to 1-9 by adding 1 to every element.
	for index, combination := range digitIndexes {
		for _, v := range combination {
			digitCombinations[index] = append(digitCombinations[index], v+1)
		}
	}

	resultMap := make(map[string]map[float64]bool)

	// Combine every permutation of digits with every permutation of operators
	for _, dc := range digitCombinations {
		for _, oc := range operatorCombinations {
			// For the purposes of our calculation, the order of digits doesn't
			// mater sort the digits and then convert to a string which we can
			// use as our map key.  This allows us to easily append the results
			// of [1,2,3,4] & [1,3,4,2] to the same map key value.
			var sorted []int
			sorted = append(sorted, dc...)
			sort.Ints(sorted)
			sortedStr := fmt.Sprintf("%v", sorted)
			// totals is a list of all possible totals
			totals := rpGroupings(dc, oc)
			// Make a new map key if it doesn't exist
			if _, ok := resultMap[sortedStr]; !ok {
				resultMap[sortedStr] = make(map[float64]bool)
			}
			// We don't care about the values, only the map keys
			for _, total := range totals {
				resultMap[sortedStr][total] = true
			}
		}
	}
	// We now have all the results but we don't know which map key holds
	// the winning sequence.
	for key, floats := range resultMap {
		var resultSlice []float64
		// Convert the float64 key values to a slice
		for f := range floats {
			resultSlice = append(resultSlice, f)
		}
		// Calculate the number of consecutive counting numbers in the result set
		conseq := maxConseqFloat64(resultSlice)
		// Update maxConseq each time a better result is found
		if conseq > maxConseq {
			maxConseq = conseq
			maxConseqKey = key
		}
	}
	fmt.Printf("Challenge 93 solution: %s, with %.0f consecutive counting numbers\n", maxConseqKey, maxConseq)
}
