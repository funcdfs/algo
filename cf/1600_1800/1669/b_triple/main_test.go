package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 784 (Div. 4) B. Triple
// URL: https://codeforces.com/contest/1669/problem/B
// Time: 2024-12-12 01:24:07

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "7\n1\n1\n3\n2 2 2\n7\n2 2 3 3 4 2 2\n8\n1 4 3 4 3 2 4 1\n9\n1 1 1 2 2 2 3 3 3\n5\n1 5 2 4 3\n4\n4 4 4 4\n",
			want:  "-1\n2\n2\n4\n3\n-1\n4\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
