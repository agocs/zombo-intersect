package zombo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	type test struct {
		inputA []string
		inputB []string
		want   []string
	}

	tests := []test{
		{inputA: []string{"a", "b", "c"}, inputB: []string{"c", "c", "c"}, want: []string{"c"}},
		{inputA: []string{"a", "b", "c"}, inputB: []string{"c"}, want: []string{"c"}},
		{inputA: []string{"a", "b", "c"}, inputB: []string{"c", "d", "e"}, want: []string{"c"}},
		{inputA: []string{"a", "b", "c"}, inputB: []string{"a", "b", "e"}, want: []string{"a", "b"}},
		{inputA: []string{"a", "b", "c"}, inputB: []string{"a", "c", "e"}, want: []string{"a", "c"}},
		{inputA: []string{"a", "b", "c"}, inputB: []string{"d", "e", "f"}, want: []string{}},
		{inputA: []string{"a", "b", "b", "c"}, inputB: []string{"b", "b"}, want: []string{"b"}},
	}

	for _, tc := range tests {
		got := intersect(tc.inputA, tc.inputB)
		assert.Equal(t, tc.want, got)
	}
}

func TestZombofy(t *testing.T) {
	type test struct {
		input [][]string
		want  []string
	}

	tests := []test{
		{input: [][]string{{"a", "b", "c"}, {"c", "c", "c"}}, want: []string{"c"}},
		{input: [][]string{{"a", "b", "c"}, {"c", "c", "c"}, {"a", "c", "e"}}, want: []string{"c"}},
		{input: [][]string{{"a", "b", "c"}, {"b", "b", "c"}, {"a", "c", "e"}}, want: []string{"c"}},
		{input: [][]string{{"a", "b", "c"}}, want: []string{"a", "b", "c"}},
	}

	for _, tc := range tests {
		got := zombofyLinear(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
