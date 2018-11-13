// +build dot

package ast

import (
	"fmt"
	"strings"

	"github.com/awalterschulze/gographviz"
)

func dotLabel(a *AST, n Node) string {
	slice := a.slice(n.Pos)
	typelabel := n.Type.String()

	if n.Type == TypeText {
		return fmt.Sprintf(`%s[%v:%v] '%v'`, typelabel, n.Pos.Start, n.Pos.End, string(slice))
	}

	return fmt.Sprintf(`%s[%v:%v]`, typelabel, n.Pos.Start, n.Pos.End)
}

// ToDot …
func (a *AST) ToDot() string {
	dot := gographviz.NewGraph()
	dot.SetDir(true)

	a.root.visit(func(n Node) Node {
		src := n.ID()

		label := strings.Replace(dotLabel(a, n), "\"", "\\\"", -1)

		dot.AddNode("G", src, map[string]string{
			"label": fmt.Sprintf(`"%v"`, label),
		})

		for _, child := range n.Children {
			dst := child.ID()
			dot.AddEdge(src, dst, true, nil)
		}

		return n
	})

	return dot.String()
}
