package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countConsistentStrings(allowed string, words []string) int {
	var result int

	for i := 0; i < len(words); i++ {
		var matched int

		for j := 0; j < len(words[i]); j++ {
			for m := 0; m < len(allowed); m++ {
				if allowed[m] == words[i][j] {
					matched++
				}
			}
		}
		if matched == len(words[i]) {
			result++
		}
	}

	return result

}

func TestCountConsistentStrings(t *testing.T) {
	for _, tc := range []struct {
		allowed string
		words   []string
		output  int
	}{
		{allowed: "ab", words: []string{"ad", "bd", "aaab", "baa", "badab"}, output: 2},
		{allowed: "abc", words: []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, output: 7},
		{allowed: "cad", words: []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}, output: 4},
	} {
		assert.Equal(t, tc.output, countConsistentStrings(tc.allowed, tc.words), tc.allowed)
	}
}
