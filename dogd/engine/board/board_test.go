package board

import (
	"fmt"
	"testing"
)

func TestPrintAsGraphviz(t *testing.T) {
	nodes := generateNodes()
	dot := NodesAsGraphviz(nodes)
	fmt.Println(dot)
}
