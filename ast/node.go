package ast

import (
	"fmt"
	"math"
)

// Node …
type Node struct {
	Type     Type
	Pos      Position
	Children []Node
}

func NilNode() Node {
	return Node{Type: TypeNil}
}

func TextNode(start int, end int) Node {
	return Node{
		Type: TypeText,
		Pos: Position{
			Start: start,
			End:   end,
		},
	}
}

func ListNode(children ...Node) Node {
	return Node{
		Type:     TypeList,
		Children: children,
		Pos:      bounds(children...),
	}
}

func SequenceNode(children ...Node) Node {
	return Node{
		Type:     TypeSequence,
		Children: children,
		Pos:      bounds(children...),
	}
}

func ArbitraryRangeNode(start Node, end Node) Node {
	children := []Node{start, end}

	return Node{
		Type:     TypeRangeArbitrary,
		Children: children,
		Pos:      bounds(children...),
	}
}

func NumericRangeNode(start Node, end Node, step Node) Node {
	children := []Node{start, end, step}

	return Node{
		Type:     TypeRangeNumeric,
		Children: children,
		Pos:      bounds(children...),
	}
}

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

// Type …
type Type int

const (
	TypeNil Type = iota
	TypeText
	TypeList
	TypeSequence
	TypeRangeNumeric
	TypeRangeArbitrary
)

// Position …
type Position struct {
	Start int
	End   int
}

const (
	RangeStepNone = 0
)

func (n Node) ID() string {
	return fmt.Sprintf("n_%d_%d_%d", n.Type, n.Pos.Start, n.Pos.End)
}

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
