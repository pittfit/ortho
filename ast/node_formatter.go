package ast

import (
	"fmt"
	"strings"
)

// NodeFormatter Format a node
type NodeFormatter struct {
	indent string
}

// NewNodeFormatter Create a new node formatter
func NewNodeFormatter(indent string) *NodeFormatter {
	return &NodeFormatter{indent: indent}
}

// ToString Convert a node to a LISP-style string
func (f *NodeFormatter) ToString(n Node) string {
	return strings.Join(f.ToLines(n), "\n")
}

// ToLines Convert a node to a LISP-style slice of strings
func (f *NodeFormatter) ToLines(n Node) []string {
	header := fmt.Sprintf("(%s [%d:%d]", n.Type.String(), n.Loc.Start, n.Loc.End)
	footer := ")"

	if len(n.Children) == 0 {
		return []string{fmt.Sprintf("%s%s", header, footer)}
	}

	lines := []string{}
	lines = append(lines, header)

	for _, child := range n.Children {
		for _, line := range f.ToLines(child) {
			lines = append(lines, fmt.Sprintf("%s%s", f.indent, line))
		}
	}

	lines = append(lines, footer)

	return lines
}
