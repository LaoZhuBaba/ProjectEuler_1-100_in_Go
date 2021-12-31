package challenge67

import (
	"bufio"
	"fmt"
	"os"
)

var val []int
var maxRowLen int

type Cell struct {
	val int
	//rollup   int
	lChild int
	rChild int
	//refCount int
}

func descendTree(t *[]Cell, n int) int {
	// fmt.Printf("walkTree called for cell: %d\n", n)
	var lVal, rVal int
	if (*t)[n].lChild == 0 && (*t)[n].rChild == 0 {
		// fmt.Printf("walkTree is returning %d for node %d\n", (*t)[n].val, n)
		return (*t)[n].val
	} else {
		lVal = descendTree(t, (*t)[n].lChild)
		rVal = descendTree(t, (*t)[n].rChild)
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

func printOptPath(tree *[]Cell, n int, level int) {
	fmt.Printf("%d\n", val[n])
	if level == maxRowLen {
		return
	}
	lVal := (*tree)[n+level].val
	rVal := (*tree)[n+level+1].val
	if lVal > rVal {
		printOptPath(tree, n+level, level+1)
	} else {
		printOptPath(tree, n+level+1, level+1)
	}
}

func Challenge67() {

	// Open the file containing the triangle data
	f, err := os.Open("triangle.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	rowLength := 1

	// The triangle file consists of two character strings, representing decimal numbers
	// separated by spaces or carriage returns.  Read one string (i.e., number) at a time
	// and convert to an int with fmt.Sscanf().  As we are doing this, keep track of the
	// number of columns in each row, which follows the triangular number pattern (each
	// row is one longer than the previous).  The point of this is to calculate the length
	// of the final row which we store in maxRowLen and use later.
	for count, column := 1, 1; scanner.Scan(); count++ {
		var i int
		fmt.Sscanf(scanner.Text(), "%d", &i)
		val = append(val, i)

		if column == rowLength {
			rowLength++
			column = 1
		} else {
			column++
		}
	}
	maxRowLen = rowLength - 1

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("Max row length is: %d\n", maxRowLen)

	// Now we have all the numbers in a slice of int called val.  And we know what the length of the
	// final row in the triangle will be.  Read these values into a slice of []Cell called tree.

	var tree = make([]Cell, len(val))

	count := 0
	var lChildNum, rChildNum int
	for row := 1; row <= maxRowLen; row++ {
		for col := 1; col <= row; col++ {
			// fmt.Printf("%d ", val[count])
			tree[count].val = val[count]
			lChildNum = count + row
			rChildNum = count + row + 1
			if lChildNum < len(tree) {
				tree[count].lChild = count + row
			}
			if rChildNum < len(tree) {
				tree[count].rChild = count + row + 1
			}
			count++
		}
		// fmt.Printf("\n")
	}

	grandTotal := descendTree(&tree, 0)
	fmt.Printf("Optimal path is:\n")
	printOptPath(&tree, 0, 1)
	fmt.Printf("=======\n%d (Challenge 67 solution)\n", grandTotal)
}
