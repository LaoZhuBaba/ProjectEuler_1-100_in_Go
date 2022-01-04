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

// IntCombinator represents a slice of integers.  It's purpose is to provide
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
	for v := range ic.ch {
		if v {
			// fmt.Printf("%v %v\n", v, ic.s)
			for k, n := range mpr.sequence {
				seq[k] = mpr.values[ic.s[n]]
			}
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
					fmt.Printf("%s\n", str)
				}
			}
			ic.ch <- true
		} else {
			return
		}
	}
}
