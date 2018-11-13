package ast

type visitor func(n Node) Node

func (n Node) visit(visitor visitor) Node {
	n = visitor(n)

	for idx, child := range n.Children {
		n.Children[idx] = child.visit(visitor)
	}

	return n
}
