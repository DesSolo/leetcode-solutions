package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isValid(s string) bool {
	var stack []byte

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default:
			ls := len(stack)
			if ls == 0 {
				return false
			}
			if stack[ls-1] != s[i] {
				return false
			}
			stack = stack[:ls-1]
		}
	}

	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	cases := []struct {
		src   string
		valid bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"]", false},
		{"(])", false},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.valid, isValid(tc.src), tc.src)
	}
}
