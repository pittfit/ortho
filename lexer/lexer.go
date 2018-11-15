package lexer

import (
	"github.com/pittfit/ortho/token"
)

// Lexer …
type Lexer struct {
	input []byte

	inEscape   bool
	openBraces int

	prevPos int
	currPos int
	nextPos int

	prevChar byte
	currChar byte
	nextChar byte

	prevTok token.Token
	currTok token.Token
	nextTok token.Token
}

// NewLexer …
func NewLexer(input []byte) *Lexer {
	l := &Lexer{
		input:    input,
		prevPos:  -2,
		currPos:  -1,
		nextChar: 0,
	}
	return l
}

func (l *Lexer) readChar() {
	l.prevPos++
	l.currPos++
	l.nextPos++

	l.prevChar = l.charAt(l.prevPos)
	l.currChar = l.charAt(l.currPos)
	l.nextChar = l.charAt(l.nextPos)

	// fmt.Printf("readChar {%v, %v, %v}\n", l.prevPos, l.currPos, l.nextPos)
}

func (l *Lexer) charAt(pos int) byte {
	if pos < 0 {
		return 0
	} else if pos >= len(l.input) {
		return 0
	}

	return l.input[pos]
}

// NextToken …
func (l *Lexer) NextToken() token.Token {
	tok := l.readToken()
	// This line causes the tests to fail
	// l.currTok = tok
	return tok
}

func (l *Lexer) readToken() token.Token {
	l.readChar()

	return l.matchToken()
}

func (l *Lexer) matchToken() token.Token {
	var tok token.Token

match:
	tok.Literal = l.input[l.currPos:l.nextPos]

	switch {
	case l.inEscape:
		l.readChar()
		l.inEscape = false

		tok.Type = token.LITERAL
		tok.Literal = l.input[l.currPos:l.nextPos]
	case l.currChar == '\\':
		l.inEscape = true
		goto match
	case l.currChar == 0:
		tok.Type = token.EOF
	case l.currChar == '{':
		tok.Type = token.BRACE_OPEN
		l.openBraces++
	case l.currChar == '}':
		tok.Type = token.BRACE_CLOSE
		l.openBraces--
	case l.currChar == ',':
		tok.Type = token.LIST_SEPARATOR
	case l.currChar == '*':
		tok.Type = token.WILDCARD
	case l.currChar == '.' && l.nextChar == '.' && l.openBraces > 0:
		l.readChar()

		tok.Type = token.RANGE_SEPARATOR
		tok.Literal = l.input[l.prevPos:l.nextPos]
	default:
		tok.Type = token.LITERAL
		tok.Literal = l.readLiteral()
	}

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: []byte{ch}}
}

func (l *Lexer) readLiteral() []byte {
	start := l.currPos

	for l.isLiteral(l.nextChar) {
		l.readChar()
	}

	return l.input[start:l.nextPos]
}

func (l *Lexer) isLiteral(b byte) bool {
	return !l.isSpecialChar(b)
}

func (l *Lexer) isSpecialChar(b byte) bool {
	if l.inEscape {
		return false
	}

	if b == 0 {
		return true
	}

	if b == '{' || b == '}' || b == ',' {
		return true
	}

	if b == '*' {
		return true
	}

	if b == '\\' {
		return true
	}

	if l.openBraces > 0 {
		if b == '.' {
			return true
		}
	}

	return false
}
