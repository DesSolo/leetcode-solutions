package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countVowels(s string) int {
	var vowels int
	for _, char := range s {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels++
		}
	}
	return vowels
}

func halvesAreAlike(s string) bool {
	sLen := len(s)

	return countVowels(s[:sLen/2]) == countVowels(s[sLen/2:])

}

func TestHalvesAreAlike(t *testing.T) {
	for _, tc := range []struct {
		Input  string
		Output bool
	}{
		{Input: "book", Output: true},
		{Input: "textbook", Output: false},
		{Input: "MerryChristmas", Output: false},
		{Input: "AbCdEfGh", Output: true},
	} {
		assert.Equal(t, tc.Output, halvesAreAlike(tc.Input), tc.Input)
	}
}
