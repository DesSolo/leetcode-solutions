package leetcode

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func toLowerCase(str string) string {
	var b strings.Builder
	for _, c := range str {
		if unicode.IsUpper(c) {
			b.WriteRune(
				unicode.ToLower(c),
			)
			continue
		}
		b.WriteRune(c)
	}
	return b.String()
}

func TestToLowerCase(t *testing.T) {
	for _, tc := range []struct {
		Input  string
		Output string
	}{
		{Input: "Hello", Output: "hello"},
		{Input: "here", Output: "here"},
		{Input: "LOVELY", Output: "lovely"},
	} {
		assert.Equal(t, tc.Output, toLowerCase(tc.Input), tc.Input)
	}
}


func BenchmarkRune(b *testing.B) {
	for n := 0; n < b.N; n++ {
		toLowerCase("Hello")
	}
}

func BenchmarkSys(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.ToLower("Hello")
	}
}
