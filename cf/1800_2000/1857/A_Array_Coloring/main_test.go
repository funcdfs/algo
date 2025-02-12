package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 891 (Div. 3) A. Array Coloring
// URL: https://codeforces.com/contest/1857/problem/A
// Time: 2025-02-12 15:20:07

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "7\n8\n1 2 4 3 2 3 5 4\n2\n4 7\n3\n3 9 8\n2\n1 7\n5\n5 4 3 2 1\n4\n4 3 4 5\n2\n50 48\n",
			want:  "YES\nNO\nYES\nYES\nNO\nYES\nYES\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
