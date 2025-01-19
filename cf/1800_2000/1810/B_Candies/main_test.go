package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - CodeTON Round 4 (Div. 1 + Div. 2, Rated, Prizes!) B. Candies
// URL: https://codeforces.com/contest/1810/problem/B
// Time: 2025-01-19 16:05:16

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "4\n2\n3\n7\n17\n",
			want:  "-1\n1\n2\n2\n2 2\n4\n2 1 1 1\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
