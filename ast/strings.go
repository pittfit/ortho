package ast

// ToStrings …
func (a *AST) ToStrings() ([]string, error) {
	return a.nodeToStrings(a.root)
}

func (a *AST) nodeToStrings(n Node) ([]string, error) {
	var strings []string

	switch n.Type {
	case TypeNil:
		// do nothing
	case TypeText:
		strings = append(strings, string(a.slice(n.Loc)))
	case TypeWildcard:
		strings = append(strings, string(a.slice(n.Loc)))
	case TypeSequence:
		strs, err := a.perChildStrings(n)
		if err != nil {
			return nil, err
		}

		// Convert each of the children to a set of strings
		// Then combine them pulling on from each list until all combos have been found
		strings = append(strings, combinations(strs)...)
	case TypeList:
		strs, err := a.perChildStrings(n)
		if err != nil {
			return nil, err
		}

		for _, strs := range strs {
			strings = append(strings, strs...)
		}
	case TypeRangeNumeric:
		start, end, step := n.Children[0].Loc, n.Children[1].Loc, n.Children[2].Loc

		numericStrs, err := stringsForNumericRange(a.slice(start), a.slice(end), a.slice(step))
		if err != nil {
			return nil, err
		}

		strings = append(strings, numericStrs...)
	}

	return strings, nil
}

func (a *AST) perChildStrings(n Node) ([][]string, error) {
	var err error
	var strings = make([][]string, len(n.Children))

	for idx, child := range n.Children {
		strings[idx], err = a.nodeToStrings(child)
		if err != nil {
			return nil, err
		}
	}

	return strings, nil
}
