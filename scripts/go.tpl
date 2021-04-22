package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

{code}

func Test{meta_name_normal}(t *testing.T) {{
	for _, tc := range []struct {{
		Input  []int
		Output int
	}}{{
		{{Input: []int{{1}}, Output: 2}},
		{{Input: []int{{0}}, Output: 5}},
	}} {{
		assert.Equal(t, tc.Output, {meta_name}(tc.Input), tc.Input)
	}}
}}