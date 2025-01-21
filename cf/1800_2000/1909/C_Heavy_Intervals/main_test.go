package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Pinely Round 3 (Div. 1 + Div. 2) C. Heavy Intervals
// URL: https://codeforces.com/contest/1909/problem/C
// Time: 2025-01-20 23:03:05

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "2\n2\n8 3\n12 23\n100 100\n4\n20 1 2 5\n30 4 3 10\n2 3 2 3\n",
			want:  "2400\n42\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
