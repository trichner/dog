package deck

import (
	"fmt"
	"math/rand"
	"strings"
)

type Deck struct {
	cards []Card
}

func NewFullDeck() *Deck {

	cards := make([]Card, 0, 52+4)

	cards = append(cards, allCardsOfSuit(SuitSpades)...)
	cards = append(cards, allCardsOfSuit(SuitHearts)...)
	cards = append(cards, allCardsOfSuit(SuitDiamonds)...)
	cards = append(cards, allCardsOfSuit(SuiteClubs)...)

	nJokers := 4
	for i := 0; i < nJokers; i++ {
		cards = append(cards, kJoker)
	}

	return NewDeck(cards)
}

// NewDeck creates a new deck from a slice of cards
func NewDeck(cards []Card) *Deck {

	// make a defensive copy
	copied := make([]Card, len(cards))
	copy(copied, cards)
	return &Deck{cards: copied}
}

// DiscardFirst discards the first matching card from the deck while maintaining the order
func (d *Deck) DiscardFirst(card Card) error {
	idx := d.Find(card)
	if idx < 0 {
		return fmt.Errorf("cannot discard card, deck does not contain %v", card)
	}

	d.cards = append(d.cards[:idx], d.cards[idx+1:]...)
	return nil
}

func (d *Deck) Contains(card Card) bool {
	return d.Find(card) >= 0
}

// Find returns the index of the card in the deck or -1 if its not in the deck
func (d *Deck) Find(card Card) int {
	for i, c := range d.cards {
		if c == card {
			return i
		}
	}
	return -1
}

// DrawN draws N cards from the deck
func (d *Deck) DrawN(n int) ([]Card, error) {
	if len(d.cards) < n {
		return nil, fmt.Errorf("less than %d cards left on deck", n)
	}
	bisection := len(d.cards) - n
	drawn := d.cards[bisection:]
	d.cards = d.cards[:bisection]

	return drawn, nil
}

func (d *Deck) Size() int {
	return len(d.cards)
}

func (d *Deck) Cards() []Card {
	return d.cards
}

func (d *Deck) String() string {
	var sb strings.Builder
	for _, c := range d.cards {
		sb.WriteString(c.String())
	}

	return sb.String()
}

func Shuffle(rand *rand.Rand, cards []Card) []Card {
	shuffled := make([]Card, len(cards))
	copy(shuffled, cards)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
