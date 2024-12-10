package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 565 (Div. 3) C. Lose it!
// URL: https://codeforces.com/contest/1176/problem/C
// Time: 2024-12-10 11:54:00

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "5\n4 8 15 16 23\n",
			want:  "5\n",
		},
		{
			name:  "case2",
			input: "12\n4 8 4 15 16 8 23 15 16 42 23 42\n",
			want:  "0\n",
		},
		{
			name:  "case3",
			input: "15\n4 8 4 8 15 16 8 16 23 15 16 4 42 23 42\n",
			want:  "3\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
