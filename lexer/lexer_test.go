package lexer

import (
	"testing"

	"github.com/pittfit/ortho/token"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  string
	tokens []token.Token
}{
	{
		input: "{,}",
		tokens: []token.Token{
			token.NewToken(token.BRACE_OPEN, 0, 1),
			token.NewToken(token.LIST_SEPARATOR, 1, 2),
			token.NewToken(token.BRACE_CLOSE, 2, 3),
			token.NewToken(token.EOF, 3, 3),
		},
	},
	{
		input: "{a,}",
		tokens: []token.Token{
			token.NewToken(token.BRACE_OPEN, 0, 1),
			token.NewToken(token.LITERAL, 1, 2),
			token.NewToken(token.LIST_SEPARATOR, 2, 3),
			token.NewToken(token.BRACE_CLOSE, 3, 4),
			token.NewToken(token.EOF, 4, 4),
		},
	},
	{
		input: "{ab,c}",
		tokens: []token.Token{
			token.NewToken(token.BRACE_OPEN, 0, 1),
			token.NewToken(token.LITERAL, 1, 3),
			token.NewToken(token.LIST_SEPARATOR, 3, 4),
			token.NewToken(token.LITERAL, 4, 5),
			token.NewToken(token.BRACE_CLOSE, 5, 6),
			token.NewToken(token.EOF, 6, 6),
		},
	},
	{
		input: "{1..10}",
		tokens: []token.Token{
			token.NewToken(token.BRACE_OPEN, 0, 1),
			token.NewToken(token.LITERAL, 1, 2),
			token.NewToken(token.RANGE_SEPARATOR, 2, 4),
			token.NewToken(token.LITERAL, 4, 6),
			token.NewToken(token.BRACE_CLOSE, 6, 7),
			token.NewToken(token.EOF, 7, 7),
		},
	},
	{
		input: "{1...10}",
		tokens: []token.Token{
			token.NewToken(token.BRACE_OPEN, 0, 1),
			token.NewToken(token.LITERAL, 1, 2),
			token.NewToken(token.RANGE_SEPARATOR, 2, 4),
			token.NewToken(token.LITERAL, 4, 7),
			token.NewToken(token.BRACE_CLOSE, 7, 8),
			token.NewToken(token.EOF, 8, 8),
		},
	},
	{
		input: "{1....10}",
		tokens: []token.Token{
			token.NewToken(token.BRACE_OPEN, 0, 1),
			token.NewToken(token.LITERAL, 1, 2),
			token.NewToken(token.RANGE_SEPARATOR, 2, 4),
			token.NewToken(token.RANGE_SEPARATOR, 4, 6),
			token.NewToken(token.LITERAL, 6, 8),
			token.NewToken(token.BRACE_CLOSE, 8, 9),
			token.NewToken(token.EOF, 9, 9),
		},
	},
	{
		input: "{01..10..2}",
		tokens: []token.Token{
			token.NewToken(token.BRACE_OPEN, 0, 1),
			token.NewToken(token.LITERAL, 1, 3),
			token.NewToken(token.RANGE_SEPARATOR, 3, 5),
			token.NewToken(token.LITERAL, 5, 7),
			token.NewToken(token.RANGE_SEPARATOR, 7, 9),
			token.NewToken(token.LITERAL, 9, 10),
			token.NewToken(token.BRACE_CLOSE, 10, 11),
			token.NewToken(token.EOF, 11, 11),
		},
	},
	{
		input: "foo/*/bar{01..10..2}.jpg",
		tokens: []token.Token{
			token.NewToken(token.LITERAL, 0, 4),
			token.NewToken(token.WILDCARD, 4, 5),
			token.NewToken(token.LITERAL, 5, 9),
			token.NewToken(token.BRACE_OPEN, 9, 10),
			token.NewToken(token.LITERAL, 10, 12),
			token.NewToken(token.RANGE_SEPARATOR, 12, 14),
			token.NewToken(token.LITERAL, 14, 16),
			token.NewToken(token.RANGE_SEPARATOR, 16, 18),
			token.NewToken(token.LITERAL, 18, 19),
			token.NewToken(token.BRACE_CLOSE, 19, 20),
			token.NewToken(token.LITERAL, 20, 24),
			token.NewToken(token.EOF, 24, 24),
		},
	},
	{
		input: "foo/*/bar{01..10..2}..jpg",
		tokens: []token.Token{
			token.NewToken(token.LITERAL, 0, 4),
			token.NewToken(token.WILDCARD, 4, 5),
			token.NewToken(token.LITERAL, 5, 9),
			token.NewToken(token.BRACE_OPEN, 9, 10),
			token.NewToken(token.LITERAL, 10, 12),
			token.NewToken(token.RANGE_SEPARATOR, 12, 14),
			token.NewToken(token.LITERAL, 14, 16),
			token.NewToken(token.RANGE_SEPARATOR, 16, 18),
			token.NewToken(token.LITERAL, 18, 19),
			token.NewToken(token.BRACE_CLOSE, 19, 20),
			token.NewToken(token.LITERAL, 20, 25),
			token.NewToken(token.EOF, 25, 25),
		},
	},
	{
		input: "\\{{a,b}",
		tokens: []token.Token{
			token.NewToken(token.LITERAL, 1, 2),
			token.NewToken(token.BRACE_OPEN, 2, 3),
			token.NewToken(token.LITERAL, 3, 4),
			token.NewToken(token.LIST_SEPARATOR, 4, 5),
			token.NewToken(token.LITERAL, 5, 6),
			token.NewToken(token.BRACE_CLOSE, 6, 7),
			token.NewToken(token.EOF, 7, 7),
		},
	},
	{
		input: "\\\\{{a,b}",
		tokens: []token.Token{
			token.NewToken(token.LITERAL, 1, 2),
			token.NewToken(token.BRACE_OPEN, 2, 3),
			token.NewToken(token.BRACE_OPEN, 3, 4),
			token.NewToken(token.LITERAL, 4, 5),
			token.NewToken(token.LIST_SEPARATOR, 5, 6),
			token.NewToken(token.LITERAL, 6, 7),
			token.NewToken(token.BRACE_CLOSE, 7, 8),
			token.NewToken(token.EOF, 8, 8),
		},
	},
	{
		input: "\\a{{a,b}",
		tokens: []token.Token{
			token.NewToken(token.LITERAL, 1, 2),
			token.NewToken(token.BRACE_OPEN, 2, 3),
			token.NewToken(token.BRACE_OPEN, 3, 4),
			token.NewToken(token.LITERAL, 4, 5),
			token.NewToken(token.LIST_SEPARATOR, 5, 6),
			token.NewToken(token.LITERAL, 6, 7),
			token.NewToken(token.BRACE_CLOSE, 7, 8),
			token.NewToken(token.EOF, 8, 8),
		},
	},
	{
		input: "{a\\,b}",
		tokens: []token.Token{
			token.NewToken(token.BRACE_OPEN, 0, 1),
			token.NewToken(token.LITERAL, 1, 5),
			token.NewToken(token.BRACE_CLOSE, 5, 6),
			token.NewToken(token.EOF, 6, 6),
		},
	},
	{
		input: "a,b",
		tokens: []token.Token{
			token.NewToken(token.LITERAL, 0, 3),
		},
	},
}

func TestNextToken(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			l := NewLexer([]byte(tC.input))

			for _, expected := range tC.tokens {
				actual := l.NextToken()

				assert.Equal(t, expected, actual)
			}
		})
	}
}

// BenchmarkNextToken
func BenchmarkNextToken(b *testing.B) {
	for _, tC := range testCases {
		b.Run(tC.input, func(b *testing.B) {
			input := []byte(tC.input)

			for i := 0; i < b.N; i++ {
				l := NewLexer(input)

				for range tC.tokens {
					l.NextToken()
				}
			}
		})
	}
}
