package ast

import "github.com/pittfit/ortho/token"

// AST …
type AST struct {
	Input []byte
	Root  Node
}

func NewAST(input []byte, root Node) *AST {
	return &AST{Input: input, Root: root}
}

func (a *AST) slice(l token.Location) []byte {
	return a.Input[l.Start:l.End]
}
