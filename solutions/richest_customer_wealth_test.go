package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maximumWealth(accounts [][]int) int {
	var result int
	n, m := len(accounts), len(accounts[0])
	for i := 0; i < n; i++ {
		var bs int
		for j := 0; j < m; j++ {
			bs += accounts[i][j]
		}
		if result < bs {
			result = bs
		}
	}

	return result
}

func TestMaximumWealth(t *testing.T) {
	for _, tc := range []struct {
		Input  [][]int
		Output int
	}{
		{Input: [][]int{{1, 2, 3}, {3, 2, 1}}, Output: 6},
		{Input: [][]int{{1, 5}, {7, 3}, {3, 5}}, Output: 10},
		{Input: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, Output: 17},
	} {
		assert.Equal(t, tc.Output, maximumWealth(tc.Input), tc.Input)
	}
}
