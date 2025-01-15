package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 582 (Div. 3) B. Bad Prices
// URL: https://codeforces.com/contest/1213/problem/B
// Time: 2025-01-14 21:49:35

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{

		{
			name:  "case1",
			input: "5\n6\n3 9 4 6 7 5\n1\n1000000\n2\n2 1\n10\n31 41 59 26 53 58 97 93 23 84\n7\n3 2 1 2 3 4 5\n",
			want:  "3\n0\n1\n8\n2\n",
		},
		{
			name:  "case2",
			input: "1\n13\n12 32 12 43 58 48 82 42 22 72 49 29 62\n",
			want:  "8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
