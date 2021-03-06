package ast

import (
	"fmt"
	"regexp"
	"strings"
)

// PatternOpts
type PatternOpts struct {
	Anchored bool
}

func (a *AST) ToPattern(opts PatternOpts) (string, error) {
	pattern, err := a.nodeToSubpattern(a.Root)

	if err != nil {
		return "", err
	}

	if opts.Anchored {
		pattern = fmt.Sprintf("^%s$", pattern)
	}

	return pattern, nil
}

func (a *AST) nodeToSubpattern(n Node) (string, error) {
	switch n.Type {
	case TypeNil:
		return "", nil
	case TypeText:
		return regexp.QuoteMeta(string(a.slice(n.Loc))), nil
	case TypeWildcard:
		return ".*?", nil
	case TypeSequence:
		patterns, err := a.perChildSubpatterns(n)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s", strings.Join(patterns, "")), nil
	case TypeList:
		patterns, err := a.perChildSubpatterns(n)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("(%s)", strings.Join(patterns, "|")), nil
	case TypeRangeNumeric:
		start, end, step := n.Children[0].Loc, n.Children[1].Loc, n.Children[2].Loc

		numericStrs, err := stringsForNumericRange(a.slice(start), a.slice(end), a.slice(step))
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("(%s)", strings.Join(numericStrs, "|")), nil
	}

	return "", nil
}

func (a *AST) perChildSubpatterns(n Node) ([]string, error) {
	var err error
	var patterns = make([]string, len(n.Children))

	for idx, child := range n.Children {
		patterns[idx], err = a.nodeToSubpattern(child)
		if err != nil {
			return nil, err
		}
	}

	return patterns, nil
}
