package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 928 (Div. 4) A. Vlad and the Best of Five
// URL: https://codeforces.com/contest/1926/problem/A
// Time: 2024-12-26 18:24:45

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "8\nABABB\nABABA\nBBBAB\nAAAAA\nBBBBB\nBABAA\nAAAAB\nBAAAA\n",
			want:  "B\nA\nB\nA\nB\nA\nA\nA\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
