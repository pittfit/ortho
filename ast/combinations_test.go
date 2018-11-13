package ast

import (
	"testing"
)

var testCases = []struct {
	desc     string
	lists    [][]string
	expected []string
}{
	{
		desc: "Simple list expansion",
		lists: [][]string{
			[]string{"1", "2", "3", "4"},
			[]string{"a", "b"},
			[]string{"_", "-", "+"},
			[]string{"1", "2", "3", "4"},
			[]string{"a", "b"},
			[]string{"_", "-", "+"},
		},
		expected: []string{
			"1a_",
			"1a-",
			"1a+",
			"1b_",
			"1b-",
			"1b+",
			"2a_",
			"2a-",
			"2a+",
			"2b_",
			"2b-",
			"2b+",
			"3a_",
			"3a-",
			"3a+",
			"3b_",
			"3b-",
			"3b+",
			"4a_",
			"4a-",
			"4a+",
			"4b_",
			"4b-",
			"4b+",
		},
	},
}

// TestConbinations …
func TestConbinations(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			combinations(tt.lists)
			// assert.Equal(t, tt.expected, actual)
		})
	}
}

// BenchmarkConbinations …
func BenchmarkConbinations(b *testing.B) {
	for _, tt := range testCases {
		b.Run(tt.desc, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				combinations(tt.lists)
			}
		})
	}
}
