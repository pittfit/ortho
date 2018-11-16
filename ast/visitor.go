package ast

import (
	"github.com/pittfit/ortho/tracing"
)

type visitor func(n Node) Node

func (n Node) visit(visitor visitor) Node {
	defer tracing.End(tracing.Begin("visit", n.ID()))

	n = visitor(n)

	for idx, child := range n.Children {
		n.Children[idx] = child.visit(visitor)
	}

	return n
}
