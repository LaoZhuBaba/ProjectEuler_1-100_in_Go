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
	rChild *Cell
}

// These are some data for testing.  If you want to use any of these
// you will need to update the size constant above which reprents the
// length of a row (or column) in the matrix.  And obviously change
// the paramenter passed to descendTree() within Challenge81().
//
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

var matrix [sizeSquared]Cell

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
func loadValues(matrix *[sizeSquared]Cell, matrixValues *[sizeSquared]int) {
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
	fmt.Printf("Challenge 81 solution: %d\n", descendTree(&matrix[0]))
}
