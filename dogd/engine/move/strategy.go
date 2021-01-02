package move

import (
	"fmt"
	"github.com/trichner/dog/dogd/engine/board"
	"github.com/trichner/dog/dogd/engine/deck"
)

//go:generate enumer -type=MoveKind
type MoveKind int

const (
	KindForward MoveKind = iota
	KindReverse
	KindSwap
	KindStart
)

type Play struct {
	paths [][]board.NodeID
	kind  MoveKind

	card deck.Card // the card chosen for the move (i.e. king)
}

func (p *Play) String() string {
	return fmt.Sprintf("%s %s %s", p.card.String(), p.kind, p.paths)
}

func GenerateRankAcePlays(b *board.Board, card deck.Card, c board.Color) []Play {
	var plays []Play

	plays = append(plays, GenerateStartPlay(b, card, c)...)

	plays = append(plays, ForwardPlaysFactory(1)(b, card, c)...)
	plays = append(plays, ForwardPlaysFactory(11)(b, card, c)...)

	return plays
}

func GenerateRankKingPlays(b *board.Board, card deck.Card, c board.Color) []Play {
	var plays []Play

	plays = append(plays, GenerateStartPlay(b, card, c)...)
	plays = append(plays, ForwardPlaysFactory(13)(b, card, c)...)

	return plays
}

func GenerateRankQueenPlays(b *board.Board, card deck.Card, c board.Color) []Play {
	return ForwardPlaysFactory(12)(b, card, c)
}

func GenerateRankJackPlays(b *board.Board, card deck.Card, c board.Color) []Play {
	var plays []Play

	colors := board.ColorValues()
	for i := 0; i < 4; i++ {
		// at least one marble must be ours
		ourMarble := b.GetMarble(c, i)

		// can potentially swap with any other marble
		for _, otherColor := range colors {
			for j := 0; j < 4; j++ {
				if i == j && c == otherColor {
					// cannot swap itself
					continue
				}

				otherMarble := b.GetMarble(otherColor, j)
				plays = append(plays, GenerateSwapPlay(card, ourMarble, otherMarble)...)
			}
		}
	}

	return plays
}

func GenerateRankTenPlays(b *board.Board, card deck.Card, c board.Color) []Play {
	return ForwardPlaysFactory(10)(b, card, c)
}

func GenerateRankNinePlays(b *board.Board, card deck.Card, c board.Color) []Play {
	return ForwardPlaysFactory(9)(b, card, c)
}

func GenerateRankEightPlays(b *board.Board, card deck.Card, c board.Color) []Play {
	return ForwardPlaysFactory(8)(b, card, c)
}

func GenerateRankSevenPlays(b *board.Board, card deck.Card, c board.Color) []Play {
	fmt.Printf("not implemented :'(\n")
	return []Play{}
}

func GenerateRankSixPlays(b *board.Board, card deck.Card, c board.Color) []Play {
	return ForwardPlaysFactory(6)(b, card, c)
}

func GenerateRankFivePlays(b *board.Board, card deck.Card, c board.Color) []Play {
	return ForwardPlaysFactory(5)(b, card, c)
}

func GenerateRankFourPlays(b *board.Board, card deck.Card, c board.Color) []Play {
	var plays []Play
	plays = append(plays, ForwardPlaysFactory(4)(b, card, c)...)
	plays = append(plays, ReversePlaysFactory(4)(b, card, c)...)
	return plays
}

func GenerateRankThreePlays(b *board.Board, card deck.Card, c board.Color) []Play {
	return ForwardPlaysFactory(5)(b, card, c)
}

func GenerateRankTwoPlays(b *board.Board, card deck.Card, c board.Color) []Play {
	return ForwardPlaysFactory(5)(b, card, c)
}

