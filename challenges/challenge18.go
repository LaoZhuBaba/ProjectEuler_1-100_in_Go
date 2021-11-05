// By starting at the top of the triangle below and moving to adjacent numbers on the row below,
// the maximum total from top to bottom is 23.

//      3
//     7 4
//    2 4 6
//   8 5 9 3

// That is, 3 + 7 + 4 + 9 = 23.

// Find the maximum total from top to bottom of the triangle below:

//  0                 75
//  1                95 64
//  3               17 47 82
//  6              18 35 87 10
// 10             20 04 82 47 65
// 15            19 01 23 75 03 34
// 21           88 02 77 73 07 63 67
// 28          99 65 04 28 06 16 70 92
// 36         41 41 26 56 83 40 80 70 33
// 45        41 48 72 33 47 32 37 16 94 29
// 55       53 71 44 65 25 43 91 52 97 51 14
// 66      70 11 33 28 77 73 17 78 39 68 17 57
// 78     91 71 52 38 17 14 91 43 58 50 27 29 48
// 91    63 66 04 68 89 53 67 30 73 16 69 87 40 31
// 105  04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

// NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route. However, Problem 67,
// is the same challenge with a triangle containing one-hundred rows; it cannot be solved by brute force, and requires a
// clever method! ;o)

package challenges

import "fmt"

const maxRowLength = 15

type Node struct {
	val int
	//rollup   int
	lChild int
	rChild int
	//refCount int
}

var t = make([]Node, len(values))

var values = []int{
	75,
	95, 64,
	17, 47, 82,
	18, 35, 87, 10,
	20, 04, 82, 47, 65,
	19, 01, 23, 75, 03, 34,
	88, 02, 77, 73, 07, 63, 67,
	99, 65, 04, 28, 06, 16, 70, 92,
	41, 41, 26, 56, 83, 40, 80, 70, 33,
	41, 48, 72, 33, 47, 32, 37, 16, 94, 29,
	53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14,
	70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57,
	91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48,
	63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31,
	04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23,
}

func walkTree(t *[]Node, n int) int {
	// fmt.Printf("walkTree called for node: %d\n", n)
	var lVal, rVal int
	if (*t)[n].lChild == 0 && (*t)[n].rChild == 0 {
		// fmt.Printf("walkTree is returning %d for node %d\n", (*t)[n].val, n)
		return (*t)[n].val
	} else {
		lVal = walkTree(t, (*t)[n].lChild)
		rVal = walkTree(t, (*t)[n].rChild)
		(*t)[n].lChild = 0
		(*t)[n].rChild = 0
		if lVal >= rVal {
			(*t)[n].val += lVal
			return (*t)[n].val
		} else {
			(*t)[n].val += rVal
			return (*t)[n].val
		}
	}
}

func printOptimalPath(t *[]Node, n int, level int) {
	fmt.Printf("%d\n", values[n])
	if level == maxRowLength {
		return
	}
	lVal := (*t)[n+level].val
	rVal := (*t)[n+level+1].val
	if lVal > rVal {
		printOptimalPath(t, n+level, level+1)
	} else {
		printOptimalPath(t, n+level+1, level+1)
	}
}

func Challenge18() {
	count := 0
	var lChildNum, rChildNum int
	for row := 1; row <= maxRowLength; row++ {
		for col := 1; col <= row; col++ {
			fmt.Printf("%d ", values[count])
			t[count].val = values[count]
			lChildNum = count + row
			rChildNum = count + row + 1
			if lChildNum < len(t) {
				t[count].lChild = count + row
			}
			if rChildNum < len(t) {
				t[count].rChild = count + row + 1
			}
			count++
		}
		fmt.Printf("\n")
	}
	//fmt.Printf("t is: %v\n", t)
	grandTotal := walkTree(&t, 0)
	fmt.Printf("Optimal path is:\n")
	printOptimalPath(&t, 0, 1)
	fmt.Printf("=======\n%d\n", grandTotal)
}
