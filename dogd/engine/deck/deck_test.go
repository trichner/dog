package deck

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestDeck_DrawN(t *testing.T) {

	deck := NewFullDeck()
	fullSize := deck.Size()

	n := 42
	cards, _ := deck.DrawN(n)

	assert.Len(t, cards, n)
	assert.Len(t, deck.cards, fullSize-n)
}

func TestDeck_Shuffle(t *testing.T) {

	cards := NewFullDeck().Cards()

	rng := rand.New(rand.NewSource(42))
	for i := 0; i < 10000; i++ {
		shuffled := Shuffle(rng, cards)
		assertCompleteDeck(t, NewDeck(shuffled))
		assert.NotEqual(t, cards, shuffled, "should veeery likely no longer be equal")
	}
}

func TestDeck_NewFullDeck(t *testing.T) {

	deck := NewFullDeck()

	assertCompleteDeck(t, deck)
}

func assertCompleteDeck(t *testing.T, deck *Deck) {
	assert.Equal(t, 56, deck.Size())

	assertContainsFullSuit(t, deck, SuitSpades)
	assertContainsFullSuit(t, deck, SuitHearts)
	assertContainsFullSuit(t, deck, SuitDiamonds)
	assertContainsFullSuit(t, deck, SuiteClubs)

	assert.Equal(t, 4, countRank(deck.cards, RankJoker))
}

func assertContainsFullSuit(t *testing.T, deck *Deck, suit Suit) {
	cards := filterSuit(deck, suit)

	ranks := []Rank{RankAce, RankTwo, RankThree, RankFour, RankFive, RankSix, RankSeven, RankEight, RankNine, RankTen, RankJack, RankQueen, RankKing}
	assert.Len(t, cards, len(ranks), "suit %d should have %d entries but has %d", suit, len(ranks), len(cards))

	for _, r := range ranks {
		assert.Equal(t, 1, countRank(cards, r))
	}
}

func countRank(cards []Card, r Rank) int {
	n := 0
	for _, c := range cards {
		if c.rank == r {
			n++
		}
	}
	return n
}

func filterSuit(deck *Deck, suit Suit) []Card {
	cards := []Card{}
	for _, c := range deck.cards {
		if c.suit == suit {
			cards = append(cards, c)
		}
	}
	return cards
}
