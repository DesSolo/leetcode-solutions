package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchInsert(nums []int, target int) int {
	total := len(nums)
	if nums[total-1] < target {
		return total
	}

	for i := 0; i < total; i++ {
		if nums[i] == target || nums[i] > target {
			return i
		}
	}
	return 0
}

func TestSearchInsert(t *testing.T) {
	for _, tc := range []struct {
		Nums   []int
		Target int
		Output int
	}{
		{Nums: []int{1, 3, 5, 6}, Target: 5, Output: 2},
		{Nums: []int{1, 3, 5, 6}, Target: 2, Output: 1},
		{Nums: []int{1, 3, 5, 6}, Target: 7, Output: 4},
		{Nums: []int{1, 3, 5, 22}, Target: 7, Output: 3},
		{Nums: []int{1, 3, 5, 6}, Target: 0, Output: 0},
		{Nums: []int{1}, Target: 0, Output: 0},
	} {
		assert.Equal(t, tc.Output, searchInsert(tc.Nums, tc.Target))
	}
}
