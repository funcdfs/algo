package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 957 (Div. 3) A. Only Pluses
// URL: https://codeforces.com/contest/1992/problem/A
// Time: 2025-01-10 01:23:49

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{

		{
			name:  "case1",
			input: "2\n2 3 4\n10 1 10\n",
			want:  "100\n600\n",
		},
		{
			name:  "wa",
			input: "2\n1 1 4\n1 1 5\n",
			want:  "48\n60\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
