package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Educational Codeforces Round 167 (Rated for Div. 2) A. Catch the Coin
// URL: https://codeforces.com/contest/1989/problem/A
// Time: 2025-01-14 16:10:04

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "5\n24 42\n-2 -1\n-1 -2\n0 -50\n15 0\n",
			want:  "YES\nYES\nNO\nNO\nYES\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
