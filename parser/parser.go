package parser

import (
	"fmt"

	"github.com/pittfit/ortho/ast"
	"github.com/pittfit/ortho/lexer"
	"github.com/pittfit/ortho/token"
	"github.com/pittfit/ortho/tracing"
)

type Parser struct {
	l *lexer.Lexer

	prevTok token.Token
	currTok token.Token
	nextTok token.Token
}

// NewParser
func NewParser(input []byte) *Parser {
	defer tracing.End(tracing.Begin("NewParser", ""))
	tracing.Debug("Input", string(input))

	p := &Parser{
		l: lexer.NewLexer(input),
	}

	p.readToken()
	p.readToken()

	return p
}

func (p *Parser) readToken() {
	p.prevTok = p.currTok
	p.currTok = p.nextTok
	p.nextTok = p.l.NextToken()

	tracing.Call("readToken", fmt.Sprintf("(%v, %v)", p.currTok, p.nextTok))
}

// Parse parse an expression into an abstract syntax tree
func (p *Parser) Parse() *ast.AST {
	defer tracing.End(tracing.Begin("parse", ""))

	// Top-level parsing continues until EOF
	nodes := []ast.Node{}

	for p.currTok.Type != token.EOF {
		nodes = append(nodes, p.parseToken())

		p.readToken()
	}

	return ast.NewAST(p.l.Input(), ast.SequenceNode(nodes...)).Optimize()
}

func (p *Parser) parseToken() ast.Node {
	defer tracing.End(tracing.Begin("parseToken", ""))

	start, end := p.currTok.Loc.Start, p.currTok.Loc.End

	if p.currTok.Type == token.BRACE_OPEN {
		return p.parseSubExpr()
	}

	if p.currTok.Type == token.LITERAL {
		return ast.TextNode(start, end)
	}

	if p.currTok.Type == token.WILDCARD {
		return ast.WildcardNode(start, end)
	}

	if p.currTok.Type == token.RANGE_SEPARATOR {
		return p.parseRange()
	}

	if p.currTok.Type == token.LIST_SEPARATOR {
		return p.parseListItem()
	}

	panic("Unexpected token " + p.currTok.Type.String())
}

func (p *Parser) parseSubExpr() ast.Node {
	defer tracing.End(tracing.Begin("parseSubExpr", ""))

	nodes := []ast.Node{}
	isList := false

	for p.currTok.Type != token.EOF && p.currTok.Type != token.BRACE_CLOSE {
		p.readToken()

		if p.prevTok.Type == token.LIST_SEPARATOR && p.currTok.Type == token.BRACE_OPEN {
			nodes = append(nodes, ast.NilNode())
		}

		if p.currTok.Type == token.LITERAL && p.nextTok.Type == token.RANGE_SEPARATOR {
			// This token will be consumed by p.parseRange()
			continue
		}

		if p.currTok.Type == token.LIST_SEPARATOR && !isList {
			// We've hit the first list separator
			// Lift the already parsed nodes into a sequence
			isList = true

			if len(nodes) > 0 {
				nodes = []ast.Node{ast.SequenceNode(nodes...)}
			}
		}

		if p.prevTok.Type == token.BRACE_OPEN && p.currTok.Type == token.LIST_SEPARATOR {
			nodes = append(nodes, ast.NilNode())
		}

		if p.currTok.Type == token.BRACE_OPEN {
		} else if p.currTok.Type == token.LITERAL {
			nodes = append(nodes, p.parseToken())
		} else if p.currTok.Type == token.WILDCARD {
			nodes = append(nodes, p.parseToken())
		} else if p.currTok.Type == token.RANGE_SEPARATOR {
			nodes = append(nodes, p.parseToken())
		} else if p.currTok.Type == token.LIST_SEPARATOR {
			nodes = append(nodes, p.parseToken())
		}
	}

	if isList {
		return ast.ListNode(nodes...)
	}

	return ast.SequenceNode(nodes...)
}

func (p *Parser) parseListItem() ast.Node {
	defer tracing.End(tracing.Begin("parseListItem", ""))

	nodes := []ast.Node{}

	for p.nextTok.Type != token.LIST_SEPARATOR && p.nextTok.Type != token.BRACE_CLOSE && p.nextTok.Type != token.EOF {
		p.readToken()

		start, end := p.currTok.Loc.Start, p.currTok.Loc.End

		if p.currTok.Type == token.BRACE_OPEN {
			nodes = append(nodes, p.parseSubExpr())
		} else if p.currTok.Type == token.LITERAL {
			nodes = append(nodes, ast.TextNode(start, end))
		} else if p.currTok.Type == token.WILDCARD {
			nodes = append(nodes, ast.WildcardNode(start, end))
		}
	}

	if len(nodes) == 0 && p.nextTok.Type == token.BRACE_CLOSE {
		nodes = append(nodes, ast.NilNode())
	}

	return ast.SequenceNode(nodes...)
}

func (p *Parser) parseRange() ast.Node {
	defer tracing.End(tracing.Begin("parseRange", ""))

	// Is valid range
	tracing.Debug("parseRange", "Check prev == literal")
	if p.prevTok.Type != token.LITERAL {
		tracing.Debug("parseRange", "Invalid range. Previous token was "+p.prevTok.Type.String())
		start, end := p.currTok.Loc.Start, p.currTok.Loc.End

		return ast.TextNode(start, end)
	}

	// The next token must be a literal
	tracing.Debug("parseRange", "Check next == literal")
	if p.nextTok.Type != token.LITERAL {
		tracing.Debug("parseRange", "Invalid range. Next token was "+p.prevTok.Type.String())
		start, end := p.currTok.Loc.Start, p.currTok.Loc.End

		return ast.TextNode(start, end)
	}

	// This looks to be a valid range
	tracing.Debug("parseRange", "Possibly valid range")

	tracing.Debug("parseRange", "Save prev literal")
	startNode := ast.TextNode(p.prevTok.Loc.Start, p.prevTok.Loc.End)

	tracing.Debug("parseRange", "Save next literal")
	endNode := ast.TextNode(p.nextTok.Loc.Start, p.nextTok.Loc.End)

	// Consume the range token
	tracing.Debug("parseRange", "Consume range")
	p.readToken()

	// Consume the next literal token
	tracing.Debug("parseRange", "Consume next literal")
	p.readToken()

	// We're left either (as the current token)
	// 1. A brace close
	if p.currTok.Type == token.BRACE_CLOSE {
		tracing.Debug("parseRange", "Current token is a closing brace")

		return ast.NumericRangeNode(startNode, endNode, ast.NilNode())
	}

	// 2. A range separator followed by a literal followed by a brace close
	if p.currTok.Type != token.RANGE_SEPARATOR {
		// FIXME: Bail properly
		return ast.NilNode()
	}

	// Consume the range separator
	stepNode := ast.TextNode(p.nextTok.Loc.Start, p.nextTok.Loc.End)
	p.readToken()

	if p.currTok.Type == token.LITERAL && p.nextTok.Type == token.BRACE_CLOSE {
		p.readToken()

		return ast.NumericRangeNode(startNode, endNode, stepNode)
	}

	// FIXME: Bail properly
	return ast.NilNode()
}
