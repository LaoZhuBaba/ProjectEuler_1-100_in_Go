package challenge90

import (
	"fmt"
)

// Subset returns a list of n-length subnets of l
func Subset(l []int, n int) (ret [][]int) {

	if n < 1 {
		return [][]int{{}}
	}
	// Each recursion decrements n.  When n reaches 1 then return
	// a list of the remaining elements.  This terminates the recursion.
	if n == 1 {
		for _, value := range l {
			ret = append(ret, []int{value})
		}
		return ret
	}

	max := len(l) - n

	// Iterate along the list so long as enough
	// elements remain to make an n-length subset.
	for count := 0; count <= max; count++ {
		for _, val := range Subset(l[count+1:], n-1) {
			// Consider l[count] to be our "cursor" position. On each iteration make a
			// recursive call to Subset, passing a slice which starts one element to the
			// right of the cursor.  For each "row" returned, create a new slice which
			// has the value at our cursor position as its first element and the rest
			// made up of that "row".  This creates a slice of slices.
			ret = append(ret, append([]int{l[count]}, val...))
		}
	}
	return ret
}

func IntInSlice(s []int, n int) bool {
	for _, val := range s {
		if val == n {
			return true
		}
	}
	return false
}
func Split2DigNum(n int) (a, b int) {
	return n / 10, n % 10
}

func ContainsN(a, b []int, n int) (ret int) {
	if IntInSlice(a, n) {
		ret += 1
	}
	if IntInSlice(b, n) {
		ret += 2
	}
	return ret
}

func ContainsNN(nn int, sa, sb []int) bool {
	a, b := Split2DigNum(nn)
	containsA := ContainsN(sa, sb, a)
	if containsA == 0 {
		return false
	}
	containsB := ContainsN(sa, sb, b)
	if containsB == 0 {
		return false
	}
	if containsA == 1 {
		if IntInSlice(sb, b) {
			return true
		}
	}
	if containsA == 2 {
		if IntInSlice(sa, b) {
			return true
		}
	}
	if containsA == 3 {
		if IntInSlice(sb, b) || IntInSlice(sa, b) {
			return true
		}
	}
	return false
}

func Contains01(sa, sb []int) bool {
	return ContainsNN(1, sa, sb)
}

func Contains06(sa, sb []int) bool {
	return ContainsNN(6, sa, sb)
}

func Contains16(sa, sb []int) bool {
	return ContainsNN(16, sa, sb)
}

func Contains25(sa, sb []int) bool {
	return ContainsNN(25, sa, sb)
}

func Contains36(sa, sb []int) bool {
	return ContainsNN(36, sa, sb)
}

func Contains46(sa, sb []int) bool {
	return ContainsNN(46, sa, sb)
}

func Contains64(sa, sb []int) bool {
	return ContainsNN(64, sa, sb)
}

func Contains81(sa, sb []int) bool {
	return ContainsNN(81, sa, sb)
}

func ContainsAll(sa, sb []int) bool {
	for _, nn := range []int{1, 4, 25, 81} {
		if !ContainsNN(nn, sa, sb) {
			return false
		}
	}

	if !ContainsNN(9, sa, sb) && !ContainsNN(6, sa, sb) {
		return false
	}
	if !ContainsNN(49, sa, sb) && !ContainsNN(46, sa, sb) {
		return false
	}
	if !ContainsNN(36, sa, sb) && !ContainsNN(39, sa, sb) {
		return false
	}
	if !ContainsNN(19, sa, sb) && !ContainsNN(16, sa, sb) {
		return false
	}
	// 64/94 are reversals of 46/49 so already covered
	// if !ContainsNN(64, sa, sb) && !ContainsNN(94, sa, sb) {
	// 	return false
	// }
	return true
}
func IsIntInSlice(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}

func testCombos(data [][]int) {
	l := len(data)
	var count int

	for a := 0; a < l; a++ {
		// Start inner iterations at a+1 to ensure we don't count
		// both (data[a], data[b]) and (data[b], data[a]).
		for b := a + 1; b < l; b++ {
			if ContainsAll(data[a], data[b]) {
				count++
				fmt.Printf("count: %04d, a: %v, b: %v\n", count, data[a], data[b])
			}
		}
	}
}

func Challenge90() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	data := Subset(input, 6)
	testCombos(data)
}
