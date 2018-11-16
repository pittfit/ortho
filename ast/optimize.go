package ast

// Optimize Modifies nodes in the AST to produce a smaller tree with the same semantics
func (a *AST) Optimize() *AST {
	root := a.root

	root = root.visit(liftNestedNodes)
	root = root.visit(mergeConsecutiveTextNodes)
	root = root.visit(liftSingleChildSequences)

	return &AST{
		buf:  a.buf,
		root: root,
	}
}

func liftNestedNodes(n Node) Node {
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
}

func mergeConsecutiveTextNodes(n Node) Node {
	if n.Type != TypeSequence {
		return n
	}

	children := []Node{}
	previousNodeType := TypeNil

	for _, child := range n.Children {
		if child.Type == TypeText && previousNodeType == TypeText {
			children[len(children)-1].Loc.End = child.Loc.End
		} else {
			children = append(children, child)
		}

		previousNodeType = child.Type
	}

	n.Children = children

	return n
}

func liftSingleChildSequences(n Node) Node {
	if n.Type != TypeSequence {
		return n
	}

	if len(n.Children) != 1 {
		return n
	}

	return n.Children[0]
}
