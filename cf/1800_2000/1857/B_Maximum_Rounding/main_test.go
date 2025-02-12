package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 891 (Div. 3) B. Maximum Rounding
// URL: https://codeforces.com/contest/1857/problem/B
// Time: 2025-02-12 15:26:05

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "10\n1\n5\n99\n913\n1980\n20444\n20445\n60947\n419860\n40862016542130810467\n",
			want:  "1\n10\n100\n1000\n2000\n20444\n21000\n100000\n420000\n41000000000000000000\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
