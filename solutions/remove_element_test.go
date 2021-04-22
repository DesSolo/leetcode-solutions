package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = remove(nums, i)
			i--
		}
	}

	return len(nums)
}

func TestRemoveElement(t *testing.T) {
	for _, tc := range []struct {
		Nums   []int
		Val    int
		Output int
	}{
		{Nums: []int{3, 2, 2, 3}, Val: 3, Output: 2},
		{Nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, Val: 2, Output: 5},
	} {
		assert.Equal(t, tc.Output, removeElement(tc.Nums, tc.Val), tc.Nums)
	}
}
