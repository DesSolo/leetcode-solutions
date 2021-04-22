package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// {"I"             1}
// "V"             5
// "X"             10
// "L"             50
// "C"             100
// "D"             500
// "M"             1000

var dict = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var ints []int
	for _, char := range s {
		ints = append(ints, dict[char])
	}

	total := len(ints)
	var result int
	for i := 0; i < total; i++ {
		curNum := ints[i]
		if i+1 >= total {
			result += curNum
			continue
		}
		nextNum := ints[i+1]
		if curNum < nextNum {
			result += nextNum - curNum
			i++
		} else {
			result += curNum
		}
	}

	return result
}

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tc := range cases {
		assert.Equal(t, romanToInt(tc.s), tc.want)
	}
}
