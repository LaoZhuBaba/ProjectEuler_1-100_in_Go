package challenge82

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	size      = 80
	reallyBig = 9_000_000_000_000_000_000
)

type Cell struct {
	distance int // The total calculated distance via the currently known best path.  Initialised to "reallyBig"
	cost     int // The cost of links to child cells in this model, the value is common to all three links
	uChild   *Cell
	dChild   *Cell
	rChild   *Cell
}

type Matrix [size][size]Cell

// Calculate the minimum cost from any cell in the leftmost column of a matrix to any cell in rightmost column.
// Uses Dijkstra's algorithm.

func (m *Matrix) dijkstra() int {
	minDistance := reallyBig
	inProgress := make(map[*Cell]bool)
	//  Seed the inProgress list with the first column of cells as a starting point
	for index := range m {
		inProgress[&m[index][0]] = true
	}

	// I'm not sure how Go handles iterating through a map which the loop is updating?
	// I'm sure the range won't be magically updated so I've done it this way to ensure
	// when the range loop completes the range will immediately be recalculated until
	// the length of inProgress zero.
	for len(inProgress) != 0 {
		for cell := range inProgress {
			for _, child := range []*Cell{cell.uChild, cell.dChild, cell.rChild} {
				// If child == cell then this is an edge node, so ignore.
				if child != cell {
					// The heart of Dijkstra's algorithm is that any child cell should be
					// assigned a distance equal to the current cell plus the local path cost,
					// UNLESS that child cell has already been assigned a shorter distance.
					distanceViaThisCell := cell.distance + cell.cost
					if child.distance > distanceViaThisCell {
						child.distance = distanceViaThisCell
						// Once a distance has been assigned, the cell goes on the inProgress
						// list so that distances to its children can be (tentatively) calculated.
						inProgress[child] = true
					}
				}
			}
			// Once a distance has been assigned for every child cell (or skipped due to a
			// lower distance already being present) then we are finished with this cell.
			delete(inProgress, cell)
		}
	}
	// The last column is a special case because for all other cells their cost is added to the
	// distance of their child cells.  But for the final cell we have to add on the cost to the
	// cell itself, or it won't be included.  The basic issue is that the number of links is
	// aways one greater than the number of cell, so when we iterate through cells there will
	// always be one extra cost to add.
	for _, row := range m {
		lastColumn := &row[len(row)-1]
		lastColumn.distance += lastColumn.cost
		if lastColumn.distance < minDistance {
			minDistance = lastColumn.distance
		}
	}
	return minDistance
}

// Load the data into an array of the correct size and return it
func loadMatrixFile(fileName string, m *Matrix) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err.Error())
	}
	s := bufio.NewScanner(f)
	var row int
	for s.Scan() {
		for k, v := range strings.Split(s.Text(), ",") {
			thisCell := &m[row][k]
			var i int
			fmt.Sscanf(v, "%d", &i)
			thisCell.cost = i
			// Initial distance is zero for the left-most starting cells or "reallyBig" for other cells
			// reallyBig is a sentinel value which should lose any comparison against a calculated path.
			if k == 0 {
				thisCell.distance = 0
			} else {
				thisCell.distance = reallyBig
			}
			// The rChild value points to the same row in the next column except for the rightmost column
			// where it points to its own cell to indicate an edge.
			if k == size-1 {
				thisCell.rChild = thisCell
			} else {
				thisCell.rChild = &m[row][k+1]
			}
			// The uChild value points to the same column in row-1.  Or self pointer for the first/top row.
			if row == 0 {
				thisCell.uChild = thisCell
			} else {
				thisCell.uChild = &m[row-1][k]
			}
			// The dChild value points to the samee colum in row+1.  Or self pointer for the last/bottom row.
			if row == size-1 {
				thisCell.dChild = thisCell
			} else {
				thisCell.dChild = &m[row+1][k]
			}
		}
		row++
	}
}

func Challenge82() {
	var matrix Matrix
	loadMatrixFile("p082_matrix.txt", &matrix)
	fmt.Printf("Challenge 82 solution is: %d\n", matrix.dijkstra())
}
