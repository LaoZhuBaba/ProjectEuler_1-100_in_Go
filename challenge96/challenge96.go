package challenge96

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const sudokuFile = "p096_sudoku.txt"

func loadFile() (all []int) {
	f, err := os.Open(sudokuFile)
	if err != nil {
		log.Fatalf("failed to open file: %s", sudokuFile)
	}
	b := bufio.NewScanner(f)
	b.Split(bufio.ScanLines)
	for b.Scan() {
		text := b.Text()
		d := make([]int, 9)
		if len(text) == 9 {
			fmt.Sscanf(text, "%1d%1d%1d%1d%1d%1d%1d%1d%1d", &d[0], &d[1], &d[2], &d[3], &d[4], &d[5], &d[6], &d[7], &d[8])
			all = append(all, d...)
		}
	}
	return all
}

func isIntInSlice(i int, s []int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

// findCandidates looks up relatedElementsIndexes to determine what values
// would clash.  Then returns a list of numbers 1-9 which don't clash.
func findCandidates(s []int, relatedElementsIndexes []int, maxValue int) (candidates []int) {
	clashes := []int{}
	candidates = []int{}

	for _, v := range relatedElementsIndexes {
		clashes = append(clashes, s[v])
	}
	for count := 1; count <= maxValue; count++ {
		if !isIntInSlice(count, clashes) {
			candidates = append(candidates, count)
		}
	}
	return candidates
}

// recursiveSolve takes i as an index into  slice s of max length where related_elements
// is a slice max length containing lists of related elements.  For in slice s, element x
// will have related_elements[x] which must all be different from x.  Any number in the
// range 1 - 9 which isn't on the related_element list is a possible solution candidate.
func recursiveSolve(i int, s []int, maxValue int, maxIndex int, relatedElements [][]int) bool {
	// start at i and skip ahead past elements which already have a number in the range 1 - 9.
	// fmt.Printf("recursiveSolve() called with i: %d maxValue: %d maxIndex: %d, s: %v\n", i, maxValue, maxIndex, s)
	// fmt.Printf("relatedElements: %v\n", relatedElements)
	count := i
	for s[count] != 0 {
		if count >= maxIndex {
			return true
		}
		count++
	}
	candidates := findCandidates(s, relatedElements[count], maxValue)
	for _, candidate := range candidates {
		s[count] = candidate
		if count == maxIndex {
			return true
		}
		if recursiveSolve(count+1, s, maxValue, maxIndex, relatedElements) {
			return true
		}
	}
	// if no candidate value provides a solution then reset s[i] back to 0 and
	// return false.  This path is no good.
	s[count] = 0
	return false
}

func solveSudoku(s []int, maxValue int, maxIndex int, related_elements [][]int) bool {
	return recursiveSolve(0, s, maxValue, maxIndex, related_elements)
}

func Challenge96() {
	var solution int
	all := loadFile()

	for count := 0; count < 50; count++ {
		if !solveSudoku(all[count*81:(count+1)*81], 9, 80, related_elements) {
			log.Fatalf("At least one Sudoku failed to solve\n")
		}
	}
	for grids := 0; grids < 81*50; grids += 81 {
		fmt.Printf("\n")
		for rows := 0; rows < 81; rows += 9 {
			thisRowStart := grids + rows
			if rows == 0 {
				solution += (all[thisRowStart] * 100) + (all[thisRowStart+1] * 10) + (all[thisRowStart+2])
			}
			fmt.Printf("%v\n", all[thisRowStart:thisRowStart+9])
		}
	}
	fmt.Printf("challenge 96 solution: %d\n", solution)

}
