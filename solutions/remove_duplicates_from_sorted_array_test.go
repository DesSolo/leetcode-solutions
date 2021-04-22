package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var i int

	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tc := range []struct {
		Input  []int
		Output int
	}{
		{Input: []int{1, 1, 2}, Output: 2},
		{Input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, Output: 5},
	} {
		assert.Equal(t, tc.Output, removeDuplicates(tc.Input))
	}
}
