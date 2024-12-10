package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 565 (Div. 3) B. Merge it!
// URL: https://codeforces.com/contest/1176/problem/B
// Time: 2024-12-10 11:43:25

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "2\n5\n3 1 2 3 1\n7\n1 1 1 1 1 2 2\n",
			want:  "3\n3\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
