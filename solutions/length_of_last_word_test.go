package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLastWord(s string) int {
	src := strings.TrimRight(s, " ")
	var counter int

	for i := len(src) - 1; i >= 0; i-- {

		if s[i] == byte(' ') {
			return counter
		}
		counter++
	}
	return counter
}

func TestLengthOfLastWord(t *testing.T) {
	for _, tc := range []struct {
		Input  string
		Output int
	}{
		{Input: "Hello World", Output: 5},
		{Input: " ", Output: 0},
		{Input: "a ", Output: 1},
	} {
		assert.Equal(t, tc.Output, lengthOfLastWord(tc.Input))
	}
}
