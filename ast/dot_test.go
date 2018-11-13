// +build dot

package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestToDot …
func TestToDot(t *testing.T) {
	ast := &AST{
		buf:  []byte("{1,2}"),
		root: ListNode(TextNode(1, 2), TextNode(3, 4)),
	}

	actual := ast.ToDot()
	expected := "digraph  {\n\tn_2_1_4->n_1_1_2;\n\tn_2_1_4->n_1_3_4;\n\tn_1_1_2 [ label=\"text[1:2] '1'\" ];\n\tn_1_3_4 [ label=\"text[3:4] '2'\" ];\n\tn_2_1_4 [ label=\"list[1:4]\" ];\n\n}\n"

	assert.Equal(t, expected, actual)
}

// BenchmarkToDot …
func BenchmarkToDot(b *testing.B) {
	ast := &AST{
		buf:  []byte("{1,2}"),
		root: ListNode(TextNode(1, 2), TextNode(3, 4)),
	}

	for i := 0; i < b.N; i++ {
		ast.ToDot()
	}
}
