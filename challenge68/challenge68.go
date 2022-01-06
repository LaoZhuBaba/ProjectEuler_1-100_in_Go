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

// Return every possible combination of the slice of integers in ic.s. This slice is
// repeatedly updated and after update "true" will be sent so the caller knows to read
// the updated value.  After all possible combinations have been provided signal
// completion with ic,done()
func (ic IntCombinator) RunCombinations() {
	var i, nextHigher int
	defer ic.done()
	// This was originally implemented with recursion but this created undesireable
	// complications.  Reimplemented later with i representing what was originally
	// the depth of recursion.
	for {
		sl := ic.s[i:]
		// If sl has only two elements then just reverse them.
		if len(sl) == 2 {
			ic.ch <- true
			<-ic.ch
			sl[0], sl[1] = sl[1], sl[0]
			ic.ch <- true
			<-ic.ch
			i--
			continue
		}
		// Keep moving i to the right if the original small-to-big order
		// has not been reversed.
		if sl[2] > sl[1] {
			i++
			continue
		}
		if getMaxInList(sl) == sl[0] {
			// Out ultimate aim is for the order of integers to be completely reversed,
			// with sl[0] being the highest numbered element.  If this condition is met
			// when i == 0 then we're done.  If this condition is met when i > 0 then
			// decrement i and continue.
			if i == 0 {
				return
			}
			i--
			continue
		}
		// If i != 2 AND sl[0] is not yet at max value and sl[2] is NOT > sl[1] then rotate
		// the slice until the next higher value is at sl[0] and then sort the other values
		// from lowest to highest.
		nextHigher = getNextHigherInList(sl[0], sl[1:])
		for nextHigher != sl[0] {
			shared.RotateRight(sl)
		}
		sort.Ints(sl[1:])
		i++
	}
}

type IntCombinator struct {
	s    []int
	ch   chan bool
	done func()
}

// IntCombinator represents a slice of integers which is sorted from low to high. Its purpose is to
// provide all possible combinations of those integers.  At the end of the process the slice will be
// ordered from high to low.
func newIntCombinator(s []int) *IntCombinator {
	ic := new(IntCombinator)
	ic.ch = make(chan bool)
	ic.s = make([]int, len(s))
	ic.done = func() {
		close(ic.ch)
	}
	// We could also do ic.s = s rather then copying, but the caller may not expect s to be updated
	// so safest to copy I think.
	copy(ic.s, s)
	sort.Ints(ic.s)
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

func sumInt(s ...int) (i int) {
	for _, v := range s {
		i += v
	}
	return
}
func Challenge68() {
	ic := newIntCombinator([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	// ic := newIntCombinator([]int{0, 1, 2, 3})
	// var s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// var s = []int{0, 1, 2, 3, 4, 5, 6}
	mpr := NewMagicPentagonRing([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	go ic.RunCombinations()

	var seq [15]int
	var sum, solution int
	for <-ic.ch {
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
			if sumInt(seq[0:3]...) == sumInt(seq[3:6]...) &&
				sumInt(seq[3:6]...) == sumInt(seq[6:9]...) &&
				sumInt(seq[6:9]...) == sumInt(seq[9:12]...) &&
				sumInt(seq[9:12]...) == sumInt(seq[12:15]...) {
				str := ""
				for _, v := range seq {
					str = (str + fmt.Sprintf("%d", v))
				}
				if len(str) == 16 {
					fmt.Sscanf(str, "%d", &sum)
					if sum > solution {
						solution = sum
					}
				}
			}
		}
		// Acknowledge receipt of message, so ic.RunCombinations() can continue
		ic.ch <- true
	}
	fmt.Printf("solution 68 is: %d\n", solution)
}
