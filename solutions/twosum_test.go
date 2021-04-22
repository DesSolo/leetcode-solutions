package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(nums []int, target int) []int {
	total := len(nums)
	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			if j == i {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		Nums   []int
		Target int
		Want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{3, 2, 3}, 6, []int{0, 2}},
	}

	for _, tc := range cases {
		assert.Equal(t, twoSum(tc.Nums, tc.Target), tc.Want)
	}
}
