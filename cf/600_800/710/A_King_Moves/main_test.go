package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Educational Codeforces Round 16 A. King Moves
// URL: https://codeforces.com/contest/710/problem/A
// Time: 2025-01-20 22:50:47

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "e4\n",
			want:  "8\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
