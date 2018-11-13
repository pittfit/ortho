package ast

// AST …
type AST struct {
	buf  []byte
	root Node
}

func (a *AST) slice(p Position) []byte {
	return a.buf[p.Start:p.End]
}
