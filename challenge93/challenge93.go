package challenge93

import (
	"fmt"
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
	// if n == len(set) {
	// 	return [][]int{set}
	// }
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
		// for _, tailSet := range setsWithinSet(tail) {
		// 	headSet = append(headSet, append([]int{head}, tailSet...))
		// }
		// sets = append(sets, headSet...)
		subsets = append(subsets, headSet...)
	}
	return subsets
}

// func permutationsWithRepeats(level int, set []rune, input []rune) (output [][]rune) {
// 	if level == 2 {
// 		return output
// 	}

//		var inputCopy []rune
//		inputCopy = append(inputCopy, input...)
//		for _, r := range set {
//			output = append(output, inputCopy)
//			output[len(output)-1][level] = r
//			// fmt.Printf("%s%c\n", strings.Repeat(".", level), r)
//		}
//		return append(permutationsWithRepeats(level+1, set, output[len(output)-1]), output...)
//	}
func Challenge93() {
	// var output [][]rune
	// start := []rune{'0', '0', '0', '0'}

	// output = permutationsWithRepeats(
	// 	0,
	// 	[]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'},
	// 	start,
	// )
	// fmt.Printf("%c\n", output)
	const max = 4
	subsets := setsWithinSet(max)
	// for count := 1; count <= max; count++ {
	// 	result := setsWithinSet(count)
	// 	fmt.Printf("result: %v\n", result)
	// 	fmt.Printf("count: %d, length %d\n", count, len(result))
	// }
	//subsets := []int{0, 1, 2, 3}
	operators := []rune{'+', '-', '*', '/'}
	fmt.Printf("subsets: %v\n", subsets)
	all := [][][]int{}
	for count := 1; count <= max; count++ {
		all = append(all, nMemberSubsets(count, max))
	}
	//result := [][]rune{}
	fmt.Printf("all: %v\n", all)

	var operatorSequence [][]rune
	for _, repeaters := range subsets {
		repeatersLength := len(repeaters)
		fmt.Printf("repeaters: %v\n", repeaters)
		fmt.Printf("repeatersLength: %d\n", repeatersLength)
		selectors := all[repeatersLength-1]
		fmt.Printf("selectors: %v\n", selectors)
		for _, selection := range selectors {
			operations := []rune{}
			for index, i := range selection {
				operator := operators[i]
				for count := 0; count < repeaters[index]; count++ {
					operations = append(operations, operator)
				}
			}
			operatorSequence = append(operatorSequence, operations)
		}
		// for index, repeater := range repeaters {
		// 	for _, selected := range selectors[index] {

		// 	}
		// }
		// for _, repeater := range repeaters {
		// 	for count := 0; count < repeater; count++ {

		// 	}
		// }
		// rowset := []rune{}
		// for _, val := range subset {
		// 	fmt.Printf("all[val-1]: %v\n", all[val-1])
		// 	// for _, v := range all[val-1] {
		// 	// 	rowset = append(rowset, operators[v])
		// 	// }
		// 	result = append(result, rowset)
		// }
		fmt.Printf("result: %c\n", operatorSequence)
	}
}
