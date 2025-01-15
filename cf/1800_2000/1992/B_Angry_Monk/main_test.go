package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 957 (Div. 3) B. Angry Monk
// URL: https://codeforces.com/contest/1992/problem/B
// Time: 2025-01-10 01:58:47

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "4\n5 3\n3 1 1\n5 2\n3 2\n11 4\n2 3 1 5\n16 6\n1 6 1 1 1 6\n",
			want:  "2\n3\n9\n15\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
