package board

import (
	"fmt"
)

type node struct {
	id            NodeID
	number        int
	kind          NodeKind
	color         Color
	nextNodes     []*node
	previousNodes []*node
}

func (n *node) addNext(next *node) *node {
	n.nextNodes = append(n.nextNodes, next)
	return n
}

func (n *node) addPrevious(previous *node) *node {
	n.previousNodes = append(n.previousNodes, previous)
	return n
}

func (n *node) shallowEqual(other *node) bool {
	if !(n.id == other.id && n.number == other.number && n.kind == other.kind && n.color == other.color) {
		return false
	}

	return equalNodeList(n.previousNodes, other.previousNodes) && equalNodeList(n.nextNodes, other.nextNodes)
}

func equalNodeList(a, b []*node) bool {

	if len(a) != len(b) {
		return false
	}

	for _, node := range a {
		found := false
		for _, other := range b {
			if node.id == other.id {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func newStartNode(color Color) *node {
	return newNode(color, KindStart, 0)
}

func newHomeNode(color Color, n int) *node {
	return newNode(color, KindHome, n)
}

func newFinishNode(color Color, n int) *node {
	return newNode(color, KindFinish, n)
}

func newNormalNode(color Color, n int) *node {
	return newNode(color, KindNormal, n)
}

func newNode(color Color, kind NodeKind, n int) *node {
	return &node{
		id:            NodeIdOf(color, kind, n),
		number:        n,
		kind:          kind,
		color:         color,
		nextNodes:     make([]*node, 0),
		previousNodes: make([]*node, 0),
	}
}

func NodeIdOf(color Color, kind NodeKind, n int) NodeID {
	var colorChar string
	switch color {
	case ColorRed:
		colorChar = "R"
	case ColorYellow:
		colorChar = "Y"
	case ColorBlue:
		colorChar = "B"
	case ColorGreen:
		colorChar = "G"
	}

	var kindChar string
	switch kind {
	case KindNormal:
		kindChar = "N"
	case KindStart:
		kindChar = "S"
	case KindFinish:
		kindChar = "F"
	case KindHome:
		kindChar = "H"
	}
	return NodeID(fmt.Sprintf("%s%s%d", colorChar, kindChar, n))
}