func GenerateRankJokerPlays(b *board.Board, card deck.Card, c board.Color) []Play {

	var plays []Play
	plays = append(plays, GenerateRankAcePlays(b, card, c)...)
	plays = append(plays, GenerateRankKingPlays(b, card, c)...)
	plays = append(plays, GenerateRankQueenPlays(b, card, c)...)
	plays = append(plays, GenerateRankJackPlays(b, card, c)...)
	plays = append(plays, GenerateRankTenPlays(b, card, c)...)
	plays = append(plays, GenerateRankNinePlays(b, card, c)...)
	plays = append(plays, GenerateRankEightPlays(b, card, c)...)
	plays = append(plays, GenerateRankSevenPlays(b, card, c)...)
	plays = append(plays, GenerateRankSixPlays(b, card, c)...)
	plays = append(plays, GenerateRankFivePlays(b, card, c)...)
	plays = append(plays, GenerateRankFourPlays(b, card, c)...)
	plays = append(plays, GenerateRankThreePlays(b, card, c)...)
	plays = append(plays, GenerateRankTwoPlays(b, card, c)...)
	return plays
}

func GenerateSwapPlay(card deck.Card, our, other board.BoardMarble) []Play {

	if !isMarbleSwappable(our) || !isMarbleSwappable(other) {
		return []Play{}
	}

	from := our.Position().ID()
	to := other.Position().ID()

	return []Play{{
		paths: [][]board.NodeID{{from, to}},
		kind:  KindSwap,
		card:  card,
	}}
}

func isMarbleSwappable(m board.BoardMarble) bool {
	position := m.Position()
	if position.Kind() == board.KindFinish || position.Kind() == board.KindHome {
		return false
	}

	if !m.IsStartPassed() {
		return false
	}

	return true
}

func ForwardPlaysFactory(distance int) func(b *board.Board, card deck.Card, c board.Color) []Play {
	return func(b *board.Board, card deck.Card, c board.Color) []Play {
		var plays []Play

		for i := 0; i < 4; i++ {
			marble := b.GetMarble(c, i)
			plays = append(plays, GenerateMoveForwardPlays(card, marble, distance)...)
		}

		return plays
	}
}

func GenerateMoveForwardPlays(card deck.Card, marble board.BoardMarble, distance int) []Play {

	var plays []Play

	from := marble.Position()
	paths := FindForwardPaths(from, distance)
	for _, p := range paths {
		plays = append(plays, Play{
			paths: [][]board.NodeID{p},
			kind:  KindForward,
			card:  card,
		})
	}

	return plays
}

func ReversePlaysFactory(distance int) func(b *board.Board, card deck.Card, c board.Color) []Play {
	return func(b *board.Board, card deck.Card, c board.Color) []Play {
		var plays []Play

		for i := 0; i < 4; i++ {
			marble := b.GetMarble(c, i)
			plays = append(plays, GenerateMoveReversePlays(card, marble, distance)...)
		}

		return plays
	}
}

func GenerateMoveReversePlays(card deck.Card, marble board.BoardMarble, distance int) []Play {

	var plays []Play

	from := marble.Position()
	paths := FindReversePaths(from, distance)
	for _, p := range paths {
		plays = append(plays, Play{
			paths: [][]board.NodeID{p},
			kind:  KindReverse,
			card:  card,
		})
	}

	return plays
}

// GenerateStartPlay returns a play to move the first available marble of the color
// from home to the start
func GenerateStartPlay(b *board.Board, card deck.Card, c board.Color) []Play {

	for i := 0; i < 4; i++ {
		m := b.GetMarble(c, i)
		if m.Position().Kind() == board.KindHome {
			from := m.Position().ID()
			toNode := m.Position().Next()[0]
			occupant := toNode.Marble()
			if occupant != nil && !occupant.IsStartPassed() {
				// start is blocked, no play possible
				return []Play{}
			}

			to := toNode.ID()

			return []Play{{
				paths: [][]board.NodeID{{from, to}},
				kind:  KindStart,
				card:  card,
			}}
		}
	}

	return []Play{}
}
