package challenge54

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Hand struct {
	cards        map[byte]int
	flush        byte
	straight     byte
	pairs        []byte
	triple       []byte
	quad         byte
	singles      []byte
	hasMultiples bool
}

func (h Hand) print() {
	fmt.Printf("=========================\n")
	fmt.Printf("card values are: %v\n", h.cards)
	fmt.Printf("flush is: %d\n", h.flush)
	fmt.Printf("straight value is: %d\n", h.straight)
	fmt.Printf("singles are: %d\n", h.singles)
	fmt.Printf("pairs are: %d\n", h.pairs)
	fmt.Printf("triple is: %d\n", h.triple)
	fmt.Printf("quad is: %d\n", h.quad)
	fmt.Printf("has multiples is: %v\n", h.hasMultiples)
}

const (
	unInitialised = iota
	nothing
	onePair
	twoPairs
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
	royalFlush
)

func (h Hand) score() []byte {
	if h.flush != 0 {
		if h.straight != 0 {
			if h.cards['A'] != 0 {
				fmt.Printf("Royal Flush!\n")
				return []byte{royalFlush, h.flush}
			}
			fmt.Printf("Straight Flush\n")
			return []byte{straightFlush, h.straight, h.flush}
		}
		fmt.Printf("Flush\n")
		return []byte{flush, h.flush}
	}
	if h.straight != 0 {
		fmt.Printf("Straight\n")
		return []byte{straight, h.straight}
	}
	if h.quad != 0 {
		fmt.Printf("Four of a Kind\n")
		return []byte{fourOfAKind, h.quad, h.singles[0]}
	}
	if h.triple != nil {
		if len(h.pairs) == 1 {
			fmt.Printf("Full House\n")
			return []byte{fullHouse, h.triple[0], h.pairs[0]}
		}
		fmt.Printf("Three of a Kind\n")
		return append([]byte{threeOfAKind}, h.singles...)
	}
	if len(h.pairs) == 2 {
		fmt.Printf("Two Pairs ")
		return []byte{twoPairs, h.pairs[0], h.pairs[1], h.singles[0]}
	}
	if len(h.pairs) == 1 {
		fmt.Printf("Pair ")
		return append([]byte{onePair, h.pairs[0]}, h.singles...)
	}
	if len(h.pairs) == 0 {
		fmt.Printf("Nothing\n")
		return append([]byte{nothing}, h.singles...)
	}
	fmt.Printf("Shouldn't reach here\n")
	return []byte{0}

}

func cardToValue(card byte) byte {
	if card >= '2' && card <= '9' {
		return card - '2'
	}
	if card == 'T' {
		return 8
	}
	if card == 'J' {
		return 9
	}
	if card == 'Q' {
		return 10
	}
	if card == 'K' {
		return 11
	}
	return 12
}

func isStraight(cards []string) byte {
	lowest := byte(12)
	highest := byte(0)

	for _, card := range cards {
		v := cardToValue(card[0])
		// fmt.Printf("v: %d\n", v)
		if v < lowest {
			// fmt.Printf("setting lowest to %d\n", v)
			lowest = v
		}
		if v > highest {
			// fmt.Printf("setting highest to %d\n", v)
			highest = v
		}

	}
	// fmt.Printf("highest is: %d; lowest is: %d\n", highest, lowest)
	if highest-lowest == 4 {
		return byte(highest)
	}
	return 0
}
func newHand(cards []string) Hand {
	var h Hand
	h.cards = make(map[byte]int)
	firstCardSuit := cards[0][1]
	h.flush = firstCardSuit
	for _, card := range cards {
		val := card[0]
		suit := card[1]
		if suit != firstCardSuit {
			h.flush = 0
		}
		h.cards[val]++
	}
	for k, v := range h.cards {
		if v == 4 {
			h.quad = cardToValue(k)
			h.hasMultiples = true
			continue
		}
		if v == 3 {
			h.triple = []byte{cardToValue(k)}
			h.hasMultiples = true
		}
		if v == 2 {
			h.pairs = append(h.pairs, cardToValue(k))
			h.hasMultiples = true
		}
		if v == 1 {
			h.singles = append(h.singles, cardToValue(k))
		}
	}
	if !h.hasMultiples {
		h.straight = isStraight(cards)
	}
	sort.Slice(h.singles, func(i int, j int) bool {
		return h.singles[i] > h.singles[j]
	})
	sort.Slice(h.pairs, func(i int, j int) bool {
		return h.pairs[i] > h.pairs[j]
	})
	return h

}
func isFirstScoreGreater(first []byte, second []byte) bool {
	if first[0] == second[0] {
		return isFirstScoreGreater(first[1:], second[1:])
	}
	if first[0] > second[0] {
		return true
	} else {
		return false
	}
}

func Challenge54() {

	var solution int
	var solutionComplement int
	f, err := os.Open("p054_poker.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 1
	for scanner.Scan() {
		fmt.Printf("----------------- line %d --------------------\n", count)
		count++
		tenCards := strings.Fields(scanner.Text())
		h1 := newHand(tenCards[:5])
		h2 := newHand(tenCards[5:])

		score1 := h1.score()
		score2 := h2.score()
		fmt.Printf("score1: %v\n", score1)
		fmt.Printf("score2: %v\n", score2)
		h1.print()
		h2.print()

		if isFirstScoreGreater(score1, score2) {
			solution++
			fmt.Printf("%d hand 1 wins\n", count-1)
		} else {
			solutionComplement++
			fmt.Printf("%d hand 2 wins\n", count-1)
		}

		fmt.Printf("Challenge 54 solution is: %d\n", solution)
		//		fmt.Printf("Solution complement is: %d\n", solutionComplement)

		// fmt.Printf("H1:\n")
		// h1.print()
		// fmt.Printf("H2:\n")
		// h2.print()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// var h Hand
	// // Nothing
	// h = newHand([]string{"7D", "3S", "QD", "5D", "6D"})
	// h.print()
	// fmt.Printf("Score is %d\n", h.score())

	// // Pair
	// h = newHand([]string{"7D", "7S", "4D", "5D", "6D"})
	// h.print()
	// fmt.Printf("Score is %d\n", h.score())

	// // Two Pairs
	// h = newHand([]string{"7D", "3S", "3D", "5D", "5H"})
	// h.print()
	// fmt.Printf("Score is %d\n", h.score())

	// // Three of a Kind
	// h = newHand([]string{"7D", "7S", "4D", "5D", "7H"})
	// h.print()
	// fmt.Printf("Score is %d\n", h.score())

	// // Straight
	// h = newHand([]string{"TD", "9D", "6S", "7H", "8C"})
	// h.print()
	// fmt.Printf("Score is %d\n", h.score())

	// // Flush
	// h = newHand([]string{"TD", "3D", "4D", "5D", "6D"})
	// h.print()
	// fmt.Printf("Score is %d\n", h.score())

	// // Full House
	// h = newHand([]string{"7D", "3S", "7C", "3C", "3H"})
	// h.print()
	// fmt.Printf("Score is %d\n", h.score())

	// // Four of a Kind
	// h = newHand([]string{"7S", "5C", "7D", "7H", "7C"})
	// h.print()
	// fmt.Printf("Score is %d\n", h.score())

	// // Straight Flush
	// h = newHand([]string{"7D", "3D", "4D", "5D", "6D"})
	// h.print()
	// fmt.Printf("Score is %d\n", h.score())

	// // Royal Flush
	// h = newHand([]string{"AD", "KD", "TD", "JD", "QD"})
	// h.print()
	// fmt.Printf("Score is %d\n", h.score())
}
