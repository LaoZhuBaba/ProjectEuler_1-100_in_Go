package challenge84

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Type definitions

// TokenMover allows us to move tokens in various ways using the same type
type TokenMover interface {
	TokenMove(*Board, string)
}

type Regress int       // For backtracking a token a certain number of squares.
type GotoSquare string // Move a token to a named square
type SkipToStr string  // Fast forward a token to the next square starting with string

// Represents a set of dice sizes gives the size(s) of the dice and values stores their current value
type Dice struct {
	sizes  []int
	values []int
	random *rand.Rand
}

// Either a Chance or Community Chest card which may have a message and may redirect a token to another square
type Card struct {
	msg         string
	redirection TokenMover
}

// Just a deck of cards
type Deck struct {
	cards    []Card
	nextCard int
}

// The board itself
type Board struct {
	squares        []string       // a list of names for squares
	squaresByDescr map[string]int // look up the index of a square using its name
	tokens         map[string]int // look up the position of a token by its name
	ccDeck         Deck           // Community Chest deck
	chDeck         Deck           // Chance deck
	accounting     map[string]int // Keep a use count of every square
}

// Variable definitions

var monopolySquares = []string{
	"GO", "A1", "CC1", "A2", "T1", "R1", "B1", "CH1", "B2", "B3",
	"JAIL", "C1", "U1", "C2", "C3", "R2", "D1", "CC2", "D2", "D3",
	"FP", "E1", "CH2", "E2", "E3", "R3", "F1", "F2", "U2", "F3",
	"G2J", "G1", "G2", "CC3", "G3", "R4", "CH3", "H1", "T2", "H2",
}

var monopolyCommunityChestCards = []Card{
	{"Advance to Go", GotoSquare("GO")},
	{"Go to Jail", GotoSquare("JAIL")},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
}
var monopolyChanceCards = []Card{
	{"Advance to Go", GotoSquare("GO")},
	{"Go to Jail", GotoSquare("JAIL")},
	{"Generic", GotoSquare("C1")},
	{"Generic", GotoSquare("E3")},
	{"Generic", GotoSquare("H2")},
	{"Generic", GotoSquare("R1")},
	{"Generic", SkipToStr("R")},
	{"Generic", SkipToStr("R")},
	{"Generic", SkipToStr("U")},
	{"Generic", Regress(3)},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
	{"Generic", nil},
}

// We actually only need one token for this model.
var monopolyTokens = []string{"doggy", "ship", "boot"}

func (r Regress) TokenMove(b *Board, tokenStr string) {
	b.incrToken(tokenStr, int(r)*-1)
}
func (gs GotoSquare) TokenMove(b *Board, tokenStr string) {
	b.placeToken(tokenStr, string(gs))
}
func (sts SkipToStr) TokenMove(b *Board, tokenStr string) {

	for location := b.getTokenLocation(tokenStr); location[0:1] != string(sts); location = b.getTokenLocation(tokenStr) {
		// These increments are final, so don't account for them against the square.
		b.incrTokenWithoutAccounting(tokenStr, 1)
	}
	location := b.getTokenLocation(tokenStr)
	// Place the token in the location where it already is, just to update the accounting
	// for that square.
	b.placeToken(tokenStr, location)
}

// Initialiser functions which allow us to populate decks, dice and board correctly
func newDeck(cards []Card) Deck {
	c := make([]Card, 0)
	c = append(c, cards...)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	r.Shuffle(len(c), func(i, j int) { c[i], c[j] = c[j], c[i] })
	return Deck{cards: c, nextCard: 0}
}
func newDice(s ...int) *Dice {
	dice := new(Dice)
	dice.sizes = append(dice.sizes, s...)
	source := rand.NewSource(time.Now().UnixNano())
	dice.random = rand.New(source)
	return dice
}

func newBoard(squares []string, tokens []string) *Board {
	b := new(Board)
	b.squares = append(b.squares, squares...)
	b.tokens = make(map[string]int)
	b.accounting = make(map[string]int)
	for _, token := range tokens {
		b.tokens[token] = 0
	}
	b.squaresByDescr = make(map[string]int)
	for i, s := range b.squares {
		b.squaresByDescr[s] = i
	}
	b.chDeck = newDeck(monopolyChanceCards)
	b.ccDeck = newDeck(monopolyCommunityChestCards)
	return b
}

func (d Dice) Roll() (int, bool) {
	var total int
	for _, size := range d.sizes {
		d.values = append(d.values, d.random.Int()%size+1)
	}
	unanimous := true
	firstDie := d.values[0]
	for _, value := range d.values {
		total += value
		if value != firstDie {
			unanimous = false
		}
	}
	return total, unanimous
}

