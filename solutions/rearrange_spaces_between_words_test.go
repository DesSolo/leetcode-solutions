package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reorderSpaces(text string) string {
	return strings.Trim(text, "")

}

func TestReloadSpaces(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{"  this   is  a sentence ", "this   is   a   sentence"},
		{" practice   makes   perfect", "practice   makes   perfect "},
		{"hello   world", "hello   world"},
		{"  walks  udp package   into  bar a", "walks  udp  package  into  bar  a "},
		{"a", "a"},
	}

	for _, tc := range cases {
		assert.Equal(t, reorderSpaces(tc.input), tc.output)
	}
}
