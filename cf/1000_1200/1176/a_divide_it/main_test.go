package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 565 (Div. 3) A. Divide it!
// URL: https://codeforces.com/contest/1176/problem/A
// Time: 2024-12-10 11:30:47

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "7\n1\n10\n25\n30\n14\n27\n1000000000000000000\n",
			want:  "0\n4\n6\n6\n-1\n6\n72\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
