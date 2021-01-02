package move

import (
	"github.com/stretchr/testify/assert"
	"github.com/trichner/dog/dogd/engine/board"
	"testing"
)

func TestFindForwardPaths_NoExitFromHome(t *testing.T) {
	aBoard := board.NewBoard()

	redHome := aBoard.GetNode("RH1")

	paths := FindForwardPaths(redHome, 3)
	assert.Empty(t, paths)
}

func TestFindForwardPaths_MoveFromStart(t *testing.T) {
	aBoard := board.NewBoard()

	redStart := aBoard.GetNode("RS0")
	redStart.SetMarble(board.NewMarble(board.ColorRed, 1, false))

	paths := FindForwardPaths(redStart, 3)
	assert.Equal(t, paths[0], pathOf("RS0", "RN1", "RN2", "RN3"))
}

func TestFindForwardPaths_MoveFromStart_PassedStart(t *testing.T) {
	aBoard := board.NewBoard()

	redStart := aBoard.GetNode("RS0")
	redStart.SetMarble(board.NewMarble(board.ColorRed, 1, true))

	paths := FindForwardPaths(redStart, 3)

	var expected [][]board.NodeID
	expected = append(expected, pathOf("RS0", "RF1", "RF2", "RF3"))
	expected = append(expected, pathOf("RS0", "RN1", "RN2", "RN3"))
	assert.Equal(t, expected, paths)
}

func TestFindForwardPaths_BlockedStart(t *testing.T) {
	aBoard := board.NewBoard()

	redStart := aBoard.GetNode("RS0")
	redStart.SetMarble(board.NewMarble(board.ColorRed, 1, false))

	blockedNode := aBoard.GetNode("GN13")
	blockedNode.SetMarble(board.NewMarble(board.ColorGreen, 1, true))

	paths := FindForwardPaths(blockedNode, 13)

	assert.Empty(t, paths)
}

func TestFindForwardPaths_UnblockedStart(t *testing.T) {
	aBoard := board.NewBoard()

	redStart := aBoard.GetNode("RS0")
	redStart.SetMarble(board.NewMarble(board.ColorRed, 1, true))

	blockedNode := aBoard.GetNode("GN13")
	blockedNode.SetMarble(board.NewMarble(board.ColorGreen, 1, true))

	paths := FindForwardPaths(blockedNode, 3)

	var expected [][]board.NodeID
	expected = append(expected, pathOf("GN13", "GN14", "GN15", "RS0"))
	assert.Equal(t, expected, paths)
}

func TestFindForwardPaths_NoPassingInFinish(t *testing.T) {
	aBoard := board.NewBoard()

	redFinish1 := aBoard.GetNode("RF1")
	redFinish1.SetMarble(board.NewMarble(board.ColorRed, 1, true))

	redFinish2 := aBoard.GetNode("RF2")
	redFinish2.SetMarble(board.NewMarble(board.ColorRed, 2, true))

	paths := FindForwardPaths(redFinish1, 2)

	assert.Empty(t, paths)
}

func TestFindReversePaths_InFinish(t *testing.T) {
	aBoard := board.NewBoard()

	redFinish1 := aBoard.GetNode("RF1")
	redFinish1.SetMarble(board.NewMarble(board.ColorRed, 1, true))

	paths := FindReversePaths(redFinish1, 2)
	assert.Empty(t, paths)
}

func TestFindReversePaths_InHome(t *testing.T) {
	aBoard := board.NewBoard()

	redFinish1 := aBoard.GetNode("RH1")
	redFinish1.SetMarble(board.NewMarble(board.ColorRed, 1, false))

	paths := FindReversePaths(redFinish1, 2)
	assert.Empty(t, paths)
}

func TestFindReversePaths_AfterStart(t *testing.T) {
	aBoard := board.NewBoard()

	red2 := aBoard.GetNode("RN2")
	red2.SetMarble(board.NewMarble(board.ColorRed, 1, true))

	paths := FindReversePaths(red2, 4)

	var expected [][]board.NodeID
	expected = append(expected, pathOf("RN2", "RN1", "RS0", "GN15", "GN14"))
	assert.Equal(t, expected, paths)
}

func TestFindReversePaths_BlockingMarble(t *testing.T) {
	aBoard := board.NewBoard()

	red2 := aBoard.GetNode("RN2")
	red2.SetMarble(board.NewMarble(board.ColorRed, 1, true))

	redStart := aBoard.GetNode("RS0")
	redStart.SetMarble(board.NewMarble(board.ColorRed, 2, false))

	paths := FindReversePaths(red2, 4)

	assert.Empty(t, paths)
}

func pathOf(nodes ...board.NodeID) []board.NodeID {
	return nodes
}
