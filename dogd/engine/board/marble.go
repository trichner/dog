package board

type BoardMarble interface {
	Color() Color
	Number() int
	IsStartPassed() bool
	Position() BoardNode
}

type Marble struct {
	color       Color
	number      int
	passedStart bool
}

func NewMarble(color Color, number int, hasPassedStart bool) *Marble {
	return &Marble{
		color:       color,
		number:      number,
		passedStart: hasPassedStart,
	}
}

func (m *Marble) Number() int {
	return m.number
}

func (m *Marble) Color() Color {
	return m.color
}

func (m *Marble) IsStartPassed() bool {
	return m.passedStart
}

type wrappedMarble struct {
	*Board
	*Marble
	position NodeID
}

func (w *wrappedMarble) Position() BoardNode {
	nodeOfMarble := w.nodes[w.position]
	return wrapNode(w.Board, nodeOfMarble)
}

func wrapMarble(b *Board, m *Marble, position NodeID) BoardMarble {
	return &wrappedMarble{
		Board:    b,
		Marble:   m,
		position: position,
	}
}
