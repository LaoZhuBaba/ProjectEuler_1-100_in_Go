package challenge91

import (
	"fmt"
	"log"
)

type Vector struct {
	x int
	y int
}

type Grid struct {
	size int
}

func (g Grid) IsWithin(v Vector) bool {
	return v.x >= 0 && v.y >= 0 && v.x <= g.size && v.y <= g.size
}

func (v Vector) lPerp() (rv Vector) {
	rv, err := Vector{-v.y, v.x}.simplify()
	if err != nil {
		log.Fatalf("in method lPerp failed to simplify vector: %s\n", rv)
	}
	return rv
}

func (v Vector) rPerp() (rv Vector) {
	rv, err := Vector{v.y, -v.x}.simplify()
	if err != nil {
		log.Fatalf("in method rPerp failed to simplify vector: %s\n", rv)
	}
	return rv
}
func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y}
}

func (v Vector) String() string {
	return fmt.Sprintf("(%d,%d)", v.x, v.y)
}

// For a given Grid return a list of all the vectors within that Grid
// where the y coordinate is greater or equal to the x coordinate. This
// This is useful because we for ever (x, y) there will be an equal number
// of trianges formed by (y, x) so we can just double the number, except
// where x == y.
func (g Grid) getGradiantsGtEq1() (ret []Vector) {
	if g.size < 2 {
		return []Vector{}
	}
	for x, y := 0, 1; x <= y; {
		ret = append(ret, Vector{x, y})
		if x == y {
			if x == g.size-1 {
				return ret
			}
			y += 1
			x = 0
		} else {
			x += 1
		}
	}
	return ret
}

// simplify shortens a vector to a minimum whole number values
// which points in the same direction.  So 2,2 simplifies to 1,1
// They point point up and to the right at a 45 degree angle.
func (v Vector) simplify() (rv Vector, err error) {
	var swap, negX, negY bool
	var big, small int

	defer func() {
		if swap {
			rv.x, rv.y = rv.y, rv.x
		}
		if negX {
			rv.x *= -1
		}
		if negY {
			rv.y *= -1
		}
	}()

	// (0, 0) vector doesn't make sense for our purposes to treat as an error
	if v.x == 0 && v.y == 0 {
		return v, fmt.Errorf("(0,0) is not a valid vector for our purposes")
	}

	// We need to convert the vector to a standard form where neither number is negative
	// and x is less than y.  Record the changes so we can rever back in defer()
	if v.x < 0 {
		v.x *= -1
		negX = true
	}

	if v.y < 0 {
		v.y *= -1
		negY = true
	}

	if v.x == v.y {
		return Vector{1, 1}, nil
	}

	if v.y > v.x {
		small, big = v.x, v.y
		swap = false
	} else {
		small, big = v.y, v.x
		swap = true
	}

	if small == 0 {
		return Vector{0, 1}, nil
	}

	if small == 1 || big == 3 {
		return Vector{small, big}, nil
	}

	max_divider := big / 2
	for divider := max_divider; divider > 1; divider-- {
		if (small%divider == 0) && (big%divider == 0) {
			return Vector{small / divider, big / divider}, nil
		}
	}
	return Vector{small, big}, nil
}

func newGrid(size int) Grid {
	return Grid{
		size: size,
	}
}

func countTriangles(v Vector, gridSize int) (count int) {
	// lp is the shortest vector which perpendicular to v on the left
	lp := v.lPerp()
	// rp is the shortest vector which perpendicular to v on the right
	rp := v.rPerp()

	// extend lp to the left perpendicular side until it goes outside the grid
	for lnext := v.Add(lp); lnext.x >= 0 && lnext.y <= gridSize; lnext = lnext.Add(lp) {
		count++
		if v.x != v.y {
			// There will also be a symetrically opposite triange except where v.x == v.y
			count++
		}
	}
	// extend rp to the right perpendicular side until it goes outside the grid
	for rnext := v.Add(rp); rnext.x <= gridSize && rnext.y >= 0; rnext = rnext.Add(rp) {
		count++
		if v.x != v.y {
			// There will also be a symetrically opposite triange except where v.x == v.y
			count++
		}
	}
	return count
}

func Challenge91() {
	// The question is framed in terms of the length if the squares side
	const lengthOfSide = 50
	// The number of grid points is always one greater per side.  I.e., a 3 x 3 grid = 2 x 2 squares
	const gridSize = lengthOfSide + 1

	//  There are two different types of right-angle triangles to count:
	//  * triangles where the right angle vertex is at grid coordinate (0,0).
	//    the number of these is equal to the square of lengthOfSize
	//  * others where the right angle is at another vertex.
	var count int = lengthOfSide * lengthOfSide

	g := newGrid(gridSize)
	list := g.getGradiantsGtEq1()
	for _, v := range list {
		count += countTriangles(v, lengthOfSide)
	}
	fmt.Printf("count: %v\n", count)
}
