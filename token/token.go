package token

import "fmt"

// TokenType
type TokenType int

// Currently only single byte runes are supported for these.
const (
	INVALID TokenType = iota
	EOF
	ESCAPE
	LITERAL
	WILDCARD
	BRACE_OPEN
	BRACE_CLOSE
	LIST_SEPARATOR
	RANGE_SEPARATOR
	STEP_SEPARATOR
)

// Token
type Token struct {
	Type TokenType
	Loc  Location
}

// NewToken …
func NewToken(typ TokenType, start int, end int) Token {
	return Token{Type: typ, Loc: Location{Start: start, End: end}}
}

func (t Token) String() string {
	switch t.Type {
	case LITERAL:
		return fmt.Sprintf("%v(%v:%v)", t.Type.String(), t.Loc.Start, t.Loc.End)
	default:
		return t.Type.String()
	}
}

func (t TokenType) String() string {
	switch t {
	case INVALID:
		return "INVALID"
	case EOF:
		return "EOF"
	case LITERAL:
		return "LITERAL"
	case WILDCARD:
		return "WILDCARD"
	case BRACE_OPEN:
		return "BRACE_OPEN"
	case BRACE_CLOSE:
		return "BRACE_CLOSE"
	case LIST_SEPARATOR:
		return "LIST_SEPARATOR"
	case RANGE_SEPARATOR:
		return "RANGE_SEPARATOR"
	case STEP_SEPARATOR:
		return "STEP_SEPARATOR"
	default:
		return "UNKNOWN"
	}
}
