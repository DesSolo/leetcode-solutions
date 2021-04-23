package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reorderSpaces(text string) string {
	if len(text) <= 1 {
		return text
	}

	var result []rune

	// count spaces
	var spTotal int
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			spTotal++
		}
	}

	src := strings.Fields(text)
	srcTotal := len(src)
	if srcTotal <= 1 {
		srcTotal = 2
	}
	spMiddle := spTotal / (srcTotal - 1)
	for index, word := range src {
		for _, char := range word {
			result = append(result, char)
		}
		if index == srcTotal-1 {
			continue
		}
		for i := 0; i < spMiddle; i++ {
			result = append(result, ' ')
		}
	}
	// add extra
	spExtra := spTotal % (srcTotal - 1)
	for i := 0; i < spExtra; i++ {
		result = append(result, ' ')
	}

	return string(result)
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
		{"  hello", "hello  "},
	}

	for _, tc := range cases {
		assert.Equal(t, reorderSpaces(tc.input), tc.output)
	}
}
