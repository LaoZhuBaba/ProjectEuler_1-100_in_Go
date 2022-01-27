package challenge85

import "fmt"

// For a rectangle with sides (x,y) calculate the number of rectangles of size not exceeding (x,y)
// which can be contained within that rectangle.  For example a (2, 1) rectangle contains two (1,1)
// rectangles, plus one (2,1) rectangle, so this function would return 3.  A (2,2) rectangle contains
// four (1,1) rectangles, two (2,1) rectangles, two (1,2) rectangles and one (2,2) rectangles, for a
// total of 9.
func rectInRect(x, y int) int {
	var total int
	for xCount := 1; xCount <= x; xCount++ {
		for yCount := 1; yCount <= y; yCount++ {
			total += (x - xCount + 1) * (y - yCount + 1)
		}
	}
	return total
}

// The challenge is to find the rectangle for which the number of contained rectangles is closest to target.
const target = 2_000_000

func Challenge85() {

	n := 1
	incr := 100
	for rectInRect(1, n) < target {
		n += incr
		if incr > 1 {
			incr -= 1
		}
	}
	// n is now higher than the minimum value where a 1 x n rectangle contains "target"
	// possible smaller rectangles.  We know from this that if we increment from rectangles
	// containing 1 row to n rows we will definitely have covered every possibility.
	bestDelta := target // An impossibly high initial value
	var bestXY [2]int
	var rir, delta int
	for x := 1; x <= n; x++ {
		y := 1
		for {
			rir = rectInRect(x, y)
			if rir > target {
				// At this point we have found an (x,y) combination which has > target sub-rectangles.
				// We need to check if this is the closest combination and also check (x-1, y) because
				// this could be even closer but LESS THAN target.
				delta = rir - target
				if delta < bestDelta {
					bestDelta = delta
					bestXY = [2]int{x, y}
				}
				rir = rectInRect(x-1, y)
				if rir < target {
					delta = target - rir
				} else {
					delta = rir - target
				}
				if delta < bestDelta {
					bestDelta = delta
					bestXY = [2]int{x - 1, y}
				}
				break
			}
			y++
		}
	}
	fmt.Printf("bestDelta is: %d\n", bestDelta)
	fmt.Printf("bestXY values are: %v\n", bestXY)
	fmt.Printf("Challenge 85 solution is: %d\n", bestXY[0]*bestXY[1])
}
