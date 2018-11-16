package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestToStrings …
func TestToStrings(t *testing.T) {
	testCases := []struct {
		desc     string
		str      string
		root     Node
		expected []string
	}{
		{
			desc:     "Simple list expansion",
			str:      "{1,2}",
			root:     ListNode(TextNode(1, 2), TextNode(3, 4)),
			expected: []string{"1", "2"},
		},

		// Numeric Ranges: start < end
		{
			desc:     "Numeric range expansion: Forwards, No Step",
			str:      "{1..5}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), NilNode()),
			expected: []string{"1", "2", "3", "4", "5"},
		},
		{
			desc:     "Numeric range expansion: Forwards, Positive Step",
			str:      "{1..5..2}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), TextNode(7, 8)),
			expected: []string{"1", "3", "5"},
		},
		{
			desc:     "Numeric range expansion: Forwards, Negative Step",
			str:      "{1..5..-2}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), TextNode(7, 9)),
			expected: []string{"5", "3", "1"},
		},

		// Numeric Ranges: end < start
		{
			desc:     "Numeric range expansion: Backwards, No Step",
			str:      "{5..1}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), NilNode()),
			expected: []string{"5", "4", "3", "2", "1"},
		},
		{
			desc:     "Numeric range expansion: Backwards, Positive Step",
			str:      "{5..1..2}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), TextNode(7, 8)),
			expected: []string{"5", "3", "1"},
		},
		{
			desc:     "Numeric range expansion: Backwards, Negative Step",
			str:      "{5..1..-2}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), TextNode(7, 9)),
			expected: []string{"1", "3", "5"},
		},

		// Nested Nodes
		{
			desc: "Nested expansions",
			str:  "ab{c,{d,{0..6..2}}}",
			root: SequenceNode(
				TextNode(0, 2),
				ListNode(
					TextNode(3, 4),
					ListNode(
						TextNode(6, 7),
						NumericRangeNode(
							TextNode(9, 10),
							TextNode(12, 13),
							TextNode(15, 16),
						),
					),
				),
			),
			expected: []string{
				"abc",
				"abd",
				"ab0",
				"ab2",
				"ab4",
				"ab6",
			},
		},

		// Other test cases
		{
			desc: "Other 1",
			str:  "x{52,{55..60}}y",
			root: SequenceNode(
				TextNode(0, 1),
				ListNode(
					TextNode(2, 4),
					NumericRangeNode(
						TextNode(6, 8),
						TextNode(10, 12),
						NilNode(),
					),
				),
				TextNode(14, 15),
			),
			expected: []string{
				"x52y",
				"x55y",
				"x56y",
				"x57y",
				"x58y",
				"x59y",
				"x60y",
			},
		},
		{
			desc: "Other 2",
			str:  "x{52,y{55..60},}z",
			root: SequenceNode(
				TextNode(0, 1),
				ListNode(
					TextNode(2, 4),
					SequenceNode(
						TextNode(5, 6),
						NumericRangeNode(
							TextNode(7, 9),
							TextNode(11, 13),
							NilNode(),
						),
					),
					NilNode(),
				),
				TextNode(16, 17),
			),
			expected: []string{
				"x52z",
				"xy55z",
				"xy56z",
				"xy57z",
				"xy58z",
				"xy59z",
				"xy60z",
				"xz",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ast := &AST{Input: []byte(tt.str), Root: tt.root}
			actual, err := ast.ToStrings()

			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}
