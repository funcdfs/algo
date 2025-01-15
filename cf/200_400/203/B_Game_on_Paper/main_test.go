package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 128 (Div. 2) B. Game on Paper
// URL: https://codeforces.com/contest/203/problem/B
// Time: 2025-01-15 17:09:22

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "4 11\n1 1\n1 2\n1 3\n2 2\n2 3\n1 4\n2 4\n3 4\n3 2\n3 3\n4 1\n",
			want:  "10\n",
		},
		{
			name:  "case2",
			input: "4 12\n1 1\n1 2\n1 3\n2 2\n2 3\n1 4\n2 4\n3 4\n3 2\n4 2\n4 1\n3 1\n",
			want:  "-1\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
