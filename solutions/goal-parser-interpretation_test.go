package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func interpret(command string) string {
	var cmd []byte

	for i := 0; i < len(command); i++ {
		switch command[i] {
		case 'G', 'o', 'a', 'l':
			cmd = append(cmd, command[i])
			continue
		}
		if command[i] == '(' && command[i+1] == ')' {
			cmd = append(cmd, 'o')
		}

	}

	return string(cmd)
}

func TestInterpret(t *testing.T) {
	for _, tc := range []struct {
		Input  string
		Output string
	}{
		{Input: "G()(al)", Output: "Goal"},
		{Input: "G()()()()(al)", Output: "Gooooal"},
		{Input: "(al)G(al)()()G", Output: "alGalooG"},
	} {
		assert.Equal(t, tc.Output, interpret(tc.Input), tc.Input)
	}
}
