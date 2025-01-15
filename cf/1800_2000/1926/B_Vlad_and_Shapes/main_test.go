package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 928 (Div. 4) B. Vlad and Shapes
// URL: https://codeforces.com/contest/1926/problem/B
// Time: 2024-12-26 18:27:08

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "6\n3\n000\n011\n011\n4\n0000\n0000\n0100\n1110\n2\n11\n11\n5\n00111\n00010\n00000\n00000\n00000\n10\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n1111111110\n0111111100\n0011111000\n0001110000\n0000100000\n3\n111\n111\n111\n",
			want:  "SQUARE\nTRIANGLE\nSQUARE\nTRIANGLE\nTRIANGLE\nSQUARE\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
