package ast

// Optimize Modifies nodes in the AST to produce a smaller tree with the same semantics
func (a *AST) Optimize() *AST {
	// Walk through the AST
	newRoot := a.root.visit(func(n Node) Node {
		if n.Type != TypeList && n.Type != TypeSequence {
			return n
		}

		// List nodes can merge children of child list nodes
		// Sequence nodes can merge children of child sequence nodes
		children := []Node{}

		for _, child := range n.Children {
			if child.Type == n.Type {
				children = append(children, child.Children...)
			} else {
				children = append(children, child)
			}
		}

		n.Children = children

		return n
	})

	return &AST{
		buf:  a.buf,
		root: newRoot,
	}
}