func (b Board) incrToken(s string, i int) {
	// Deal with negative case.  In the unlikely event that i has a negative size greater
	// than len(b.squares) just keep adding len(b.squares until it becomes positive)

	for i < 0 {
		i += len(b.squares)
	}
	b.tokens[s] += i
	b.tokens[s] %= len(b.squares)
	location := b.getTokenLocation(s)
	b.accounting[location]++
}

func (b Board) incrTokenWithoutAccounting(s string, i int) {
	// Deal with negative case.  In the unlikely event that i has a negative size greater
	// than len(b.squares) just keep adding len(b.squares until it becomes positive)

	for i < 0 {
		i += len(b.squares)
	}
	b.tokens[s] += i
	b.tokens[s] %= len(b.squares)
}

func (b Board) placeToken(token string, location string) {
	b.tokens[token] = b.squaresByDescr[location]
	b.accounting[location]++
}

func (b Board) getTokenLocation(token string) string {
	return b.squares[b.tokens[token]]
}

func (b *Board) getCcCard() Card {
	next := b.ccDeck.nextCard
	fmt.Printf("nextCard is %d\n", b.ccDeck.nextCard)
	b.ccDeck.nextCard = (b.ccDeck.nextCard + 1) % len(b.ccDeck.cards)
	fmt.Printf("nextCard is now %d\n", b.ccDeck.nextCard)
	return b.ccDeck.cards[next]
}
func (b *Board) getChCard() Card {
	next := b.chDeck.nextCard
	fmt.Printf("nextCard is %d\n", b.chDeck.nextCard)
	b.chDeck.nextCard = (b.chDeck.nextCard + 1) % len(b.chDeck.cards)
	fmt.Printf("nextCard is now %d\n", b.chDeck.nextCard)
	return b.chDeck.cards[next]

}
func takeTurn(b *Board, d *Dice, tokenStr string) {
	var consecutiveDbls int
	for count := 1; count > 0; count-- {
		roll, dbl := d.Roll()
		var location string
		location = b.getTokenLocation(tokenStr)
		fmt.Printf("Location is %s.  Rolled dice: %d\n", location, roll)
		b.incrToken(tokenStr, roll)
		if dbl {
			fmt.Printf("Rolled a DOUBLE!\n")
			count++
			consecutiveDbls++
		} else {
			consecutiveDbls = 0
		}
		location = b.getTokenLocation(tokenStr)
		fmt.Printf("Moved to new location %s\n", location)
		if location == "G2J" || consecutiveDbls > 2 {
			fmt.Printf("Too many doubles!  Sent to Jail.  End of turn!\n")
			b.placeToken(tokenStr, "JAIL")
			return
		}
		if location[0:2] == "CH" {
			chCard := b.getChCard()
			if chCard.redirection != nil {
				chCard.redirection.TokenMove(b, tokenStr)
				location = b.getTokenLocation(tokenStr)
				fmt.Printf("Picked Chance Card with message '%s' which redirected to: %s\n", chCard.msg, location)
			} else {
				fmt.Printf("Picked Chance Card with message '%s' which has no redirection\n", chCard.msg)
			}
		}
		if location[0:2] == "CC" {
			ccCard := b.getCcCard()
			if ccCard.redirection != nil {
				ccCard.redirection.TokenMove(b, tokenStr)
				location = b.getTokenLocation(tokenStr)
				fmt.Printf("Picked CC Card with message '%s' which redirected to: %s\n", ccCard.msg, location)
			} else {
				fmt.Printf("Picked CC Card with message '%s' which has no redirection\n", ccCard.msg)
			}
		}
		location = b.getTokenLocation(tokenStr)
		if location == "G2J" || consecutiveDbls > 2 {
			b.placeToken(tokenStr, "JAIL")
			return
		}
		if !dbl {
			break
		}
	}
}
func Challenge84() {
	b := newBoard(monopolySquares, monopolyTokens)
	d := newDice(4, 4)
	// Simulate a number of turns with a single counter
	for count := 100_000; count > 0; count-- {
		takeTurn(b, d, "doggy")
	}

	// Sort the squares using the values in b.accounting
	sort.Slice(b.squares, func(i, j int) bool { return b.accounting[b.squares[j]] < b.accounting[b.squares[i]] })

	fmt.Printf("\nAccounting table: %v\n", b.accounting)
	fmt.Printf("\nSquares sorted by frequency: %v\n", b.squares)
	first := b.squaresByDescr[b.squares[0]]
	second := b.squaresByDescr[b.squares[1]]
	third := b.squaresByDescr[b.squares[2]]
	fmt.Printf("Challenge 84 solution is: %02d%02d%02d\n", first, second, third)

}
