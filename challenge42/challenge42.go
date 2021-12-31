// The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:

// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

// By converting each letter in a word to a number corresponding to its alphabetical position and adding these values
// we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle
// number then we shall call the word a triangle word.

// Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common
// English words, how many are triangle words?

package challenge42

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func isTriangular(n int) bool {
	intSqrRoot := int(math.Sqrt(float64(2 * n)))
	return intSqrRoot*(intSqrRoot+1)/2 == n
}
func Challenge42() {

	//words := make([]string, 0)
	f, err := os.Open("p042_words.txt")
	if err != nil {
		fmt.Printf("%d\n", err)
		os.Exit(-1)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var solution int

	for {
		word, err := r.ReadString(',')
		word = strings.TrimRight(word, ",")
		word = strings.Trim(word, "\"")
		// fmt.Printf("Name is: %s\n", word)
		// words = append(words, word)
		if err == io.EOF {
			break
		}
		// fmt.Printf("%v\n", words)
		var wordValue int
		for _, r := range word {
			wordValue += int(r - 'A' + 1)
		}
		if isTriangular(wordValue) {
			solution++
		}

	}
	fmt.Printf("Challenge 42 solution is: %d\n", solution)
}
