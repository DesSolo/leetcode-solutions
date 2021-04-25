package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pTypes = map[string]int{
	"type": 0,
	"color": 1,
	"name": 2,
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	sIndex := pTypes[ruleKey]

	var counter int
	for i:=0; i<len(items); i++ {
		if items[i][sIndex] == ruleValue {
			counter++
		}
	}

	return counter
}

func TestCountMatches(t *testing.T) {
	for _, tc := range []struct {
		items              [][]string
		ruleKey, ruleValue string
		output             int
	}{
		{items: [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, ruleKey: "color", ruleValue: "silver", output: 1},
		{items: [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, ruleKey: "type", ruleValue: "phone", output: 2},
	} {
		assert.Equal(t, tc.output, countMatches(tc.items, tc.ruleKey, tc.ruleValue), tc.items)
	}
}
