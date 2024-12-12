package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 886 (Div. 4) A. To My Critics
// URL: https://codeforces.com/contest/1850/problem/A
// Time: 2024-12-12 01:16:39

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "5\n8 1 2\n4 4 5\n9 9 9\n0 0 0\n8 5 3\n",
			want:  "YES\nNO\nYES\nNO\nYES\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
