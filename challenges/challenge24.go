// A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the
// digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it
// lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

// 012   021   102   120   201   210

// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

package challenges

import (
	"fmt"
	"sort"
)

// func isMaxInList(n int, l []int) bool {
// 	for _, x := range l {
// 		if x > n {
// 			return false
// 		}
// 	}
// 	return true
// }

func getMaxInList(l []int) int {
	var max int
	for _, x := range l {
		if x > max {
			max = x
		}
	}
	return max
}
func getNextHigherInList(n int, l []int) int {
	minHigherThan := -1
	for _, x := range l {
		if (minHigherThan == -1 && x > n) || (x > n && x < minHigherThan) {
			minHigherThan = x
		}
	}
	return minHigherThan
}

var permu = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func rotateRight(l []int) {
	// fmt.Printf("rotateRight() called with %v\n", l)
	var temp int
	maxIndex := len(l) - 1
	if maxIndex < 1 {
		return
	}
	for c := maxIndex; c > 0; c-- {
		if c == maxIndex {
			temp = l[c]
		}
		l[c] = l[c-1]
	}
	l[0] = temp
}

func advance(p []int, countDown int) int {
	if countDown < 1 {
		return countDown
	}
	if len(p) == 2 {
		//		fmt.Printf("%v: %d\n", permu, countDown)
		p[0], p[1] = p[1], p[0]
		return countDown - 1
	}
	var nextHigher int
	for {
		countDown = advance(p[1:], countDown)
		if countDown < 1 {
			return countDown
		}
		if getMaxInList(p) == p[0] {
			break
		} else {
			//			fmt.Printf("%v: %d\n", permu, countDown)
			countDown--
			nextHigher = getNextHigherInList(p[0], p[1:])
			for nextHigher != p[0] {
				rotateRight(p)
			}
			sort.Ints(p[1:])
		}
	}
	return countDown
}

func Challenge24() {
	advance(permu, 1_000_000)
	fmt.Printf("%v\n", permu)
}
