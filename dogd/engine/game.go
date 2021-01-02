package engine

import (
	"fmt"
	engine "github.com/trichner/dog/dogd/engine/board"
	"github.com/trichner/dog/dogd/engine/deck"
	rnd "github.com/trichner/dog/dogd/engine/rand"
	"math/rand"
)

type Player struct {
	color engine.Color
	cards *deck.Deck
}

type Game struct {
	board   *engine.Board
	players []Player
	cards   *deck.Deck
	turn    int
	random  *rand.Rand
}

func NewGame() *Game {
	random, err := rnd.NewRandom([]byte{1, 3, 3, 7})
	if err != nil {
		panic(err)
	}

	board := engine.NewBoard()
	players := newPlayers()
	deckOfCards := newShuffledDoubleDeck(random)

	return &Game{
		board:   board,
		players: players,
		cards:   deckOfCards,
		random:  random,
		turn:    0,
	}
}

func (g *Game) Move(color engine.Color, card deck.Card, move MarbleMove) error {
	var err error

	err = g.playCard(color, card)
	if err != nil {
		return err
	}

}

func (g *Game) playCard(color engine.Color, card deck.Card) error {
	player := g.players[color]

	err := player.cards.DiscardFirst(card)
	if err != nil {
		return fmt.Errorf("player %s cannot discard %s: %w", color, card, err)
	}
}

func newShuffledDoubleDeck(r *rand.Rand) *deck.Deck {
	cards := append(
		deck.NewFullDeck().Cards(),
		deck.NewFullDeck().Cards()...,
	)
	shuffled := deck.Shuffle(r, cards)
	return deck.NewDeck(shuffled)
}

func newPlayers() []Player {
	var players []Player
	for _, c := range engine.ColorValues() {
		players = append(players, Player{
			color: c,
			cards: nil,
		})
	}
	return players
}
