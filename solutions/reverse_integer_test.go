package leetcode

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse(x int) int {
	s := fmt.Sprintf("%d", x)
	if s[0] == '-' {
		s = s[1:] + "-"
	}
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	d, _ := strconv.Atoi(string(runes))
	if d > 2147483648 || d < -2147483648 {
		return 0
	}
	return d
}

func TestReverse(t *testing.T) {
	cases := []struct {
		x    int
		want int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{1534236469, 0},
		{-1534236469, 0},
	}

	for _, tc := range cases {
		assert.Equal(t, reverse(tc.x), tc.want)
	}
}
