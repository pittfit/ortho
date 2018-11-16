package parser

import (
	"github.com/pittfit/ortho/ast"
	"github.com/pittfit/ortho/lexer"
	"github.com/pittfit/ortho/token"
)

type Parser struct {
	l *lexer.Lexer

	prevTok token.Token
	currTok token.Token
	nextTok token.Token
}

// NewParser
func NewParser(input []byte) *Parser {
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
}

// Parse parse an expression into an abstract syntax tree
func (p *Parser) Parse() *ast.AST {
	return ast.NewAST(p.l.Input(), p.parseSubExpr()).Optimize()
}

func (p *Parser) parseSubExpr() ast.Node {
	nodes := []ast.Node{}
	isList := false

	for p.currTok.Type != token.EOF && p.currTok.Type != token.BRACE_CLOSE {
		p.readToken()

		start, end := p.currTok.Loc.Start, p.currTok.Loc.End

		if p.currTok.Type == token.BRACE_OPEN {
			if p.prevTok.Type == token.LIST_SEPARATOR {
				nodes = append(nodes, ast.NilNode())
			}

			nodes = append(nodes, p.parseSubExpr())
		} else if p.currTok.Type == token.LITERAL {
			nodes = append(nodes, ast.TextNode(start, end))
		} else if p.currTok.Type == token.WILDCARD {
			nodes = append(nodes, ast.WildcardNode(start, end))
		} else if p.currTok.Type == token.LIST_SEPARATOR {
			// We've hit the first list separator
			// Lift the already parsed nodes into a sequence
			if !isList {
				isList = true

				if len(nodes) > 0 {
					nodes = []ast.Node{ast.SequenceNode(nodes...)}
				}
			}

			if p.prevTok.Type == token.BRACE_OPEN {
				nodes = append(nodes, ast.NilNode())
			}

			nodes = append(nodes, p.parseListItem())
		}
	}

	if isList {
		return ast.ListNode(nodes...)
	}

	return ast.SequenceNode(nodes...)
}

func (p *Parser) parseListItem() ast.Node {
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
