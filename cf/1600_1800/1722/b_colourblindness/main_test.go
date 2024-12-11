package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 817 (Div. 4) B. Colourblindness
// URL: https://codeforces.com/contest/1722/problem/B
// Time: 2024-12-11 12:08:12

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "6\n2\nRG\nRB\n4\nGRBG\nGBGB\n5\nGGGGG\nBBBBB\n7\nBBBBBBB\nRRRRRRR\n8\nRGBRRGBR\nRGGRRBGR\n1\nG\nG\n",
			want:  "YES\nNO\nYES\nNO\nYES\nYES\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
