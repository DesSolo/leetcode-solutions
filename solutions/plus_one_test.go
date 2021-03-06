package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			continue
		}
		digits[i]++
		return digits
	}
	return append([]int{1}, digits...)
}

func TestPlusOne(t *testing.T) {
	for _, tc := range []struct {
		Input  []int
		Output []int
	}{
		{Input: []int{1, 2, 3}, Output: []int{1, 2, 4}},
		{Input: []int{4, 3, 2, 1}, Output: []int{4, 3, 2, 2}},
		{Input: []int{0}, Output: []int{1}},
		{Input: []int{9}, Output: []int{1, 0}},
		{Input: []int{9, 9}, Output: []int{1, 0, 0}},
	} {
		assert.Equal(t, tc.Output, plusOne(tc.Input))
	}
}
