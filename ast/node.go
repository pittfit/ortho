package ast

import (
	"fmt"
	"math"
)

// Node A node in the AST
type Node struct {
	Type     Type
	Pos      Position
	Children []Node
}

// NilNode Create a new blank node
func NilNode() Node {
	return Node{Type: TypeNil}
}

// TextNode Create a new text node
func TextNode(start int, end int) Node {
	return Node{
		Type: TypeText,
		Pos: Position{
			Start: start,
			End:   end,
		},
	}
}

// ListNode Create a new list node
func ListNode(children ...Node) Node {
	return Node{
		Type:     TypeList,
		Children: children,
		Pos:      bounds(children...),
	}
}

// SequenceNode Create a new sequence node
func SequenceNode(children ...Node) Node {
	return Node{
		Type:     TypeSequence,
		Children: children,
		Pos:      bounds(children...),
	}
}

// ArbitraryRangeNode Create a new range node for ascii characters
func ArbitraryRangeNode(start Node, end Node) Node {
	children := []Node{start, end}

	return Node{
		Type:     TypeRangeArbitrary,
		Children: children,
		Pos:      bounds(children...),
	}
}

// NumericRangeNode Create a new range node for numeric ranges
func NumericRangeNode(start Node, end Node, step Node) Node {
	children := []Node{start, end, step}

	return Node{
		Type:     TypeRangeNumeric,
		Children: children,
		Pos:      bounds(children...),
	}
}

// Given a list of children return the bounds that
// would cover all children from start to end
func bounds(children ...Node) Position {
	var pos = Position{Start: math.MaxInt64, End: 0}

	for _, child := range children {
		if child.Type != TypeNil {
			if pos.Start > child.Pos.Start {
				pos.Start = child.Pos.Start
			}

			if pos.End < child.Pos.End {
				pos.End = child.Pos.End
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

// Position The start/end position into the buffer for this node
type Position struct {
	Start int
	End   int
}

// ID A unique identifier for this node
// Two separate nodes sharing the same
// ID mean that when given an identical
// buffer would produce the same output
// This does NOT mean that two different
// nodes cannot also produce the same output
func (n Node) ID() string {
	return fmt.Sprintf("n_%d_%d_%d", n.Type, n.Pos.Start, n.Pos.End)
}

// String A readable version of the node type
func (t Type) String() string {
	switch t {
	case TypeNil:
		return "nil"
	case TypeText:
		return "text"
	case TypeSequence:
		return "sequence"
	case TypeList:
		return "list"
	case TypeRangeNumeric:
		return "range_numeric"
	}
	return "unknown"
}
