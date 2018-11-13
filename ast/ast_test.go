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
			desc: "Numeric range expansion: Backwards, No Step",
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
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ast := &AST{buf: []byte(tt.str), root: tt.root}
			actual, err := ast.ToStrings()

			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}
