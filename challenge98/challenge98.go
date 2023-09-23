package challenge98

import (
	"encoding/csv"
	"errors"
	"euler/shared"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

const File = "0098_words.txt"

type squarePair struct {
	n       int
	nsquare int
}

type squareMap map[string][]squarePair
type transformSet struct {
	square1   int
	square2   int
	transform []int
}

type tfSetMap map[string][]transformSet

// wordToSortedString takes a string and returns it in alphabetical order.
// E.g., "base" -> "abes"
func wordToSortedString(word string) (s string) {
	var strSlice []string
	for _, v := range word {
		strSlice = append(strSlice, string(v))
	}
	sort.Strings(strSlice)
	for _, v := range strSlice {
		s += v
	}
	return s
}

// intSliceToString converts a slice of integers to a string.
// E.g., []int{2,6,5} -> "265"
func intSliceToString(is []int) (s string) {
	for _, v := range is {
		s += fmt.Sprintf("%d", v)
	}
	return s
}

// intToSortedString takes an integer and returns a string
// which an ordered string of digits in the base10 representation
// of that integer.  E.g., 8398 -> "3889".  A second return value
// is the number of unique digits in the strings so 8389 would
// return 3 as its second value.
func intToSortedString(n int) (s string, numUniq int) {
	m := make(map[int]bool)
	islice := shared.IntToIntSlice(n)
	for _, v := range islice {
		m[v] = true
	}
	sort.Ints(islice)
	for _, v := range islice {
		s += fmt.Sprintf("%d", v)
	}
	return s, len(m)
}

// countUniqueChars returns the number of unique characters in a string.
// E.g., "apple" -> 4 and "dodo" -> 2
func countUniqueChars(s string) int {
	m := make(map[rune]bool)
	for _, c := range s {
		m[c] = true
	}
	return len(m)
}

func loadFile() (all []string) {
	f, err := os.Open(File)
	if err != nil {
		log.Fatalf("failed to open file: %s", File)
	}
	all, err = csv.NewReader(f).Read()
	if err != nil {
		log.Fatalf("failed to read CSV file: %s", File)
	}
	return all
}

type anagram struct {
	//	sorted       string
	ana1 []rune
	ana2 []rune
	//	transform    []int
	transformStr string
}

// getSquaresByNumDigits takes a number of base 10 digits and returns the lowest
// and highest integers which have squares with that number of digits.  E.g., if n is 2
// then min will 4 and max will be 9.
func getSquaresByNumDigits(n int) (min, max int) {
	nDigitNumMin := math.Round(math.Pow10(n - 1))
	nDigitNumMax := math.Round(nDigitNumMin*10 - 1)
	min = int(math.Sqrt(nDigitNumMin-1)) + 1
	max = int(math.Sqrt(nDigitNumMax - 1))
	return min, max
}

// makeTransform takes two slices of runes and returns a slice of integers
// which represent how to transform one to another.  E.g., the transform
// "tool" -> "loot" is has a transform [3,1,2,0] and "stop" -> "post" has
// a transorm of [2,3,1,0]
func makeTransform(ana1, ana2 []rune) (ret []int, err error) {
	if len(ana1) != len(ana2) {
		return nil, errors.New("in makeTransform, incompatible strings passed")
	}
	ret = make([]int, len(ana1))

outer:
	for k1, v1 := range ana1 {
		for k2, v2 := range ana2 {
			if v2 == v1 {
				ret[k1] = k2
				continue outer
			}
		}
		return nil, errors.New("in makeTransform, incompatible strings passed")
	}
	return ret, nil
}

// func makeTransformForSquares is similar to makeTransform() but it takes
// two integers which converts to strings before passing to makeTransform
func makeTransformForSquares(n1, n2 int) (ret []int, err error) {
	nSlice1 := shared.IntToIntSlice(n1)
	nSlice2 := shared.IntToIntSlice(n2)
	if len(nSlice1) != len(nSlice2) {
		return nil, errors.New("in makeTransformForSquares, incompatible values passed")
	}
	runeSlice1 := make([]rune, len(nSlice1))
	runeSlice2 := make([]rune, len(nSlice2))
	for idx := range nSlice1 {
		runeSlice1[idx] = '0' + int32(nSlice1[idx])
		runeSlice2[idx] = '0' + int32(nSlice2[idx])
	}
	return makeTransform(runeSlice1, runeSlice2)
}

// findAnagrams takes a slice of strings and returns a map
// where each value is a slice of strings which are anagrams
// of one another and the key is the a string made up of the
// same characters as the anagrams but in sorted sorte order.
// For example one map element might be:
//
//	{"opst": ["stop", "post", "pots"]}
func findAnagrams(words []string) (m map[string][]string) {
	m = make(map[string][]string)
	for _, word := range words {
		key := wordToSortedString(word)
		if _, ok := m[key]; !ok {
			m[key] = []string{}
		}
		m[key] = append(m[key], word)
	}
	for k, v := range m {
		if len(v) < 2 {
			delete(m, k)
		}
	}
	return m
}

// getPairsInSet is a utility function which returns all the two member
// subsets of a set of a given size.  So getPairsInSet(3) returns:
// [[0,1], [0,2], [1,0], [1,2], [2,0], [2,1]]
func getPairsInSet(n int) (pairs [][2]int) {
	for count1 := 0; count1 < n; count1++ {
		for count2 := 0; count2 < n; count2++ {
			if count2 == count1 {
				continue
			}
			pairs = append(pairs, [2]int{count1, count2})
		}
	}
	return pairs
}

// makeTransformSetMap takes a map containing pairs of square numbmers
// and returns a map of anagram square numbers keyed on a string representing
// a transform from one to the other.  For example, the anagram pair of square
// numbers 1296 & 9216 will be stored together in a slice under the key
// "2103" where "2103" transforms 1296 -> 9216.
func makeTransformSetMap(sm squareMap) (transformSetMap tfSetMap) {
	transformSetMap = tfSetMap{}
	for _, v := range sm {
		combinations := getPairsInSet(len(v))
		for _, combination := range combinations {
			square1 := v[combination[0]].nsquare
			square2 := v[combination[1]].nsquare
			transform, err := makeTransformForSquares(square1, square2)
			transformStr := intSliceToString(transform)
			if err != nil {
				log.Fatalf("in makeSquareTransformList: failed to calculate transform from %d & %d", square1, square2)
			}
			if _, ok := transformSetMap[transformStr]; !ok {
				transformSetMap[transformStr] = []transformSet{}
			}
			transformSetMap[transformStr] = append(
				transformSetMap[transformStr], transformSet{
					square1:   square1,
					square2:   square2,
					transform: transform,
				})
		}
	}
	return transformSetMap
}

func Challenge98() {
	// keep track of the longest word
	var longest int

	// Load all the words
	words := loadFile()

	// Find the anagrams
	anaMap := findAnagrams((words))

	// We know from looking at the data that none of the anagrams have repeated letters.
	// And the because each leter must correspond with a digit we can discard any anagrams
	// that are longer than 10 characters.
	for k, v := range anaMap {
		if len(v) > 10 {
			delete(anaMap, k)
		}
	}

	// An anagram stores a bunch of fields associated with the anagram pair
	var anaSlice []anagram
	for _, v := range anaMap {
		for k2, a2 := range v {
			for k3, a3 := range v {
				// Compare every anagram to every other anagram of the same letters.
				// (There may be > 2, e.g., "pots", "stop", "post").  The continue
				// is to avoid compare an anagram to itself.
				if k3 == k2 {
					continue
				}
				// Convert the anagram strings to rune slices
				r2 := []rune(a2)
				r3 := []rune(a3)
				// Create a transform which maps one anagram to another.  E.g.,
				// the transform from "stop" to "post" is []int{2,3,1,0}.  Each
				// integer tells you where an element moves to in the transform.
				// So the first element in the transform is 2 because "s" moves
				// to index 2 in the transformed string.
				trans, err := makeTransform(r2, r3)
				if err != nil {
					log.Fatalf("couldn't create transform for %v & %v", r2, r3)
				}
				// To solve the challenge we only actually need transformStr but it's
				// nice to be able to show the anagrams which generate matching transform
				// for evidentiary purposes.
				anaSlice = append(
					anaSlice,
					anagram{
						ana1:         r2,
						ana2:         r3,
						transformStr: intSliceToString(trans),
					},
				)
			}
		}
	}
	for word := range anaMap {
		if len(word) > longest {
			longest = len(word)
		}
	}

	fmt.Printf("Starting with %d...\n", longest)

	for count := longest; count > 1; count-- {

		fmt.Printf("Searching for a solution with words of length %d\n", count)
		squareMin, squareMax := getSquaresByNumDigits(count)

		// m is a map storing square numbers keyed on a string which is
		// a sorted string of digits.  For example the square numbers
		// 1296 and 9216 will both be stored under key "1269".  This is
		// an easy way of identifying numbers which are anagrams of one
		// another
		m := make(squareMap)

		// Step through all the square numbers a certain digit length
		for count := squareMin; count <= squareMax; count++ {
			sqStr, numUniq := intToSortedString(count * count)
			// ignore numbers with repeated digits because none of our anagrams
			// contain repeated characters.
			if numUniq != len(sqStr) {
				continue
			}
			// Create a map entry if none exists
			if _, ok := m[sqStr]; !ok {
				m[sqStr] = []squarePair{}
			}
			// append to the matching map entry
			m[sqStr] = append(m[sqStr], squarePair{n: count, nsquare: count * count})
		}

		// If a map key only stores one number then there are no anagrams so discard
		for k, v := range m {
			if len(v) == 1 {
				delete(m, k)
			}
		}

		// Reorganise the data contained in m so that each possible pair of square numbers
		// which are anagrams of each other are stored in a map which is keyed on a string
		// that represents the transforom from ono to another.  For example under the the
		// parif of square numbers 1296 & 9216 will be stored under the key "2103".  At
		// this point we have a transform string for all the square number anagram pairs
		// and all the anagram words, so we can compare the transform strings and find a
		// match.
		transportSetMap := makeTransformSetMap(m)

		for _, av := range anaSlice {
			if ts, ok := transportSetMap[av.transformStr]; ok {
				fmt.Printf("found: %v %c %c\n", ts, av.ana1, av.ana2)
				var solution int
				for _, v := range ts {
					if v.square1 > solution {
						solution = v.square1
					}
					if v.square2 > solution {
						solution = v.square2
					}
				}
				fmt.Printf("challenge 98 solution is: %d with words %s & %s\n", solution, string(av.ana1), string(av.ana2))
				return
			}
		}
	}
}
