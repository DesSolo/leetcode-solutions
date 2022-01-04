package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func defangIPaddrIF(address string) string {
	var result []byte
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			result = append(result, []byte("[.]")...)
			continue
		}
		result = append(result, address[i])
	}
	return string(result)
}

func defangIPaddrSWITCH(address string) string {
	var result []byte
	for i := 0; i < len(address); i++ {
		switch address[i] {
		case '.':
			result = append(result, []byte("[.]")...)
		default:
			result = append(result, address[i])
		}
	}
	return string(result)
}

func defangIPaddrREPLACE(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}

func TestDefangIPaddr(t *testing.T) {
	for _, tc := range []struct {
		Input  string
		Output string
	}{
		{Input: "1.1.1.1", Output: "1[.]1[.]1[.]1"},
		{Input: "255.100.50.0", Output: "255[.]100[.]50[.]0"},
	} {
		assert.Equal(t, tc.Output, defangIPaddrSWITCH(tc.Input), tc.Input)
	}
}

func BenchmarkIF(b *testing.B) {
	for n := 0; n < b.N; n++ {
		defangIPaddrIF("255.100.50.0")
	}
}

func BenchmarkSWITCH(b *testing.B) {
	for n := 0; n < b.N; n++ {
		defangIPaddrSWITCH("255.100.50.0")
	}
}

func BenchmarkREPLACE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		defangIPaddrREPLACE("255.100.50.0")
	}
}
