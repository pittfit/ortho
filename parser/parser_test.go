package parser

import (
	"testing"

	"github.com/pittfit/ortho/ast"
	"github.com/pittfit/ortho/tracing"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input string
	root  ast.Node
}{
	{
		input: "{,}",
		root:  ast.ListNode(ast.NilNode(), ast.NilNode()),
	},
	{
		input: "{a,}",
		root:  ast.ListNode(ast.TextNode(1, 2), ast.NilNode()),
	},
	{
		input: "{a,b}",
		root:  ast.ListNode(ast.TextNode(1, 2), ast.TextNode(3, 4)),
	},
	{
		input: "{a,b,c}",
		root:  ast.ListNode(ast.TextNode(1, 2), ast.TextNode(3, 4), ast.TextNode(5, 6)),
	},
	{
		input: "{a,b,}",
		root:  ast.ListNode(ast.TextNode(1, 2), ast.TextNode(3, 4), ast.NilNode()),
	},
	{
		input: "{0..10}",
		root:  ast.NumericRangeNode(ast.TextNode(1, 2), ast.TextNode(4, 6), ast.NilNode()),
	},
	{
		input: "{0..10..02}",
		root:  ast.NumericRangeNode(ast.TextNode(1, 2), ast.TextNode(4, 6), ast.TextNode(8, 10)),
	},
}

func TestParse(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			input := []byte(tC.input)

			tracing.Enable()
			ast := NewParser(input).Parse()
			tracing.Disable()

			actual := ast.Root
			expected := tC.root

			assert.Equal(t, expected.String(), actual.String())
		})
	}
}

func BenchmarkParse(b *testing.B) {
	for _, tC := range testCases {
		b.Run(tC.input, func(b *testing.B) {
			input := []byte(tC.input)

			for i := 0; i < b.N; i++ {
				NewParser(input).Parse()
			}
		})
	}
}
