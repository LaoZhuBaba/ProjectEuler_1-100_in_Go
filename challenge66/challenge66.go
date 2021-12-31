package challenge66

import "fmt"

const max = 1

func increasingPairs() {
	for count1 := 1; count1 < max; count1++ {
		for count2 := 1; count2 <= count1; count2++ {
			fmt.Printf("%d %d\n", count1, count1)
			for ; count2 < count1; count2++ {
				fmt.Printf("%d %d\n", count1, count2)
				fmt.Printf("%d %d\n", count2, count1)
			}
		}

	}
}

type Squares struct {
	squareRootCache map[int]int
}

//var squares = &Squares{cache : make(map[int]int) }
func newSquares() *Squares {
	return &Squares{squareRootCache: make(map[int]int)}
}
func (s Squares) squareRoot(i int) (int, bool) {
	if sr, ok := s.squareRootCache[i]; ok {
		return sr, sr != 0
	} else {
		for n, square := 1, 1; square < i; n++ {
			square = n * n
			if _, ok := s.squareRootCache[square]; !ok {
				s.squareRootCache[square] = n
			}
			if square == i {
				return n, true
			}
		}
	}
	// Negative cache
	s.squareRootCache[i] = 0
	return 0, false
}
func Challenge66() {
	increasingPairs()
	squares := newSquares()
	for count := 20; count > 1; count -= 2 {
		root, isSquare := squares.squareRoot((count))
		fmt.Printf("count down: %d %d %v\n", count, root, isSquare)
	}
	fmt.Println("------------------")
	for count := 1; count < 30; count++ {
		root, isSquare := squares.squareRoot((count))
		fmt.Printf("count: %d %d %v\n", count, root, isSquare)
	}
}
