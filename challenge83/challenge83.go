package challenge83

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
	distance int     // The total calculated distance via the currently known best path.  Initialised to "reallyBig"
	cost     int     // The cost of links to child cells. In this model, the value is common to all three links
	siblings []*Cell // A slice of pointers to other cells.  For example each non-edge cell has four siblings.
}

type Matrix [size][size]Cell

// Calculate the path from top left corner cell to bottom right using Dijkstra's algorithm.
func (m *Matrix) dijkstra() int {
	inProgress := make(map[*Cell]bool)
	//  Seed the inProgress list with the top right corner cell as the starting point
	inProgress[&m[0][0]] = true

	// We have to be careful here because we are iterating through values using range but
	// within the loop we are updating the same map which range is based on.  There is no
	// conflict because the list we get from range is a COPY.  However, it means we need
	// an outer for statement so that a new range is recreated as soon as the old completes,
	// until inProgress is empty.
	for len(inProgress) != 0 {
		for cell := range inProgress {
			for _, child := range cell.siblings {
				// The heart of Dijkstra's algorithm is that any child cell should be
				// assigned a distance equal to the current cell plus the local path cost,
				// UNLESS that child cell has already been assigned a shorter distance.
				distanceViaThisCell := cell.distance + cell.cost
				if child.distance > distanceViaThisCell {
					child.distance = distanceViaThisCell
					// Once a distance has been assigned, the cell goes onto the inProgress
					// list so that distances to its children can be (tentatively) calculated.
					// If the distance for a cell is updated later then it has to go back
					// onto the inProgress list.  This may trigger a chain of dependent
					// cells going back on the list, but it will all work itself out. I
					// think it might be possible for a cell distance to be updated multiple
					// times before it is recalculated but because we are using a map it
					// it can't actually be on the list twice.  This is good.  As far as I
					// know, the order of cells on the inProgress list will not affect the
					// final result--although I guess it may affect performance.
					inProgress[child] = true
				}
			}
			// Once a distance has been assigned for every child cell (or skipped due to a lower distance
			// already being present) then we are finished with this cell--for the moment at least.
			delete(inProgress, cell)
		}
	}
	finalCell := &m[size-1][size-1]
	// For all the other cells the link cost associated with the cell is added to the next cell in
	// the chain.  This doesn't happen with the last cell so we need to remember to add this value
	// to the final tally.  The underlying point is that the number of cells in the chain is one
	// greater than the number of links, so we always have an extra value value to add at some point.
	finalCell.distance += finalCell.cost
	return finalCell.distance
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
			thisCell.siblings = make([]*Cell, 0)
			var i int
			fmt.Sscanf(v, "%d", &i)
			thisCell.cost = i
			// Initial distance is zero for the starting cell in the top left corner, Initialise all other cells
			// with a distance of reallyBig.

			if k == 0 && row == 0 {
				thisCell.distance = 0
			} else {
				thisCell.distance = reallyBig
			}

			// All cells have a pointer to the adjacent cell on the left except cells in the leftmost column.
			if k != 0 {
				thisCell.siblings = append(thisCell.siblings, &m[row][k-1])
			}
			// All cells have a pointer to the adjacent cell on the right except cells in the rightmost column.
			if k != size-1 {
				thisCell.siblings = append(thisCell.siblings, &m[row][k+1])
			}
			// All cells have a pointer to the cell immediately above except cells in the first row.
			if row != 0 {
				thisCell.siblings = append(thisCell.siblings, &m[row-1][k])
			}
			// All cells have a pointer to the cell immediately below except cells in the last row.
			if row != size-1 {
				thisCell.siblings = append(thisCell.siblings, &m[row+1][k])
			}
		}
		row++
	}
}

func Challenge83() {
	var matrix Matrix
	loadMatrixFile("p082_matrix.txt", &matrix)
	solution := matrix.dijkstra()
	//fmt.Printf("%v\n", matrix)
	fmt.Printf("Challenge 83 solution is: %d\n", solution)
}
