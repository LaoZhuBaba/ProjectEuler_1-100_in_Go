package challenges

import "fmt"

func createPtrList(included, excluded []*[]int) []*[]int {
	returnList := make([]*[]int, 0)
	var skip bool
	for _, i := range included {
		skip = false
		for _, e := range excluded {
			if i == e {
				skip = true
			}
		}
		if skip {
			continue
		}
		returnList = append(returnList, i)
	}
	return returnList
}

func triangle(i int) int {

	return i * (i + 1) / 2
}
func square(i int) int {
	return i * i
}
func pentagon(i int) int {

	return i * (3*i - 1) / 2
}

func hexagon(i int) int {

	return i * (2*i - 1)
}
func heptagon(i int) int {

	return i * (5*i - 3) / 2
}
func octagon(i int) int {

	return i * (3*i - 2)
}

// An example of a match is be 1234 & 3456
func checkMiddleOverlap(i1, i2 int) bool {
	return (i1-(i2/100))%100 == 0
}

// An example of a match is 1234 & 5612
func checkOuterOverlap(i1, i2 int) bool {
	return (i2-(i1/100))%100 == 0
}
func makePolygonalList(floor, ceiling int, f func(int) int) []int {
	l := make([]int, 0)
	for count := 1; ; count++ {
		transformed := f(count)
		if transformed < floor {
			continue
		}
		if transformed > ceiling {
			return l
		}
		l = append(l, transformed)
	}
}

func linkTogether(include, exclude []*[]int, chain []int) bool {
	ptrList := createPtrList(include, exclude)
	if len(ptrList) > 0 {
		for _, p := range ptrList {
			for _, v := range *p {
				// Check to see if first two digits of v match the last two digits of the last integer
				// in the chain
				if checkMiddleOverlap(chain[len(chain)-1], v) {
					// Call this function recursively with p added to the exclude list to ensure each type
					// of number is only considered once.
					if linkTogether(include, append(exclude, p), append(chain, v)) {
						return true
					}
				}
			}
		}
	} else {
		// This is the case where we've linked together one of each type of number so we now have
		// to check if the first two digits of the first integer in the chain match the last two digits
		// in the last integer in the change.  If they do, then print the solution and then return
		// 'true' which will cause all nested instances of this function to return true in sequence.
		if checkOuterOverlap(chain[0], chain[len(chain)-1]) {
			var solution int
			for _, v := range chain {
				solution += v
			}
			fmt.Printf("Completed the change: %v with total of %d\n", chain, solution)
			return true
		}
	}
	return false
}

func Challenge61() {
	triangleList := makePolygonalList(1000, 9999, triangle)
	squareList := makePolygonalList(1000, 9999, square)
	pentagonList := makePolygonalList(1000, 9999, pentagon)
	hexagonList := makePolygonalList(1000, 9999, hexagon)
	heptagonList := makePolygonalList(1000, 9999, heptagon)
	octagonList := makePolygonalList(1000, 9999, octagon)

	// triangleList doesn't go into the list of pointers to lists because we use triangleList as our
	// starting point.  We have no matches to compare at this point so we need to consider every
	// triangular number as a potential solution.  We could have started with any of the other pointers,
	// but just happened to choose triangleList as the starting point.
	bigList := []*[]int{&squareList, &pentagonList, &hexagonList, &heptagonList, &octagonList}
	excludeList := []*[]int{}
	for _, v := range triangleList {
		if linkTogether(bigList, excludeList, []int{v}) {
			return
		}
	}
}
