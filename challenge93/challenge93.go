package challenge93

import (
	"fmt"
	"math"
	"sort"
)

func setsWithinSet(n int) (sets [][]int) {
	sets = [][]int{{n}}
	if n == 1 {
		return [][]int{{1}}
	}

	for head := n - 1; head > 0; head-- {
		tail := n - head
		headSet := [][]int{}
		for _, tailSet := range setsWithinSet(tail) {
			headSet = append(headSet, append([]int{head}, tailSet...))
		}
		sets = append(sets, headSet...)
	}
	return sets
}

func nMemberSubsets(subsetSize, setSize int) (subsets [][]int) {

	set := make([]int, setSize)
	for index := range set {
		set[index] = index
	}
	return nMemberSubsets2(subsetSize, set)
}

func nMemberSubsets2(subsetSize int, set []int) (subsets [][]int) {
	if subsetSize == 1 {
		for _, memb := range set {
			subsets = append(subsets, []int{memb})
		}
	}
	for index, member := range set {
		var left, right []int
		left = append(left, set[:index]...)
		right = append(right, set[index+1:]...)
		oneLessSet := append(left, right...)
		headSet := [][]int{}
		for _, tailSet := range nMemberSubsets2(subsetSize-1, oneLessSet) {
			headSet = append(headSet, append([]int{member}, tailSet...))
		}
		subsets = append(subsets, headSet...)
	}
	return subsets
}

const (
	ADD = iota
	SUBTRACT
	DIVIDE
	MULTIPLY
)

func Arithmetic(parameters []int, operators []int) (result float64) {
	result = float64(parameters[0])
	for index, operator := range operators {
		switch operator {
		case ADD:
			result += float64(parameters[index+1])

		case SUBTRACT:
			result -= float64(parameters[index+1])

		case MULTIPLY:
			result *= float64(parameters[index+1])

		case DIVIDE:
			result /= float64(parameters[index+1])
		}
	}
	return result
}

func Challenge93() {
	resultMap := make(map[string]map[float64]bool)
	digitCombinations := nMemberSubsets(4, 9)
	// fmt.Printf("digitCombinations: %v (length: %d)\n", digitCombinations, len(digitCombinations))
	// ascending := [][]int{}
	// for _, v := range digitCombinations {
	// 	if v[0] < v[1] && v[1] < v[2] && v[2] < v[3] {
	// 		v[0]++
	// 		v[1]++
	// 		v[2]++
	// 		v[3]++
	// 		ascending = append(ascending, v)
	// 	}
	// }
	// fmt.Printf("ascending: %v (length: %d)\n", ascending, len(ascending))

	result := AllOrderedCombinations(4)
	fmt.Printf("result: %v (length: %d)\n", result, len(result))

	for _, r := range result {
		for _, c := range digitCombinations {
			var sorted []int
			var cc []int
			cc = append(cc, c...)
			cc[0]++
			cc[1]++
			cc[2]++
			cc[3]++
			sorted = append(sorted, cc...)
			sort.Ints(sorted)
			sortedStr := fmt.Sprintf("%v", sorted)
			total := Arithmetic(cc, r)
			if total < 1 {
				continue
			}
			if math.Round(total) != total {
				continue
			}
			if _, ok := resultMap[sortedStr]; !ok {
				resultMap[sortedStr] = make(map[float64]bool)
			}
			if sortedStr == "[1 2 3 4]" {
				fmt.Printf("cc is: %v with total %f and r: %d\n", cc, total, r)
			}
			resultMap[sortedStr][total] = true
		}
		// if math.Round(total) != total {
		// 	continue
		// }
		//fmt.Printf("%f (%v)\n", total, r)
		fmt.Printf("resultsMap['[1 2 3 4]']: %v\n", resultMap["[1 2 3 4]"])
	}
	// fmt.Printf("1 + 2 / 3 * 4 = %f\n", Arithmetic([]int{1, 2, 3, 4}, []int{ADD, DIVIDE, MULTIPLY}))
	// fmt.Printf("5 + 2 / 4 - 3 = %f\n", Arithmetic([]int{5, 2, 4, 3}, []int{ADD, DIVIDE, SUBTRACT}))
	// fmt.Printf("1 + 3 / 2 * 4 = %f\n", Arithmetic([]int{1, 3, 2, 4}, []int{ADD, DIVIDE, MULTIPLY}))
	// fmt.Printf("1 + 4 / 2 - 6 = %f\n", Arithmetic([]int{1, 4, 2, 6}, []int{ADD, DIVIDE, SUBTRACT}))
	// fmt.Printf("8 / 3 + 7 * 3 = %f\n", Arithmetic([]int{8, 3, 7, 3}, []int{DIVIDE, ADD, MULTIPLY}))
	fmt.Printf("2 + 3 * 7 - 1 = %f\n", Arithmetic([]int{2, 3, 4, 1}, []int{ADD, MULTIPLY, SUBTRACT}))
	f := Arithmetic([]int{8, 3, 7, 3}, []int{DIVIDE, ADD, MULTIPLY})
	if math.Round(f) == f {
		fmt.Printf("true!\n")
	}
	ordered := []int{2, 4, 3, 1}
	sort.Ints(ordered)
	orderedStr := fmt.Sprintf("%v", ordered)
	fmt.Printf("orderedStr: %s\n", orderedStr)
}

// AllOrderedCombinations returns all possible ways that a set of things can be
// combined and repeated to make other sets of the same length.  For example,
// AllOrderedCombinations(2) returns: [[0,0],[1,1],[0,1],[1,0]].
func AllOrderedCombinations(max int) (result [][]int) {
	// addends will be a list of all possible ordered addends of max.  E.g.,
	// if max is 3 then addends will be [[3], [2,1], [1,2], [1,1,1]]
	// For our purposes the above means that a set of three things could
	// contain 3 things of the same type, 3 different things or 2 + 1 in
	// two different orders.
	addends := setsWithinSet(max - 1)

	// allSubsets will be a list of all possible ordered subsets.  E.g.,
	// if max is 3 then all Subsets will be:
	// [
	//  [[0],[1],[2]],
	//  [[0,1],[0,2],[1,0],[1,2],[2,0],[2,1]],
	//  [[0,1,2],[0,2,1],[1,0,2],[1,2,0],[2,0,1],[2,1,0]]
	// ]
	allSubsets := [][][]int{}
	for count := 1; count <= max; count++ {
		allSubsets = append(allSubsets, nMemberSubsets(count, max))
	}

	// We interpret each addend as a list of repeats.  For example [2,1] means we have
	// a set of three things made up of two of one thing followed by one of another thing.
	for _, repeaters := range addends {
		// repeatersLength describes the number of different types of things, [2,1] applied to apples
		// and pears means we have two apples and one pear, but still only two differnt types.
		repeatersLength := len(repeaters)
		// Index into allSubsets to get a list of the correct number of things.  E.g.,
		// allSubsets[0] is a list of single things.
		selectors := allSubsets[repeatersLength-1]
		// selection is a single case of a subset of the right length.  For example
		// allSubsets[1] will be a list of all possible sets of two things.  Each of these things
		// might occur once or several times.  We use repeaters to get all the possible
		// repetitions.
		for _, selection := range selectors {
			resultMember := []int{}
			for index, value := range selection {
				for count := 0; count < repeaters[index]; count++ {
					// repeaters[index] is the number of times that value is appended resultMember
					resultMember = append(resultMember, value)
				}
			}
			result = append(result, resultMember)
		}
	}
	return result
}
