package challenge89

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Token struct {
	value int
	count int
}

type TokenSlice []Token

type RomanCount struct {
	symbol string
	count  int
	prefix string
}

func (rc RomanCount) String() string {
	if rc.prefix != "" {
		return strings.Repeat(rc.symbol, rc.count-1) + rc.prefix + rc.symbol
	} else {
		return strings.Repeat(rc.symbol, rc.count)
	}
}

func newToken(r rune, count int) (Token, error) {
	var t Token
	switch r {
	case 'i', 'I':
		if count > 9 {
			return t, errors.New("too many consecutive 'I' characters")
		}
		t.value = 1
		t.count = count
	case 'v', 'V':
		if count > 1 {
			return t, errors.New("consecutive 'V' characters are not valid")
		}
		t.value = 5
		t.count = count
	case 'x', 'X':
		if count > 9 {
			return t, errors.New("too many consecutive 'V' characters")
		}
		t.value = 10
		t.count = count
	case 'l', 'L':
		if count > 1 {
			return t, errors.New("consecutive 'L' characters are not valid")
		}
		t.value = 50
		t.count = count
	case 'c', 'C':
		if count > 9 {
			return t, errors.New("too many consecutive 'C' characters")
		}
		t.value = 100
		t.count = count
	case 'd', 'D':
		if count > 1 {
			return t, errors.New("consecutive 'D' characters are not valid")
		}
		t.value = 500
		t.count = count
	case 'm', 'M':
		t.value = 1000
		t.count = count
	default:
		return t, errors.New("invalid character or can't parse")
	}
	return t, nil
}
func getToken(runes []rune) (Token, error) {
	count := 1
	for count < len(runes) && unicode.ToUpper(runes[count]) == unicode.ToUpper(runes[0]) {
		count++
	}
	return newToken(runes[0], count)
}
func romNumTokenise(s string, logger *log.Logger) TokenSlice {
	var count int
	var tokens TokenSlice
	runes := []rune(s)
	for count < len(runes) {
		tok, err := getToken(runes[count:])
		if err != nil {
			logger.Fatal(err)
		}

		count += tok.count
		tokens = append(tokens, tok)
	}
	return tokens
}

func (t Token) tokenToInteger() int {
	return t.value * t.count
}

func decToRoman(n int) string {

	thousands := RomanCount{count: n / 1000, symbol: "M"}
	five100s := RomanCount{count: n % 1000 / 500, symbol: "D"}
	hundreds := RomanCount{count: n % 500 / 100, symbol: "C"}
	fifties := RomanCount{count: n % 100 / 50, symbol: "L"}
	tens := RomanCount{count: n % 50 / 10, symbol: "X"}
	fives := RomanCount{count: n % 10 / 5, symbol: "V"}
	units := RomanCount{count: n % 5, symbol: "I"}
	if hundreds.count == 4 {
		if five100s.count == 1 {
			five100s.count = 0
			thousands.count++
			thousands.prefix = "C"
		} else {
			five100s.count = 1
			five100s.prefix = "C"
		}
		hundreds.count = 0
	}
	if tens.count == 4 {
		if fifties.count == 1 {
			fifties.count = 0
			hundreds.count++
			hundreds.prefix = "X"
		} else {
			fifties.count = 1
			fifties.prefix = "X"
		}
		tens.count = 0
	}
	if units.count == 4 {
		if fives.count == 1 {
			fives.count = 0
			tens.count++
			tens.prefix = "I"

		} else {
			fives.count = 1
			fives.prefix = "I"
		}
		units.count = 0
	}
	return thousands.String() + five100s.String() + hundreds.String() + fifties.String() + tens.String() + fives.String() + units.String()
}

//func parseTokens(tokens TokenSlice, logger *log.Logger) int {
func (tokens TokenSlice) parseTokens(logger *log.Logger) int {
	// Iterate from first token to second-to-last token.  This means it is always
	// look-ahead to the next token.  But obviously there will be one extra token
	// to handle at the end.
	var totalParseSum int
	var subtractor int
	last := len(tokens) - 1
	for count := 0; count < last; count++ {
		// Detect the case where a lower value token with a count of 1 precedes a
		// higher valued token. E.g., IV, XL, CM.  In this case tokens[count] is
		// saved as a subtractor.
		thisValue := tokens[count].value
		nextValue := tokens[count+1].value
		if thisValue < nextValue {
			if last-count >= 2 {
				if tokens[count+2].value == tokens[count].value {
					logger.Fatal("Invalid: if a value has a preceding subtractor the following value must be less than the subtractor")
				}
			}
			if tokens[count].count != 1 {
				logger.Fatal("subtractor token count can only be 1")
			}
			switch thisValue {
			case 1:
				if nextValue == 5 || nextValue == 10 {
					subtractor = tokens[count].tokenToInteger()
				} else {
					logger.Fatal("'I' can only be subtractor preceding 'V' or 'X'!")
				}
			case 10:
				if nextValue == 50 || nextValue == 100 {
					subtractor = tokens[count].tokenToInteger()
				} else {
					logger.Fatal("'X' can only be subtractor before 'L' or 'C'!")
				}
			case 100:
				if nextValue == 500 || nextValue == 1000 {
					subtractor = tokens[count].tokenToInteger()
				} else {
					logger.Fatal("'C' can only be subtractor before 'D' or 'M'!")
				}
			default:
				logger.Fatal("Cannot parse!")
			}
		} else {
			totalParseSum += tokens[count].tokenToInteger() - subtractor
			subtractor = 0
		}
	}
	totalParseSum += tokens[last].tokenToInteger() - subtractor
	return totalParseSum
}

func loadRoman(fileName string) (s []string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s': %w", fileName, err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	return s, err
}
func Challenge89() {
	var originalTotalStrLen, finalTotalStrLen int
	var max int
	logger := log.New(os.Stderr, "", log.Lshortfile)
	tokens := romNumTokenise("MM", logger)
	//logger.Printf("Tokens: %v\n", tokens)
	sum := tokens.parseTokens(logger)
	logger.Printf("%d\n", sum)
	s, err := loadRoman("p089_roman.txt")
	if err != nil {
		log.Fatal(err)
	}
	var grandTotal int
	for _, romNum := range s {
		originalTotalStrLen += len(romNum)
		tokens := romNumTokenise(romNum, logger)
		sum := tokens.parseTokens(logger)
		if sum > max {
			max = sum
		}
		simplifiedRoman := decToRoman(sum)
		finalTotalStrLen += len(simplifiedRoman)
		logger.Printf("%s: %s: %d\n", romNum, simplifiedRoman, sum)
		grandTotal += sum
	}
	logger.Printf("Grand Total: %d (max: %d)\n", grandTotal, max)
	logger.Printf("Challenge 89 solution: %d\n", originalTotalStrLen-finalTotalStrLen)
}
