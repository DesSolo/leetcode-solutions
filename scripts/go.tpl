package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

{code}

func Test{meta_name_normal}(t *testing.T) {{
	for _, tc := range []struct {{
		name   string
		input  []int
		output int
	}}{{
		{{name: "Example 1", input: []int{{1}}, output: 2}},
	}} {{
		t.Run(tc.name, func(t *testing.T) {{
			assert.Equal(t, tc.output, {meta_name}(tc.input), tc.input)
		}})
	}}
}}