// For this challenge I wanted to play with channels and the idea that a function could
// be called once and then return each member of a the required result set via a channel--
// one at a time.  A bit like using "yield" in Python.  This means that the solution is
// really unintelligent in some respects because we look at the every possible combination
// of values in a set, but in reality, many of them could be excluded without any deep
// consideration.  I eventually found quite a long list of potential answers and then had
// to go back and read the question carefully to work out which ones could be excluded.
// For example, it is pretty obvious from the question that the first digit in the solution
// must be 6.  Another point is that I effectively found each solution several times, with
// only a difference in ordering.

package challenge68

import (
	"euler/shared"
	"fmt"
	"sort"
)

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
func runCombinations(p []int, ch chan bool) {
	if len(p) == 2 {
		ch <- true
		<-ch
		p[0], p[1] = p[1], p[0]
		ch <- true
		<-ch
		return
	}
	var nextHigher int
	for {
		runCombinations(p[1:], ch)
		if getMaxInList(p) == p[0] {
			break
		} else {
			nextHigher = getNextHigherInList(p[0], p[1:])
			for nextHigher != p[0] {
				shared.RotateRight(p)
			}
			sort.Ints(p[1:])
		}
	}
}

// Return every possibly combination the slice of integers in ic.s.
// This slice is repeatedly updated and after update "true" will be
// sent so the caller knows to read the updated value.  After all
// possible combinations have been provide send "false" to the.
// caller.
func (ic IntCombinator) RunCombinations() {
	runCombinations(ic.s, ic.ch)
	ic.ch <- false
}

type IntCombinator struct {
	s  []int
	ch chan bool
}

// IntCombinator represents a slice of integers.  Its purpose is to provide
// all possible combinations of those integers.
func newIntCombinator(s []int) *IntCombinator {
	ic := new(IntCombinator)
	ic.ch = make(chan bool)
	ic.s = make([]int, len(s))
	copy(ic.s, s)
	return ic
}

type magicPentagonRing struct {
	values   [10]int
	sequence [15]int
}

func NewMagicPentagonRing(values [10]int) *magicPentagonRing {
	mpr := new(magicPentagonRing)
	mpr.values = [10]int{values[0], values[1], values[2], values[3], values[4], values[5], values[6], values[7], values[8], values[9]}
	mpr.sequence = [15]int{5, 0, 1, 6, 1, 2, 7, 2, 3, 8, 3, 4, 9, 4, 0}
	return mpr
}
func sumThree(a []int) int {
	return a[0] + a[1] + a[2]
}
func Challenge68() {
	ic := newIntCombinator([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	// var s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// var s = []int{0, 1, 2, 3, 4, 5, 6}
	mpr := NewMagicPentagonRing([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	go ic.RunCombinations()

	var seq [15]int
	var sum, solution int
	for v := range ic.ch {
		if v {
			// fmt.Printf("%v %v\n", v, ic.s)
			for k, n := range mpr.sequence {
				seq[k] = mpr.values[ic.s[n]]
			}
			// See initial notes above.  The correct answer has to begin with 6 because the question
			// makes clear that the inner "pentagon" is made of the numbers 1-5 and the instructionss
			// state: "starting from ... the numerically lowest external node" which only mean 6.  Abd
			// this means that the second and third values must be less than 6 because because each
			// group of three values must start with a value from the outer pentagon and continue with
			// two values from the inner pentagon.
			if seq[0] == 6 && seq[1] < 6 && seq[2] < 6 {

				//			fmt.Printf("sequence: %v\n", seq)
				if sumThree(seq[0:3]) == sumThree(seq[3:6]) &&
					sumThree(seq[3:6]) == sumThree(seq[6:9]) &&
					sumThree(seq[6:9]) == sumThree(seq[9:12]) &&
					sumThree(seq[9:12]) == sumThree(seq[12:15]) {
					// i := shared.IntSliceToInt(seq[:])
					// if len(fmt.Sprintf("%d", i)) != 16 {
					// 	fmt.Printf("Too long\n")
					// }
					// fmt.Printf("BINGO %v\n", seq)
					str := ""
					for _, v := range seq {
						str = (str + fmt.Sprintf("%d", v))
					}
					if len(str) == 16 {
						fmt.Sscanf(str, "%d", &sum)
						if sum > solution {
							solution = sum
						}
						// fmt.Printf("%s\n", str)
					}
				}
			}
			ic.ch <- true
		} else {
			fmt.Printf("solution 68 is: %d\n", solution)
			return
		}
	}
}
