package ast

import (
	"strings"
)

func combinations(lists [][]string) []string {
	listCount := len(lists)

	pos := make([]int, listCount)
	indexes := make([]int, listCount)
	iterations := 1

	for idx := range lists {
		len := len(lists[idx])
		pos[idx] = 0
		indexes[idx] = len - 1

		iterations *= len
	}

	// One of the lists is empty…
	if iterations == 0 {
		return []string{}
	}

	lastListIndex := listCount - 1
	lastListLastIndex := indexes[lastListIndex]

	var strs []string
	var sb strings.Builder

	for i := 0; i < iterations; i++ {
		// Get the current combination
		sb.Reset()
		for j := 0; j < listCount; j++ {
			sb.WriteString(lists[j][pos[j]])
		}
		strs = append(strs, sb.String())

		// Move to the next combination
		// when no wrapping is needed
		if pos[lastListIndex] != lastListLastIndex {
			pos[lastListIndex]++
			continue
		}

		// Move to the next combination by
		// wrapping around from len() to 0
		for j := lastListIndex; j >= 0; j-- {
			// 1. Increment the position of the furthest list that has not hit the end
			if pos[j] != indexes[j] {
				pos[j]++
				break
			}

			// 2. Reset positions for lists that have hit the end
			pos[j] = 0
		}
	}

	return strs
}
