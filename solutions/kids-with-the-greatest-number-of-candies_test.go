package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	totalCandies := len(candies)
	result := make([]bool, totalCandies)
	var maxCandies int
	for i := 0; i < totalCandies; i++ {
		if maxCandies < candies[i] {
			maxCandies = candies[i]
		}
	}
	for i := 0; i < totalCandies; i++ {
		result[i] = (candies[i] + extraCandies) >= maxCandies
	}
	return result
}

func TestKidsWithCandies(t *testing.T) {
	for _, tc := range []struct {
		candies      []int
		extraCandies int
		Output       []bool
	}{
		{candies: []int{2, 3, 5, 1, 3}, extraCandies: 3, Output: []bool{true, true, true, false, true}},
		{candies: []int{4, 2, 1, 1, 2}, extraCandies: 1, Output: []bool{true, false, false, false, false}},
		{candies: []int{12, 1, 12}, extraCandies: 10, Output: []bool{true, false, true}},
	} {
		assert.Equal(t, tc.Output, kidsWithCandies(tc.candies, tc.extraCandies), tc.candies)
	}
}
