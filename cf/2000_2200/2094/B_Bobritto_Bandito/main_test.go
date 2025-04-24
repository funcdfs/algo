package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 1017 (Div. 4) B. Bobritto Bandito
// URL: https://codeforces.com/contest/2094/problem/B
// Time: 2025-04-24 20:56:01

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "4\n4 2 -2 2\n4 1 0 4\n3 3 -1 2\n9 8 -6 3\n",
			want:  "-1 1\n0 1\n-1 2\n-5 3\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
