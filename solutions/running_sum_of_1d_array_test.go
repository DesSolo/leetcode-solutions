package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func runningSum(nums []int) []int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		counter += nums[i]
		nums[i] = counter
	}
	return nums
}

func TestRunningSum(t *testing.T) {
	for _, tc := range []struct {
		Input  []int
		Output []int
	}{
		{Input: []int{1, 2, 3, 4}, Output: []int{1, 3, 6, 10}},
		{Input: []int{1, 1, 1, 1, 1}, Output: []int{1, 2, 3, 4, 5}},
		{Input: []int{3, 1, 2, 10, 1}, Output: []int{3, 4, 6, 16, 17}},
	} {
		assert.Equal(t, tc.Output, runningSum(tc.Input), tc.Input)
	}
}
