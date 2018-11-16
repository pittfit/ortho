package ast

import "github.com/pittfit/ortho/token"

// AST …
type AST struct {
	input []byte
	root  Node
}

func (a *AST) slice(l token.Location) []byte {
	return a.input[l.Start:l.End]
}
