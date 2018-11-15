package lexer

import (
	"testing"

	"github.com/pittfit/ortho/token"
	"github.com/stretchr/testify/assert"
)

type result struct {
	tokTyp token.TokenType
	tokLit string
}

func TestNextToken(t *testing.T) {
	testCases := []struct {
		input  string
		tokens []result
	}{
		{
			input: "{,}",
			tokens: []result{
				{token.BRACE_OPEN, "{"},
				{token.LIST_SEPARATOR, ","},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "{a,}",
			tokens: []result{
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "a"},
				{token.LIST_SEPARATOR, ","},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "{ab,c}",
			tokens: []result{
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "ab"},
				{token.LIST_SEPARATOR, ","},
				{token.LITERAL, "c"},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "{1..10}",
			tokens: []result{
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "1"},
				{token.RANGE_SEPARATOR, ".."},
				{token.LITERAL, "10"},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "{1...10}",
			tokens: []result{
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "1"},
				{token.RANGE_SEPARATOR, ".."},
				{token.LITERAL, ".10"},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "{1....10}",
			tokens: []result{
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "1"},
				{token.RANGE_SEPARATOR, ".."},
				{token.RANGE_SEPARATOR, ".."},
				{token.LITERAL, "10"},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "{01..10..2}",
			tokens: []result{
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "01"},
				{token.RANGE_SEPARATOR, ".."},
				{token.LITERAL, "10"},
				{token.RANGE_SEPARATOR, ".."},
				{token.LITERAL, "2"},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "foo/*/bar{01..10..2}.jpg",
			tokens: []result{
				{token.LITERAL, "foo/"},
				{token.WILDCARD, "*"},
				{token.LITERAL, "/bar"},
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "01"},
				{token.RANGE_SEPARATOR, ".."},
				{token.LITERAL, "10"},
				{token.RANGE_SEPARATOR, ".."},
				{token.LITERAL, "2"},
				{token.BRACE_CLOSE, "}"},
				{token.LITERAL, ".jpg"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "foo/*/bar{01..10..2}..jpg",
			tokens: []result{
				{token.LITERAL, "foo/"},
				{token.WILDCARD, "*"},
				{token.LITERAL, "/bar"},
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "01"},
				{token.RANGE_SEPARATOR, ".."},
				{token.LITERAL, "10"},
				{token.RANGE_SEPARATOR, ".."},
				{token.LITERAL, "2"},
				{token.BRACE_CLOSE, "}"},
				{token.LITERAL, "..jpg"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "\\{{a,b}",
			tokens: []result{
				{token.LITERAL, "{"},
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "a"},
				{token.LIST_SEPARATOR, ","},
				{token.LITERAL, "b"},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "\\\\{{a,b}",
			tokens: []result{
				{token.LITERAL, "\\"},
				{token.BRACE_OPEN, "{"},
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "a"},
				{token.LIST_SEPARATOR, ","},
				{token.LITERAL, "b"},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "\\a{{a,b}",
			tokens: []result{
				{token.LITERAL, "a"},
				{token.BRACE_OPEN, "{"},
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "a"},
				{token.LIST_SEPARATOR, ","},
				{token.LITERAL, "b"},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
		{
			input: "{a\\,b}",
			tokens: []result{
				{token.BRACE_OPEN, "{"},
				{token.LITERAL, "a"},
				{token.LITERAL, ","},
				{token.LITERAL, "b"},
				{token.BRACE_CLOSE, "}"},
				{token.EOF, "\x00"},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			l := NewLexer([]byte(tC.input))

			for _, expected := range tC.tokens {
				tok := l.NextToken()

				actual := result{tok.Type, string(tok.Literal)}

				assert.Equal(t, expected, actual)
			}
		})
	}
}
