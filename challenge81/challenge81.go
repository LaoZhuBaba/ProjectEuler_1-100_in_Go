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
	lChild int
	rChild int
}

// These are some data for testing...
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

// Read the data as a simple series in an array.  From these create an array of cells
// with each cell having a value and a reference number for a left and right child cell.
// These child references all us to move down or right within the matrix.  Depending
// on whether you think of the data as a matrix or a tree, "down" is also "left".
func loadValues(matrix *[sizeSquared]Cell, matrixValues *[sizeSquared]int) {
	for k, v := range *matrixValues {
		matrix[k].val = v
		if k%size != size-1 {
			matrix[k].rChild = k + 1
		} else {
			matrix[k].rChild = -1
		}
		if k/size != size-1 {
			matrix[k].lChild = k + size
		} else {
			matrix[k].lChild = -1
		}
	}
}

func descendTree(t *[sizeSquared]Cell, n int) int {
	var lVal, rVal int
	// If we are called with n == -1  this means there are no adjoing cells.
	// Return a big number to ensure that any number in the data the will be preferred.
	if n == -1 {
		return reallyBig
	}
	if t[n].lChild == -1 && t[n].rChild == -1 {
		return t[n].val
	} else {
		lVal = descendTree(t, t[n].lChild)
		rVal = descendTree(t, t[n].rChild)
		if rVal < lVal {
			t[n].val += rVal
		} else {
			t[n].val += lVal
		}
		t[n].rChild = -1
		t[n].lChild = -1
		return t[n].val
	}
}

func Challenge81() {
	rawValues := loadMatrixFile()
	loadValues(&matrix, &rawValues)
	fmt.Printf("Challenge 81 solution: %d\n", descendTree(&matrix, 0))
}
