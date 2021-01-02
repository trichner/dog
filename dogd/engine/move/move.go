package move

import (
	"github.com/trichner/dog/dogd/engine/board"
)


func FindForwardPaths(start board.BoardNode, distance int) [][]board.NodeID {

	// cannot move from Home directly forward
	if start.Kind() == board.KindHome {
		return [][]board.NodeID{}
	}

	return findForwardPaths(start, start, distance)
}

func findForwardPaths(start, last board.BoardNode, distance int) [][]board.NodeID {
	self := []board.NodeID{last.ID()}
	marble := start.Marble()

	if distance <= 0 {
		return [][]board.NodeID{self}
	}

	var allPaths [][]board.NodeID
	for _, n := range last.Next() {

		occupant := n.Marble()

		// deal with finish
		if n.Kind() == board.KindFinish {
			// only same color marble can go to finish
			if marble.Color() != n.Color() {
				continue
			}

			// marble must have start once before
			if !marble.IsStartPassed() {
				continue
			}

			// can not pass occupied places in finish
			if occupant != nil {
				continue
			}
		}

		// if the start field is occupied by a marble that never passed start, it's blocking
		if n.Kind() == board.KindStart {
			if occupant != nil && !occupant.IsStartPassed() {
				continue
			}
		}

		paths := findForwardPaths(start, n, distance-1)
		for _, path := range paths {
			extendedPath := append(self, path...)
			allPaths = append(allPaths, extendedPath)
		}
	}
	return allPaths
}

func FindReversePaths(start board.BoardNode, distance int) [][]board.NodeID {
	return findReversePaths(start, start, distance)
}

func findReversePaths(start, last board.BoardNode, distance int) [][]board.NodeID {
	self := []board.NodeID{last.ID()}

	if distance <= 0 {
		return [][]board.NodeID{self}
	}

	var allPaths [][]board.NodeID
	for _, n := range last.Previous() {

		occupant := n.Marble()

		// if the start field is occupied by a marble that never passed start, it's blocking
		if n.Kind() == board.KindStart {
			if occupant != nil && !occupant.IsStartPassed() {
				continue
			}
		}

		paths := findReversePaths(start, n, distance-1)
		for _, path := range paths {
			extendedPath := append(self, path...)
			allPaths = append(allPaths, extendedPath)
		}
	}
	return allPaths
}
