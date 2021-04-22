package leetcode

import (
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func myAtoi(s string) int {
	var st []string
	isStartFlag := false
	for _, char := range s {
		if !isStartFlag {
			if char == '-' || char == '+' || unicode.IsDigit(char) {
				st = append(st, string(char))
				isStartFlag = true
				continue
			}
			if char == ' ' {
				continue
			}
			if !unicode.IsDigit(char) {
				return 0
			}
			continue
		}
		if !unicode.IsDigit(char) {
			break
		}
		st = append(st, string(char))

	}
	iv, _ := strconv.Atoi(strings.Join(st, ""))
	switch {
	case iv > 2147483647:
		return 2147483647
	case iv < -2147483648:
		return -2147483648

	}
	return iv
}

func TestMyAtoi(t *testing.T) {
	for _, tc := range []struct {
		Input  string
		Output int
	}{
		// {Input: "42", Output: 42},
		{Input: "   -42", Output: -42},
		{Input: "4193 with words", Output: 4193},
		{Input: "words and 987", Output: 0},
		{Input: "-91283472332", Output: -2147483648},
		{Input: "21474836460", Output: 2147483647},
		{Input: "   +0 123", Output: 0},
		{Input: "-5-", Output: -5},
		{Input: "3.14159", Output: 3},
	} {
		assert.Equal(t, tc.Output, myAtoi(tc.Input), tc.Input)
	}
}
