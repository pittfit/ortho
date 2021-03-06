package lexer

import (
	"unicode/utf8"

	"github.com/pittfit/ortho/token"
)

// Lexer …
type Lexer struct {
	input []byte

	inEscape   bool
	openBraces int

	currPos int
	nextPos int
	readPos int

	currChar rune
	nextChar rune

	currTok token.Token
	nextTok token.Token
}

// NewLexer …
func NewLexer(input []byte) *Lexer {
	l := &Lexer{
		input:    input,
		currPos:  -1,
		nextChar: 0,
	}
	l.readToken()
	return l
}

// Input
func (l *Lexer) Input() []byte {
	return l.input
}

func (l *Lexer) readChar() {
	if l.readPos < len(l.input) {
		l.currPos = l.readPos

		r, w := utf8.DecodeRune(l.input[l.readPos:])
		l.currChar = r

		l.readPos += w
	} else {
		l.currChar = 0
	}
}

func (l *Lexer) charAt(pos int) byte {
	if pos < 0 {
		return 0
	} else if pos >= len(l.input) {
		return 0
	}

	return l.input[pos]
}

func (l *Lexer) pos(pos int) int {
	if pos < 0 {
		return 0
	} else if pos >= len(l.input) {
		return len(l.input)
	}

	return pos
}

// All Get all tokens present in the byte stream
func (l *Lexer) All() []token.Token {
	tokens := make([]token.Token, 0, 64)

	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)

		if tok.Type == token.EOF {
			break
		}
	}

	return tokens
}

// NextToken …
func (l *Lexer) NextToken() token.Token {
	l.readToken()

	tok := l.currTok

	if l.currTok.Type != token.LITERAL {
		return tok
	}

	for l.nextTok.Type == token.LITERAL {
		l.readToken()

		tok.Loc.End = l.currTok.Loc.End
	}

	return tok
}

func (l *Lexer) readToken() {
	l.readChar()

	l.currTok, l.nextTok = l.nextTok, l.matchToken()
}

func (l *Lexer) matchToken() token.Token {
	var tok token.Token

match:
	tok.Loc = token.Location{Start: l.currPos, End: l.nextPos}

	switch {
	case l.inEscape:
		l.readChar()
		l.inEscape = false

		tok.Type = token.LITERAL
		tok.Loc = token.Location{Start: l.currPos, End: l.nextPos}
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
		start := l.currPos
		l.readChar()

		tok.Type = token.RANGE_SEPARATOR
		tok.Loc = token.Location{Start: start, End: l.nextPos}
	default:
		tok.Type = token.LITERAL
		tok.Loc = l.readLiteral()
	}

	return tok
}

func (l *Lexer) readLiteral() token.Location {
	start := l.currPos

	for l.isLiteral(l.nextChar) {
		l.readChar()
	}

	return token.Location{Start: start, End: l.nextPos}
}

func (l *Lexer) isLiteral(b rune) bool {
	return !l.isSpecialChar(b)
}

func (l *Lexer) isSpecialChar(b rune) bool {
	if l.inEscape {
		return false
	}

	if b == 0 {
		return true
	}

	// These can appear outside braces and are still considered "operators"
	if b == '{' || b == '}' || b == '*' || b == '\\' {
		return true
	}

	if l.openBraces > 0 {
		if b == '.' {
			return true
		}

		if b == ',' {
			return true
		}
	}

	return false
}
