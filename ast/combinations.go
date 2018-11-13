package ast

import (
	"strings"
)

func combinations(lists [][]string) []string {
	listCount := len(lists)

	// position in each of the lists
	// [0, 0, 0, 0, 0, 0]
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

	// fmt.Printf("Lists: %v\n", lists)
	// fmt.Printf("Last Indexes: %v\n", indexes)

	var strs []string
	var sb strings.Builder

	for i := 0; i < iterations; i++ {
		// fmt.Printf("[i=%v] Positions %v\n", i, pos)

		// Concat the strings from each of the lists based on the current positions
		sb.Reset()
		for j := 0; j < listCount; j++ {
			sb.WriteString(lists[j][pos[j]])
		}
		strs = append(strs, sb.String())

		// Increment the list counter
		if pos[lastListIndex] != lastListLastIndex {
			// fmt.Printf("[i=%v] lastListIndex++\n", i)
			pos[lastListIndex]++
			continue
		}

		// Walk backwards and:
		// 1. Reset positions for lists that have hit the end
		// 2. Increment the position of the furthest list that has not hit the end
		for j := lastListIndex; j >= 0; j-- {
			// fmt.Printf("[i=%v] Check list: %v\n", i, j)
			if pos[j] != indexes[j] {
				pos[j]++
				break
			}

			pos[j] = 0
		}
	}

	return strs
}
