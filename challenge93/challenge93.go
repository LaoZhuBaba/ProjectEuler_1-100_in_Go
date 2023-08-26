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
	for count := 1; count <= 5; count++ {
		result := setsWithinSet(count)
		fmt.Printf("result: %v\n", result)
		fmt.Printf("count: %d, length %d\n", count, len(result))
	}
}
