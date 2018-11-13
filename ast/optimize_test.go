package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestOptimize …
func TestOptimize(t *testing.T) {
	ast := &AST{
		root: SequenceNode(
			TextNode(0, 2),
			SequenceNode(
				TextNode(2, 3),
			),
			ListNode(
				TextNode(3, 4),
				ListNode(
					SequenceNode(
						TextNode(6, 7),
						SequenceNode(
							TextNode(7, 8),
							TextNode(8, 9),
						),
					),
					NumericRangeNode(
						TextNode(9, 10),
						TextNode(12, 13),
						TextNode(15, 16),
					),
				),
			),
		),
	}

	expected := &AST{
		root: SequenceNode(
			TextNode(0, 2),
			TextNode(2, 3),
			ListNode(
				TextNode(3, 4),
				SequenceNode(
					TextNode(6, 7),
					TextNode(7, 8),
					TextNode(8, 9),
				),
				NumericRangeNode(
					TextNode(9, 10),
					TextNode(12, 13),
					TextNode(15, 16),
				),
			),
		),
	}

	actual := ast.Optimize()

	assert.Equal(t, expected, actual)
}

// BenchmarkOptimize …
func BenchmarkOptimize(b *testing.B) {
	ast := &AST{
		root: SequenceNode(
			TextNode(0, 2),
			SequenceNode(
				TextNode(2, 3),
			),
			ListNode(
				TextNode(3, 4),
				ListNode(
					SequenceNode(
						TextNode(6, 7),
						SequenceNode(
							TextNode(7, 8),
							TextNode(8, 9),
						),
					),
					NumericRangeNode(
						TextNode(9, 10),
						TextNode(12, 13),
						TextNode(15, 16),
					),
				),
			),
		),
	}

	for n := 0; n < b.N; n++ {
		_ = ast.Optimize()
	}
}
