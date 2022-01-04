package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkIfPangram(sentence string) bool {
	var seq = make([]bool, 26)

	for _, char := range sentence {
		seq[char-'a'] = true
	}

	for _, i := range seq {
		if !i {
			return false
		}
	}
	return true
}

func TestCheckIfPangram(t *testing.T) {
	for _, tc := range []struct {
		Input  string
		Output bool
	}{
		{Input: "thequickbrownfoxjumpsoverthelazydog", Output: true},
		{Input: "thequickbrownfoxjumpsoverthelazdddddddddddydog", Output: true},
		{Input: "leetcode", Output: false},
		{Input: "e", Output: false},
	} {
		assert.Equal(t, tc.Output, checkIfPangram(tc.Input), tc.Input)
	}
}
