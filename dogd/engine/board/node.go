package board

import "fmt"

//go:generate enumer -type=Color
type Color int

//go:generate enumer -type=NodeKind
type NodeKind int

const (
	ColorRed Color = iota
	ColorYellow
	ColorBlue
	ColorGreen
)

const (
	KindNormal NodeKind = iota
	KindStart
	KindFinish
	KindHome
)

type NodeID string

const InvalidNode NodeID = "InvalidNode"

type BoardNode interface {
	ID() NodeID
	Kind() NodeKind
	Color() Color
	Next() []BoardNode
	Previous() []BoardNode
	Marble() *Marble
	SetMarble(m *Marble)
}

var _ BoardNode = &boardNode{}

type boardNode struct {
	id    NodeID
	board *Board
	node  *node
}

func (b *boardNode) ID() NodeID {
	return b.id
}

func (b *boardNode) Kind() NodeKind {
	return b.node.kind
}

func (b *boardNode) Color() Color {
	return b.node.color
}

func (b *boardNode) Next() []BoardNode {
	return wrapNodes(b.board, b.node.nextNodes)
}

func (b *boardNode) Previous() []BoardNode {
	return wrapNodes(b.board, b.node.previousNodes)
}

func (b *boardNode) Marble() *Marble {
	marble, ok := b.board.marbles[b.id]
	if !ok {
		return nil
	}
	return marble
}

func (b *boardNode) SetMarble(m *Marble) {
	marbles := b.board.marbles
	for k, v := range b.board.marbles {
		if v.number == m.number && v.color == m.color {
			delete(marbles, k)
			marbles[b.node.id] = m
			return
		}
	}

	panic(fmt.Errorf("invalid marble %s%d, cannot find it on the board", m.color, m.number))
}

func wrapNodes(b *Board, n []*node) []BoardNode {
	var nodes []BoardNode
	for _, node := range n {
		nodes = append(nodes, wrapNode(b, node))
	}
	return nodes
}

func wrapNode(b *Board, n *node) BoardNode {
	return &boardNode{
		id:    n.id,
		board: b,
		node:  n,
	}
}
