package challenge81

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	size        = 80
	sizeSquared = size * size
	reallyBig   = 9_000_000_000_000_000_000
)

type Cell struct {
	val    int
	lChild *Cell
	lCost  int
	rChild *Cell
	rCost  int
}

// These are some data for testing.  If you want to use any of these
// you will need to update the size constant above which reprents the
// length of a row (or column) in the matrix.  And obviously change
// the paramenter passed to descendTree() within Challenge81().

// var matrixValues = [sizeSquared]int{
// 	131, 673, 234, 103, 18,
// 	201, 96, 342, 965, 150,
// 	630, 803, 746, 422, 111,
// 	537, 699, 497, 121, 956,
// 	805, 732, 524, 37, 331,
// }

// var matrixValues = [sizeSquared]int{
// 	1, 6, 2, 10,
// 	2, 9, 3, 6,
// 	6, 8, 7, 5,
// 	25, 6, 4, 2,
// }

// var matrixValues = [sizeSquared]int{
// 	1, 6, 3,
// 	4, 5, 10,
// 	7, 2, 9,
// }

// var matrixValues = [sizeSquared]int{
// 	1, 2,
// 	3, 4,
// }

type Matrix [sizeSquared]Cell

var matrix Matrix

// Load the data into an array of the correct size and return it
func loadMatrixFile() [sizeSquared]int {
	var allValues [sizeSquared]int
	var count int
	f, err := os.Open("p081_matrix.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err.Error())
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		for _, v := range strings.Split(s.Text(), ",") {
			var i int
			fmt.Sscanf(v, "%d", &i)
			allValues[count] = i
			count++
		}
	}
	return allValues
}

// Read the data as a simple series in an array.  From these create an array of
// cells with each cell having a value and pointers to left and right child cell.
// These pointers to child cells allow us to move down or right within the matrix.
// Depending on whether you think of the data as a matrix or a tree, "down" can
// also be considered "left".
func loadValues(matrix *Matrix, matrixValues *[sizeSquared]int) {
	for k, v := range *matrixValues {
		matrix[k].val = v
		if k%size != size-1 {
			matrix[k].rChild = &matrix[k+1]
		} else {
			// There is no right child for the last cell in each matrix row
			matrix[k].rChild = nil
		}
		if k/size != size-1 {
			matrix[k].lChild = &matrix[k+size]
		} else {
			// Cells in the last row in the matrix have no left child.
			matrix[k].lChild = nil
		}
	}
}
func dijkstraLoadValues(matrix *Matrix, matrixValues *[sizeSquared]int) {
	for k, v := range *matrixValues {
		if k == 0 {
			matrix[k].val = 0
		} else {
			matrix[k].val = reallyBig
		}
		if k%size != size-1 {
			matrix[k].rChild = &matrix[k+1]
			matrix[k].rCost = v
		} else {
			// There is no right child for the last cell in each matrix row
			matrix[k].rChild = nil
		}
		if k/size != size-1 {
			matrix[k].lChild = &matrix[k+size]
			matrix[k].lCost = v
		} else {
			// Cells in the last row in the matrix have no left child.
			matrix[k].lChild = nil
			// This is just to ensure we have somewhere to store the last value
			matrix[k].lCost = v
		}
	}
}

func dijkstraGenerateNextCell(m *Matrix) func() *Cell {
	count := 0
	distance := 0
	x := 0
	y := 0

	return func() *Cell {
		var ret *Cell
		if count >= size*size {
			return nil
		}
		// This loop will only iterate once unless the x or y are out of bounds (i.e., have a value > size).
		// This happens naturally because x & y are coordinates to points which are increasing distant.
		// Once we reach the (x,y) coordinate matching (size,size) then all greater distances will out-of
		// bounds.  This situation will be captured by the about check of count.
		for {
			if x < size && y < size {
				ret = &matrix[y*size+x]
			} else {
				ret = nil
			}
			// Let x grow up to distance after which increment distance and reset
			if x == distance {
				x = 0
				distance++
				y = distance

			} else {
				y = distance - x - 1
				x++
			}
			if ret == nil {
				continue
			}
			count++
			return ret
		}
	}
}

// An alternative solution using the Dijkstra algorithm.
// This algorithm requires a cost for each path (rather than each cell).  So in this solution cell.val
// records the cumulative total path cost and is initially set to reallyBig so that the first estimate
// of cumulative total is always better (i.e., lower) than the initial value.  One catch is that because
// we store the individual path costs in val.lCost and val.rCost it may seem like there is nowhere to
// store this value for the terminal cell because it has no valid children.  To work around this for the
// final row the dijkstraLoadValues() function stores this value in cell.lCost, even cell.lChild is nil.

func dijkstra(generator func() *Cell) int {
	var currentVal int
	for {
		cell := generator()
		if cell == nil {
			return currentVal
		}
		if cell.lChild != nil {
			if cell.lChild.val > cell.val+cell.lCost {
				cell.lChild.val = cell.val + cell.lCost
			}
		}
		if cell.rChild != nil {
			if cell.rChild.val > cell.val+cell.rCost {
				cell.rChild.val = cell.val + cell.rCost
			}
		}
		currentVal = cell.val + cell.lCost
	}
}

func descendTree(cell *Cell) int {
	var lVal, rVal int
	// This is a recursive function with two special cases which return directly...

	// Special case 1: if we are called with a nil pointer  this means there are no adjoining
	// cells. Return a big number to ensure that any number in the data will be preferred.
	if cell == nil {
		return reallyBig // This is just a sentinel value
	}
	// Special case 2: if we are at the lowest level of the matrix/triangle then both child
	// pointers will be nil.  In this case just return the value (which triggers a recursive
	// roll-up).
	if cell.lChild == nil && cell.rChild == nil {
		return cell.val
	}
	//  This is the non-special case which recurses...
	lVal = descendTree(cell.lChild)
	rVal = descendTree(cell.rChild)
	// After getting the left and right values (which represent a sum of the lower level data
	// to the left and right respectively) we add this sum to the current cell's value and
	// then set lVal and rVal to nil so that those level cells can no longer be visited.
	// Those cells are no longer relevant along this path because they have been accounted
	// for.  Finally, return the sumary value.
	if rVal < lVal {
		cell.val += rVal
	} else {
		cell.val += lVal
	}
	cell.rChild = nil
	cell.lChild = nil
	return cell.val
}

func Challenge81() {
	rawValues := loadMatrixFile()
	loadValues(&matrix, &rawValues)
	//	loadValues(&matrix, &matrixValues)
	fmt.Printf("Challenge 81 solution solved by recursion: %d\n", descendTree(&matrix[0]))
	dijkstraLoadValues(&matrix, &rawValues)
	generator := dijkstraGenerateNextCell(&matrix)
	fmt.Printf("Challenge 81 solution solved by Dijkstra algorithm: %d\n", dijkstra(generator))
}
