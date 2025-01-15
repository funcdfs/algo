package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 957 (Div. 3) D. Test of Love
// URL: https://codeforces.com/contest/1992/problem/D
// Time: 2025-01-12 13:33:41

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "6\n6 2 0\nLWLLLW\n6 1 1\nLWLLLL\n6 1 1\nLWLLWL\n6 2 15\nLWLLCC\n6 10 0\nCCCCCC\n6 6 1\nWCCCCW\n",
			want:  "YES\nYES\nNO\nNO\nYES\nYES\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
