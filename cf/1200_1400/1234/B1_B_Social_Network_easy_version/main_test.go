package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 590 (Div. 3) B1. Social Network (easy version)
// URL: https://codeforces.com/contest/1234/problem/B1
// Time: 2025-01-15 22:39:43

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "7 2\n1 2 3 2 1 3 2\n",
			want:  "2\n2 1\n",
		},
		{
			name:  "case2",
			input: "10 4\n2 3 3 1 1 2 1 2 3 3\n",
			want:  "3\n1 3 2\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
