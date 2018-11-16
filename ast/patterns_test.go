package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestToPattern …
func TestToPattern(t *testing.T) {
	testCases := []struct {
		desc     string
		str      string
		root     Node
		expected string
	}{
		{
			desc:     "Simple list expansion",
			str:      "{1,2}",
			root:     ListNode(TextNode(1, 2), TextNode(3, 4)),
			expected: "(1|2)",
		},

		// Numeric Ranges: start < end
		{
			desc:     "Numeric range expansion: Forwards, No Step",
			str:      "{1..5}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), NilNode()),
			expected: "(1|2|3|4|5)",
		},
		{
			desc:     "Numeric range expansion: Forwards, Positive Step",
			str:      "{1..5..2}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), TextNode(7, 8)),
			expected: "(1|3|5)",
		},
		{
			desc:     "Numeric range expansion: Forwards, Negative Step",
			str:      "{1..5..-2}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), TextNode(7, 9)),
			expected: "(5|3|1)",
		},

		// Numeric Ranges: end < start
		{
			desc:     "Numeric range expansion: Backwards, No Step",
			str:      "{5..1}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), NilNode()),
			expected: "(5|4|3|2|1)",
		},
		{
			desc:     "Numeric range expansion: Backwards, Positive Step",
			str:      "{5..1..2}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), TextNode(7, 8)),
			expected: "(5|3|1)",
		},
		{
			desc:     "Numeric range expansion: Backwards, Negative Step",
			str:      "{5..1..-2}",
			root:     NumericRangeNode(TextNode(1, 2), TextNode(4, 5), TextNode(7, 9)),
			expected: "(1|3|5)",
		},

		// Nested Nodes
		{
			desc: "Nested expansions",
			str:  "ab{c,{d,{0..6..2}}}",
			root: SequenceNode(
				TextNode(0, 2),
				ListNode(
					TextNode(3, 4),
					TextNode(6, 7),
					NumericRangeNode(
						TextNode(9, 10),
						TextNode(12, 13),
						TextNode(15, 16),
					),
				),
			),
			expected: "ab(c|d|(0|2|4|6))",
		},

		// Wildcards
		{
			desc: "Wildcard nodes",
			str:  "{images,thumbs}/*/{0..10}.jpg",
			root: SequenceNode(
				ListNode(
					TextNode(1, 7),
					TextNode(8, 14),
				),
				TextNode(15, 16),
				WildcardNode(16, 17),
				TextNode(17, 18),
				NumericRangeNode(
					TextNode(19, 20),
					TextNode(22, 24),
					NilNode(),
				),
				TextNode(25, 29),
			),
			expected: "(images|thumbs)/.*?/(0|1|2|3|4|5|6|7|8|9|10)\\.jpg",
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
			expected: "x(52|(55|56|57|58|59|60))y",
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
			expected: "x(52|y(55|56|57|58|59|60)|)z",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ast := &AST{Input: []byte(tt.str), Root: tt.root}
			actual, err := ast.ToPattern(PatternOpts{Anchored: false})

			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}
