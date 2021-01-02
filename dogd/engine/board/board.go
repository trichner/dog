package board

import (
	"fmt"
	"strings"
)

type Board struct {
	nodes   map[NodeID]*node
	marbles map[NodeID]*Marble
}

func NewBoard() *Board {
	nodes := generateNodes()
	marbles := generateMarbles()

	return &Board{
		nodes:   nodes,
		marbles: marbles,
	}
}

func (b *Board) GetNode(id NodeID) BoardNode {
	internalNode := b.nodes[id]
	return wrapNode(b, internalNode)
}

func (b *Board) GetMarble(color Color, number int) BoardMarble {
	position, marble := b.findMarble(color, number)
	return wrapMarble(b, marble, position)
}

func (b *Board) findMarble(color Color, number int) (NodeID, *Marble) {
	for k, v := range b.marbles {
		if v.number == number && v.color == color {
			return k, v
		}
	}
	panic(fmt.Errorf("cannot find marble on board %s%d", color, number))
}

func generateMarbles() map[NodeID]*Marble {

	marbles := make(map[NodeID]*Marble)

	colors := []Color{ColorRed, ColorYellow, ColorBlue, ColorGreen}
	for _, c := range colors {
		for i := 0; i < 4; i++ {
			pos := NodeIdOf(c, KindHome, i+1)
			marbles[pos] = NewMarble(c, i, false)
		}
	}

	return marbles
}

func generateNodes() map[NodeID]*node {

	nodes := make(map[NodeID]*node)

	colors := []Color{ColorRed, ColorYellow, ColorBlue, ColorGreen}
	var first *node
	var previous *node
	for _, color := range colors {

		start := newStartNode(color)
		nodes[start.id] = start

		// keep the very first one around to close the loop
		if first == nil {
			first = start
		}

		// link up with previous color
		if previous != nil {
			start.addPrevious(previous)
			previous.addNext(start)
		}

		// add home nodes of color
		for i := 1; i <= 4; i++ {
			h := newHomeNode(color, i)
			nodes[h.id] = h
			h.addNext(start)
		}

		// add finish nodes of color
		var previousFinish = start
		for i := 1; i <= 4; i++ {
			f := newFinishNode(color, i)
			nodes[f.id] = f
			previousFinish.addNext(f)
			previousFinish = f
		}

		// add normal nodes for the current side
		var previousNormal = start
		for i := 1; i <= 15; i++ {
			f := newNormalNode(color, i)
			nodes[f.id] = f
			previousNormal.addNext(f)
			f.addPrevious(previousNormal)
			previousNormal = f
		}
		previous = previousNormal
	}

	// close loop
	first.addPrevious(previous)
	previous.addNext(first)

	return nodes
}

func NodesAsGraphviz(nodes map[NodeID]*node) string {
	var sb strings.Builder
	sb.WriteString("strict digraph {\n")

	for k, v := range nodes {
		sb.WriteString(fmt.Sprintf("  %s [fillcolor = \"%s\", style = \"filled\"]\n", k, toHexColor(v.color)))
		for _, next := range v.nextNodes {
			sb.WriteString(fmt.Sprintf("  %s -> %s\n", k, next.id))
		}
		for _, previous := range v.previousNodes {
			sb.WriteString(fmt.Sprintf("  %s -> %s\n", previous.id, k))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("}")
	return sb.String()
}

func toHexColor(c Color) string {
	switch c {
	case ColorRed:
		return "#FF0000"
	case ColorYellow:
		return "#FFFF00"
	case ColorBlue:
		return "#0000FF"
	case ColorGreen:
		return "#00FF00"
	}
	panic(fmt.Sprintf("unknown color: %d", c))
}
