package move

import (
	"fmt"
	"github.com/trichner/dog/dogd/engine/board"
	"github.com/trichner/dog/dogd/engine/deck"
	"testing"
)

func TestGenerateRankJokerPlays(t *testing.T) {
	aBoard := board.NewBoard()

	aBoard.GetNode("GN14").SetMarble(board.NewMarble(board.ColorGreen, 1, true))
	aBoard.GetNode("GN15").SetMarble(board.NewMarble(board.ColorRed, 1, true))

	plays := GenerateRankJokerPlays(aBoard, deck.Card{}, board.ColorRed)
	for _, p := range plays {
		fmt.Printf("joker play: %s\n", p.String())
	}

}
