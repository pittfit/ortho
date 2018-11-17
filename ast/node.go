package ast

import (
	"fmt"
	"math"
	"strings"

	"github.com/pittfit/ortho/token"
	"github.com/pittfit/ortho/tracing"
)

// Node A node in the AST
type Node struct {
	Type     Type
	Loc      token.Location
	Children []Node
}

// NilNode Create a new blank node
func NilNode() Node {
	tracing.Call("NilNode", "")

	return Node{Type: TypeNil}
}

// TextNode Create a new text node
func TextNode(start int, end int) Node {
	tracing.Call("TextNode", "")

	return Node{
		Type: TypeText,
		Loc:  token.Location{Start: start, End: end},
	}
}

// WildcardNode Create a new wildcard node
func WildcardNode(start int, end int) Node {
	tracing.Call("WildcardNode", "")

	return Node{
		Type: TypeWildcard,
		Loc:  token.Location{Start: start, End: end},
	}
}

// ListNode Create a new list node
func ListNode(children ...Node) Node {
	tracing.Call("ListNode", "")

	return Node{
		Type:     TypeList,
		Children: children,
		Loc:      bounds(children...),
	}
}

// SequenceNode Create a new sequence node
func SequenceNode(children ...Node) Node {
	tracing.Call("SequenceNode", "")

	return Node{
		Type:     TypeSequence,
		Children: children,
		Loc:      bounds(children...),
	}
}

// ArbitraryRangeNode Create a new range node for ascii characters
func ArbitraryRangeNode(start Node, end Node) Node {
	tracing.Call("ArbitraryRangeNode", "")

	children := []Node{start, end}

	return Node{
		Type:     TypeRangeArbitrary,
		Children: children,
		Loc:      bounds(children...),
	}
}

// NumericRangeNode Create a new range node for numeric ranges
func NumericRangeNode(start Node, end Node, step Node) Node {
	tracing.Call("NumericRangeNode", "")

	children := []Node{start, end, step}

	return Node{
		Type:     TypeRangeNumeric,
		Children: children,
		Loc:      bounds(children...),
	}
}

// Given a list of children return the bounds that
// would cover all children from start to end
func bounds(children ...Node) token.Location {
	if len(children) == 0 {
		return token.Location{}
	}

	pos := children[0].Loc

	for _, child := range children {
		if child.Type != TypeNil {
			if pos.Start > child.Loc.Start {
				pos.Start = child.Loc.Start
			}

			if pos.End < child.Loc.End {
				pos.End = child.Loc.End
			}
		}
	}

	return pos
}

// Type The type of node in the AST
type Type int

const (
	// TypeNil A blank node used to omit values
	// It is used in lists to produce an item with no output
	// It is used in ranges to omit the step node
	TypeNil Type = iota

	// TypeText a literal text string
	TypeText

	// TypeWildcard a wildcard in the text string
	// This is just a string of text that
	// can be considered a wildcard
	TypeWildcard

	// TypeList an enumeration of possible nodes
	// When output as a string this produces a
	// separate string for each child node
	TypeList

	// TypeSequence A match from each of the child nodes
	// When output as a string this produces a list of
	// all possible outputs picking from each child
	TypeSequence

	// TypeRangeNumeric A range of numbers.
	// Can be ascending or descending
	// Can be have a negative step (which flips its order)
	// A TypeNil step node means the step is omitted which steps by 1
	TypeRangeNumeric

	// TypeRangeArbitrary A range of ascii characters
	TypeRangeArbitrary
)

// ID A unique identifier for this node
// Two separate nodes sharing the same
// ID mean that when given an identical
// buffer would produce the same output
// This does NOT mean that two different
// nodes cannot also produce the same output
func (n Node) ID() string {
	return fmt.Sprintf("n_%d_%d_%d", n.Type, n.Loc.Start, n.Loc.End)
}

// String A readable version of the node type
func (t Type) String() string {
	switch t {
	case TypeNil:
		return "nil"
	case TypeText:
		return "text"
	case TypeWildcard:
		return "wildcard"
	case TypeSequence:
		return "sequence"
	case TypeList:
		return "list"
	case TypeRangeNumeric:
		return "range_numeric"
	}
	return "unknown"
}

// String A readable version of the node
func (n Node) String() string {
	str := strings.Builder{}

	for _, line := range n.lines() {
		str.WriteString(line)
		str.WriteString("\n")
	}

	return str.String()
}

// String A readable version of the node
func (n Node) lines() []string {
	header := fmt.Sprintf("(%s [%d:%d]", n.Type.String(), n.Loc.Start, n.Loc.End)
	footer := ")"

	if len(n.Children) == 0 {
		return []string{fmt.Sprintf("%s%s", header, footer)}
	}

	lines := []string{}
	lines = append(lines, header)

	for _, child := range n.Children {
		for _, line := range child.lines() {
			lines = append(lines, fmt.Sprintf("\t%s", line))
		}
	}

	lines = append(lines, footer)

	return lines
}
