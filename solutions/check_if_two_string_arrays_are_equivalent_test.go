package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func TestArrayStringsAreEqual(t *testing.T) {
	for _, tc := range []struct {
		word1  []string
		word2  []string
		output bool
	}{
		{word1: []string{"ab", "c"}, word2: []string{"a", "bc"}, output: true},
		{word1: []string{"a", "cb"}, word2: []string{"ab", "c"}, output: false},
		{word1: []string{"abc", "d", "defg"}, word2: []string{"abcddefg"}, output: true},
	} {
		assert.Equal(t, tc.output, arrayStringsAreEqual(tc.word1, tc.word2), tc.word1, tc.word2)
	}
}
