// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names,
// begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value
// by its alphabetical position in the list to obtain a name score.
// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th
// name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

// What is the total of all the name scores in the file?

package challenges

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func alphabeticValue(s string) int {
	sum := 0
	for _, r := range s {
		sum = sum + int(r-'A'+1)
	}
	return sum
}

func Challenge22() {

	names := make([]string, 0)
	f, err := os.Open("p022_names.txt")
	if err != nil {
		fmt.Printf("%d\n", err)
		os.Exit(-1)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		name, err := r.ReadString(',')
		name = strings.TrimRight(name, ",")
		name = strings.Trim(name, "\"")
		// fmt.Printf("Name is: %s\n", name)
		names = append(names, name)
		if err == io.EOF {
			break
		}

	}
	sort.Strings(names)
	nameScore := 0
	sumTotal := 0
	for count, name := range names {
		alphaValue := alphabeticValue(name)
		nameScore = alphaValue * (count + 1)
		sumTotal += nameScore
		fmt.Printf("Name: %s has alphabetical value of: %d count of %d and nameScore of %d\n", name, alphaValue, count+1, nameScore)
		// fmt.Printf("%s\n", name)
	}
	fmt.Printf("Sum total is: %d\n", sumTotal)
}
