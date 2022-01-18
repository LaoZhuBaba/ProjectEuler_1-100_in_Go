package challenge79

import (
	"bufio"
	"fmt"
	"os"
)

func isTripletInString(t [3]rune, s string) bool {
	// Determine whether a triplet of 3 runes occurs in a string in order.
	// E.g., [1 2 3] would be true for "641818253".
	// This is true because the string contains a 3 which occurs after a 2 which occurs after a 1.
	i := 0
	for _, v := range s {
		if v == t[i] {
			i++
			if i == 3 {
				return true
			}
		}
	}
	return false
}

func Challenge79() {
	s := make([][]rune, 0)
	f, err := os.Open("p079_keylog.txt")
	if err != nil {
		fmt.Printf("Failed to open file\n")
	} else {
		defer f.Close()
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s = append(s, []rune(scanner.Text()))
	}
	// Starting at 1 is obviously wasteful because no triple could return true for a one or two digit number.
	// However, I have to start somewhere and once you look at the data it is pretty obvious that the solution
	// must be quite a lot higher then 100.  If I attempted to work out "by hand" what the number to start
	// counting from I would at least partly have solved the challenge, so just start at 1 and let the
	// code find the answer.
	for count := 1; ; count++ {
		broken := false
		for _, v := range s {
			if !isTripletInString([3]rune{v[0], v[1], v[2]}, fmt.Sprintf("%d", count)) {
				// As soon as isTripletInString returns false for any triplet we raise a flag and break
				broken = true
				break
			}
		}
		// If we reach here without the broken flag being raised then we must have found out solution
		if !broken {
			fmt.Printf("Challenge 79 solution is: %d\n", count)
			return
		}
	}
}
