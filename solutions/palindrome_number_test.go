package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := fmt.Sprintf("%d", x)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return s == string(runes)
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		x    int
		want bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{-101, false},
	}

	for _, tc := range cases {
		assert.Equal(t, isPalindrome(tc.x), tc.want)
	}
}
