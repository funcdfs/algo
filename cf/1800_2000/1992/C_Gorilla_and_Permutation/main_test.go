package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 957 (Div. 3) C. Gorilla and Permutation
// URL: https://codeforces.com/contest/1992/problem/C
// Time: 2025-01-10 02:30:36

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "3\n5 2 5\n3 1 3\n10 3 8\n",
			want:  "5 3 4 1 2\n3 2 1\n10 9 8 4 7 5 6 1 2 3\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
