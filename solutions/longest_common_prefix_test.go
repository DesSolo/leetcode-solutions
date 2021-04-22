package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	pref := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], pref) {
			pref = pref[:len(pref)-1]
		}
	}
	return pref
}

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		sl   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"pop", "p", "p"}, "p"},
		{[]string{"lol", "lol", ""}, ""},
	}

	for _, tc := range cases {
		assert.Equal(t, longestCommonPrefix(tc.sl), tc.want)
	}
}
