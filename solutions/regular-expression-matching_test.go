package leetcode

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isMatch(s string, p string) bool {
	matched, _ := regexp.MatchString(
		fmt.Sprintf("^%s$", p), s,
	)
	return matched
}

func TestIsMatch(t *testing.T) {
	for _, tc := range []struct {
		name   string
		s, p   string
		output bool
	}{
		{name: "Example 1", s: "aa", p: "a", output: false},
		{name: "Example 2", s: "aa", p: "a*", output: true},
		{name: "Example 3", s: "ab", p: ".*", output: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.output, isMatch(tc.s, tc.p))
		})
	}
}
