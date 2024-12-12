package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 886 (Div. 4) B. Ten Words of Wisdom
// URL: https://codeforces.com/contest/1850/problem/B
// Time: 2024-12-12 01:19:04

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "3\n5\n7 2\n12 5\n9 3\n9 4\n10 1\n3\n1 2\n3 4\n5 6\n1\n1 43\n",
			want:  "4\n3\n1\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
