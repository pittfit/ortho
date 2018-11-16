package ast

import "github.com/pittfit/ortho/token"

// AST …
type AST struct {
	buf  []byte
	root Node
}

func (a *AST) slice(l token.Location) []byte {
	return a.buf[l.Start:l.End]
}
